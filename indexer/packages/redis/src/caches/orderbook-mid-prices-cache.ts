import Big from 'big.js';
import { Callback, RedisClient } from 'redis';

import { getOrderBookMidPrice } from './orderbook-levels-cache';
import {
  addOrderbookMidPricesScript,
  getOrderbookMedianPriceScript,
} from './scripts';

// Cache of orderbook prices for each clob pair
// Each price is cached for a 5 second window and in a ZSET
export const ORDERBOOK_MID_PRICES_CACHE_KEY_PREFIX: string = 'v4/orderbook_mid_prices/';

/**
 * Generates a cache key for a given ticker's orderbook mid price.
 * @param ticker The ticker symbol
 * @returns The cache key string
 */
function getOrderbookMidPriceCacheKey(ticker: string): string {
  return `${ORDERBOOK_MID_PRICES_CACHE_KEY_PREFIX}${ticker}`;
}

/**
 * Fetches and caches mid prices for multiple tickers.
 * @param client The Redis client
 * @param tickers An array of ticker symbols
 * @returns A promise that resolves when all prices are fetched and cached
 */
export async function fetchAndCacheOrderbookMidPrices(
  client: RedisClient,
  tickers: string[],
): Promise<void> {
  // Fetch midPrices and filter out undefined values
  const cacheKeyPricePairs = await Promise.all(
    tickers.map(async (ticker) => {
      const cacheKey = getOrderbookMidPriceCacheKey(ticker);
      const midPrice = await getOrderBookMidPrice(cacheKey, client);
      if (midPrice !== undefined) {
        return { cacheKey, midPrice };
      }
      return null; // Return null for undefined midPrice
    }),
  );

  // Filter out null values
  const validPairs = cacheKeyPricePairs.filter(
    (pair): pair is { cacheKey: string, midPrice: string } => pair !== null,
  );
  if (validPairs.length === 0) {
    // No valid midPrices to cache
    return;
  }

  const nowSeconds = Math.floor(Date.now() / 1000); // Current time in seconds
  // Extract cache keys and prices
  const priceCacheKeys = validPairs.map((pair) => pair.cacheKey);
  const priceValues = validPairs.map((pair) => pair.midPrice);

  return new Promise<void>((resolve, reject) => {
    client.evalsha(
      addOrderbookMidPricesScript.hash,
      priceCacheKeys.length,
      ...priceCacheKeys,
      ...priceValues,
      nowSeconds,
      (err: Error | null) => {
        if (err) {
          reject(err);
        } else {
          resolve();
        }
      },
    );
  });
}

/**
 * Retrieves the median price for a given ticker from the cache.
 * Uses a Lua script to fetch either the middle element (for odd number of prices)
 * or the two middle elements (for even number of prices) from a sorted set in Redis.
 * If two middle elements are returned, their average is calculated in JavaScript.
 * @param client The Redis client
 * @param ticker The ticker symbol
 * @returns A promise that resolves with the median price as a string, or null if not found
 */
export async function getMedianPrice(client: RedisClient, ticker: string): Promise<string | null> {
  let evalAsync: (
    marketCacheKey: string,
  ) => Promise<string[]> = (
    marketCacheKey,
  ) => {
    return new Promise((resolve, reject) => {
      const callback: Callback<string[]> = (
        err: Error | null,
        results: string[],
      ) => {
        if (err) {
          return reject(err);
        }
        return resolve(results);
      };

      client.evalsha(
        getOrderbookMedianPriceScript.hash,
        1,
        marketCacheKey,
        callback,
      );
    });
  };
  evalAsync = evalAsync.bind(client);

  const prices = await evalAsync(
    getOrderbookMidPriceCacheKey(ticker),
  );

  if (!prices || prices.length === 0) {
    return null;
  }

  if (prices.length === 1) {
    return Big(prices[0]).toFixed();
  }

  if (prices.length === 2) {
    const [price1, price2] = prices.map((price) => {
      return Big(price);
    });
    return price1.plus(price2).div(2).toFixed();
  }

  return null;
}
