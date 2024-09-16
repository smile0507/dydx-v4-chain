package types

import (
	"fmt"
	"sync"
	"time"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/dydxprotocol/v4-chain/protocol/daemons/pricefeed/api"
	pricefeedmetrics "github.com/dydxprotocol/v4-chain/protocol/daemons/pricefeed/metrics"
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/lib/metrics"
	"github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	gometrics "github.com/hashicorp/go-metrics"
)

// MarketToExchangePrices maintains price info for multiple markets. Each
// market can support prices from multiple exchange sources. Specifically,
// MarketToExchangePrices supports methods to update prices and to retrieve
// median prices. Methods are goroutine safe.
type MarketToExchangePrices struct {
	sync.Mutex                                         // lock
	marketToExchangePrices map[uint32]*ExchangeToPrice // {k: market id, v: exchange prices}
	// maxPriceAge is the maximum age of a price before it is considered too stale to be used.
	// Prices older than this age will not be used to calculate the median price.
	maxPriceAge time.Duration
}

// NewMarketToExchangePrices creates a new MarketToExchangePrices.
func NewMarketToExchangePrices(maxPriceAge time.Duration) *MarketToExchangePrices {
	return &MarketToExchangePrices{
		marketToExchangePrices: make(map[uint32]*ExchangeToPrice),
		maxPriceAge:            maxPriceAge,
	}
}

// UpdatePrices updates market prices given a list of price updates. Prices are
// only updated if the timestamp on the updates are greater than the timestamp
// on existing prices.
func (mte *MarketToExchangePrices) UpdatePrices(
	updates []*api.MarketPriceUpdate) {
	mte.Lock()
	fmt.Println("Locked mutex in UpdatePrices")

	defer func() {
		telemetry.ModuleMeasureSince(
			types.ModuleName,
			time.Now(),
			metrics.UpdateCachePrices,
			metrics.Latency,
		)
		fmt.Println("Unlocking mutex in UpdatePrices")
		mte.Unlock()
	}()
	for _, marketPriceUpdate := range updates {
		marketId := marketPriceUpdate.MarketId
		exchangeToPrices, ok := mte.marketToExchangePrices[marketId]
		if !ok {
			exchangeToPrices = NewExchangeToPrice(marketId)
			mte.marketToExchangePrices[marketId] = exchangeToPrices
		}
		exchangeToPrices.UpdatePrices(marketPriceUpdate.ExchangePrices)
	}
}

// GetValidMedianPrices returns median prices for multiple markets.
// Specifically, it returns a map where the key is the market ID and the value
// is the median price for the market. It only returns "valid" prices where
// a price is valid iff
// 1) the last update time is within a predefined threshold away from the given
// read time.
func (mte *MarketToExchangePrices) GetValidMedianPrices(
	logger log.Logger,
	marketParams []types.MarketParam,
	readTime time.Time,
) map[uint32]uint64 {
	cutoffTime := readTime.Add(-mte.maxPriceAge)
	marketIdToMedianPrice := make(map[uint32]uint64)

	mte.Lock()
	fmt.Println("Locked mutex in GetValidMedianPrices")

	defer func() {
		telemetry.ModuleMeasureSince(
			types.ModuleName,
			time.Now(),
			metrics.GetMedianPrices,
			metrics.Latency,
		)
		fmt.Println("Unlocking mutex in GetValidMedianPrices")
		mte.Unlock()
	}()
	for _, marketParam := range marketParams {
		marketId := marketParam.Id
		exchangeToPrice, ok := mte.marketToExchangePrices[marketId]
		if !ok {
			// No market price info yet, skip this market.
			logger.Warn("No market price info", metrics.MarketId, marketId)
			logger.Warn("State of marketToExchangePrices", mte.marketToExchangePrices)
			telemetry.IncrCounterWithLabels(
				[]string{
					metrics.PricefeedServer,
					metrics.NoMarketPrice,
					metrics.Count,
				},
				1,
				[]gometrics.Label{
					pricefeedmetrics.GetLabelForMarketId(marketId),
				},
			)
			continue
		}

		// GetValidPriceForMarket filters prices based on cutoff time.
		validPrices := exchangeToPrice.GetValidPrices(logger, cutoffTime)

		// Calculate the median. Returns an error if the input is empty.
		median, err := lib.Median(validPrices)
		if err != nil {
			logger.Error("No valid median price", metrics.MarketId, marketId, metrics.Error, err)
			telemetry.IncrCounterWithLabels(
				[]string{
					metrics.PricefeedServer,
					metrics.NoValidMedianPrice,
					metrics.Count,
				},
				1,
				[]gometrics.Label{
					pricefeedmetrics.GetLabelForMarketId(marketId),
				},
			)
			continue
		}
		marketIdToMedianPrice[marketId] = median
	}

	return marketIdToMedianPrice
}
