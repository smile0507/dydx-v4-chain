// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	indexer_manager "github.com/dydxprotocol/v4-chain/protocol/indexer/indexer_manager"
	clobtypes "github.com/dydxprotocol/v4-chain/protocol/x/clob/types"

	log "cosmossdk.io/log"

	mock "github.com/stretchr/testify/mock"

	subaccountstypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// MemClobKeeper is an autogenerated mock type for the MemClobKeeper type
type MemClobKeeper struct {
	mock.Mock
}

// AddOrderToOrderbookSubaccountUpdatesCheck provides a mock function with given fields: ctx, subaccountId, order
func (_m *MemClobKeeper) AddOrderToOrderbookSubaccountUpdatesCheck(ctx types.Context, subaccountId subaccountstypes.SubaccountId, order clobtypes.PendingOpenOrder) subaccountstypes.UpdateResult {
	ret := _m.Called(ctx, subaccountId, order)

	if len(ret) == 0 {
		panic("no return value specified for AddOrderToOrderbookSubaccountUpdatesCheck")
	}

	var r0 subaccountstypes.UpdateResult
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, clobtypes.PendingOpenOrder) subaccountstypes.UpdateResult); ok {
		r0 = rf(ctx, subaccountId, order)
	} else {
		r0 = ret.Get(0).(subaccountstypes.UpdateResult)
	}

	return r0
}

// AddPreexistingStatefulOrder provides a mock function with given fields: ctx, order, memclob
func (_m *MemClobKeeper) AddPreexistingStatefulOrder(ctx types.Context, order *clobtypes.Order, memclob clobtypes.MemClob) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, order, memclob)

	if len(ret) == 0 {
		panic("no return value specified for AddPreexistingStatefulOrder")
	}

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 *clobtypes.OffchainUpdates
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.Order, clobtypes.MemClob) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, order, memclob)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.Order, clobtypes.MemClob) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, order, memclob)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.Order, clobtypes.MemClob) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, order, memclob)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.Order, clobtypes.MemClob) *clobtypes.OffchainUpdates); ok {
		r2 = rf(ctx, order, memclob)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(3).(func(types.Context, *clobtypes.Order, clobtypes.MemClob) error); ok {
		r3 = rf(ctx, order, memclob)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// CanDeleverageSubaccount provides a mock function with given fields: ctx, subaccountId, perpetualId
func (_m *MemClobKeeper) CanDeleverageSubaccount(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32) (bool, bool, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId)

	if len(ret) == 0 {
		panic("no return value specified for CanDeleverageSubaccount")
	}

	var r0 bool
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) (bool, bool, error)); ok {
		return rf(ctx, subaccountId, perpetualId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) bool); ok {
		r0 = rf(ctx, subaccountId, perpetualId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32) bool); ok {
		r1 = rf(ctx, subaccountId, perpetualId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(types.Context, subaccountstypes.SubaccountId, uint32) error); ok {
		r2 = rf(ctx, subaccountId, perpetualId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CancelShortTermOrder provides a mock function with given fields: ctx, msgCancelOrder
func (_m *MemClobKeeper) CancelShortTermOrder(ctx types.Context, msgCancelOrder *clobtypes.MsgCancelOrder) error {
	ret := _m.Called(ctx, msgCancelOrder)

	if len(ret) == 0 {
		panic("no return value specified for CancelShortTermOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r0 = rf(ctx, msgCancelOrder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIndexerEventManager provides a mock function with given fields:
func (_m *MemClobKeeper) GetIndexerEventManager() indexer_manager.IndexerEventManager {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIndexerEventManager")
	}

	var r0 indexer_manager.IndexerEventManager
	if rf, ok := ret.Get(0).(func() indexer_manager.IndexerEventManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(indexer_manager.IndexerEventManager)
		}
	}

	return r0
}

// GetLongTermOrderPlacement provides a mock function with given fields: ctx, orderId
func (_m *MemClobKeeper) GetLongTermOrderPlacement(ctx types.Context, orderId clobtypes.OrderId) (clobtypes.LongTermOrderPlacement, bool) {
	ret := _m.Called(ctx, orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetLongTermOrderPlacement")
	}

	var r0 clobtypes.LongTermOrderPlacement
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (clobtypes.LongTermOrderPlacement, bool)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) clobtypes.LongTermOrderPlacement); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(clobtypes.LongTermOrderPlacement)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) bool); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetOrderFillAmount provides a mock function with given fields: ctx, orderId
func (_m *MemClobKeeper) GetOrderFillAmount(ctx types.Context, orderId clobtypes.OrderId) (bool, subaccountstypes.BaseQuantums, uint32) {
	ret := _m.Called(ctx, orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderFillAmount")
	}

	var r0 bool
	var r1 subaccountstypes.BaseQuantums
	var r2 uint32
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (bool, subaccountstypes.BaseQuantums, uint32)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) bool); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) subaccountstypes.BaseQuantums); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(2).(func(types.Context, clobtypes.OrderId) uint32); ok {
		r2 = rf(ctx, orderId)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	return r0, r1, r2
}

// GetStatePosition provides a mock function with given fields: ctx, subaccountId, clobPairId
func (_m *MemClobKeeper) GetStatePosition(ctx types.Context, subaccountId subaccountstypes.SubaccountId, clobPairId clobtypes.ClobPairId) *big.Int {
	ret := _m.Called(ctx, subaccountId, clobPairId)

	if len(ret) == 0 {
		panic("no return value specified for GetStatePosition")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, clobtypes.ClobPairId) *big.Int); ok {
		r0 = rf(ctx, subaccountId, clobPairId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// IsLiquidatable provides a mock function with given fields: ctx, subaccountId
func (_m *MemClobKeeper) IsLiquidatable(ctx types.Context, subaccountId subaccountstypes.SubaccountId) (bool, error) {
	ret := _m.Called(ctx, subaccountId)

	if len(ret) == 0 {
		panic("no return value specified for IsLiquidatable")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) (bool, error)); ok {
		return rf(ctx, subaccountId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) bool); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId) error); ok {
		r1 = rf(ctx, subaccountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logger provides a mock function with given fields: ctx
func (_m *MemClobKeeper) Logger(ctx types.Context) log.Logger {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Logger")
	}

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(types.Context) log.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// OffsetSubaccountPerpetualPosition provides a mock function with given fields: ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal, isFinalSettlement
func (_m *MemClobKeeper) OffsetSubaccountPerpetualPosition(ctx types.Context, liquidatedSubaccountId subaccountstypes.SubaccountId, perpetualId uint32, deltaQuantumsTotal *big.Int, isFinalSettlement bool) ([]clobtypes.MatchPerpetualDeleveraging_Fill, *big.Int) {
	ret := _m.Called(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal, isFinalSettlement)

	if len(ret) == 0 {
		panic("no return value specified for OffsetSubaccountPerpetualPosition")
	}

	var r0 []clobtypes.MatchPerpetualDeleveraging_Fill
	var r1 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int, bool) ([]clobtypes.MatchPerpetualDeleveraging_Fill, *big.Int)); ok {
		return rf(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal, isFinalSettlement)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int, bool) []clobtypes.MatchPerpetualDeleveraging_Fill); ok {
		r0 = rf(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal, isFinalSettlement)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.MatchPerpetualDeleveraging_Fill)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int, bool) *big.Int); ok {
		r1 = rf(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal, isFinalSettlement)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	return r0, r1
}

// ProcessSingleMatch provides a mock function with given fields: ctx, matchWithOrders
func (_m *MemClobKeeper) ProcessSingleMatch(ctx types.Context, matchWithOrders *clobtypes.MatchWithOrders) (bool, subaccountstypes.UpdateResult, subaccountstypes.UpdateResult, error) {
	ret := _m.Called(ctx, matchWithOrders)

	if len(ret) == 0 {
		panic("no return value specified for ProcessSingleMatch")
	}

	var r0 bool
	var r1 subaccountstypes.UpdateResult
	var r2 subaccountstypes.UpdateResult
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MatchWithOrders) (bool, subaccountstypes.UpdateResult, subaccountstypes.UpdateResult, error)); ok {
		return rf(ctx, matchWithOrders)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MatchWithOrders) bool); ok {
		r0 = rf(ctx, matchWithOrders)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MatchWithOrders) subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, matchWithOrders)
	} else {
		r1 = ret.Get(1).(subaccountstypes.UpdateResult)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MatchWithOrders) subaccountstypes.UpdateResult); ok {
		r2 = rf(ctx, matchWithOrders)
	} else {
		r2 = ret.Get(2).(subaccountstypes.UpdateResult)
	}

	if rf, ok := ret.Get(3).(func(types.Context, *clobtypes.MatchWithOrders) error); ok {
		r3 = rf(ctx, matchWithOrders)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// ReplayPlaceOrder provides a mock function with given fields: ctx, msg
func (_m *MemClobKeeper) ReplayPlaceOrder(ctx types.Context, msg *clobtypes.MsgPlaceOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for ReplayPlaceOrder")
	}

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 *clobtypes.OffchainUpdates
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MsgPlaceOrder) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MsgPlaceOrder) *clobtypes.OffchainUpdates); ok {
		r2 = rf(ctx, msg)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(3).(func(types.Context, *clobtypes.MsgPlaceOrder) error); ok {
		r3 = rf(ctx, msg)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// SendOrderbookFillUpdate provides a mock function with given fields: ctx, orderbookFills
func (_m *MemClobKeeper) SendOrderbookFillUpdate(ctx types.Context, orderbookFill clobtypes.StreamOrderbookFill) {
	_m.Called(ctx, orderbookFill)
}

// SendOrderbookUpdates provides a mock function with given fields: ctx, offchainUpdates
func (_m *MemClobKeeper) SendOrderbookUpdates(ctx types.Context, offchainUpdates *clobtypes.OffchainUpdates) {
	_m.Called(ctx, offchainUpdates)
}

// SendTakerOrderStatus provides a mock function with given fields: ctx, takerOrder
func (_m *MemClobKeeper) SendTakerOrderStatus(ctx types.Context, takerOrder clobtypes.StreamTakerOrder) {
	_m.Called(ctx, takerOrder)
}

// SetLongTermOrderPlacement provides a mock function with given fields: ctx, order, blockHeight
func (_m *MemClobKeeper) SetLongTermOrderPlacement(ctx types.Context, order clobtypes.Order, blockHeight uint32) {
	_m.Called(ctx, order, blockHeight)
}

// ValidateSubaccountEquityTierLimitForStatefulOrder provides a mock function with given fields: ctx, order
func (_m *MemClobKeeper) ValidateSubaccountEquityTierLimitForStatefulOrder(ctx types.Context, order clobtypes.Order) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for ValidateSubaccountEquityTierLimitForStatefulOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMemClobKeeper creates a new instance of MemClobKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMemClobKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MemClobKeeper {
	mock := &MemClobKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
