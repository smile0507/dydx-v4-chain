import { stats, STATS_FUNCTION_NAME } from '@dydxprotocol-indexer/base';
import {
  AssetCreateEventV1,
  IndexerTendermintBlock,
  IndexerTendermintEvent,
  Timestamp,
} from '@dydxprotocol-indexer/v4-protos';
import {
  AssetColumns,
  AssetFromDatabase,
  AssetTable,
  dbHelpers,
  MarketTable,
  Ordering,
  testConstants,
  marketRefresher,
  BlockTable,
  TendermintEventTable,
  assetRefresher,
} from '@dydxprotocol-indexer/postgres';
import { KafkaMessage } from 'kafkajs';
import { createKafkaMessage } from '@dydxprotocol-indexer/kafka';
import { onMessage } from '../../src/lib/on-message';
import { DydxIndexerSubtypes } from '../../src/lib/types';
import {
  binaryToBase64String,
  createIndexerTendermintBlock,
  createIndexerTendermintEvent,
} from '../helpers/indexer-proto-helpers';
import { AssetCreationHandler } from '../../src/handlers/asset-handler';
import {
  defaultAssetCreateEvent,
  defaultHeight,
  defaultPreviousHeight,
  defaultTime,
  defaultTxHash,
} from '../helpers/constants';
import { updateBlockCache } from '../../src/caches/block-cache';

describe('assetHandler', () => {
  beforeAll(async () => {
    await dbHelpers.migrate();
    jest.spyOn(stats, 'increment');
    jest.spyOn(stats, 'timing');
    jest.spyOn(stats, 'gauge');
  });

  beforeEach(async () => {
    await Promise.all([
      BlockTable.create(testConstants.defaultBlock),
      BlockTable.create(testConstants.defaultBlock2),
    ]);
    await Promise.all([
      TendermintEventTable.create(testConstants.defaultTendermintEvent),
      TendermintEventTable.create(testConstants.defaultTendermintEvent2),
      TendermintEventTable.create(testConstants.defaultTendermintEvent3),
    ]);
    updateBlockCache(defaultPreviousHeight);
  });

  afterEach(async () => {
    await dbHelpers.clearData();
    jest.clearAllMocks();
  });

  afterAll(async () => {
    await dbHelpers.teardown();
    jest.resetAllMocks();
  });

  describe('getParallelizationIds', () => {
    it('returns the correct parallelization ids', () => {
      const transactionIndex: number = 0;
      const eventIndex: number = 0;

      const indexerTendermintEvent: IndexerTendermintEvent = createIndexerTendermintEvent(
        DydxIndexerSubtypes.ASSET,
        binaryToBase64String(
          AssetCreateEventV1.encode(defaultAssetCreateEvent).finish(),
        ),
        transactionIndex,
        eventIndex,
      );
      const block: IndexerTendermintBlock = createIndexerTendermintBlock(
        0,
        defaultTime,
        [indexerTendermintEvent],
        [defaultTxHash],
      );

      const handler: AssetCreationHandler = new AssetCreationHandler(
        block,
        indexerTendermintEvent,
        0,
        defaultAssetCreateEvent,
      );

      expect(handler.getParallelizationIds()).toEqual([]);
    });
  });

  it('fails when market doesnt exist for asset', async () => {
    const transactionIndex: number = 0;
    const kafkaMessage: KafkaMessage = createKafkaMessageFromAssetEvent({
      assetEvent: defaultAssetCreateEvent,
      transactionIndex,
      height: defaultHeight,
      time: defaultTime,
      txHash: defaultTxHash,
    });

    const message: string = 'Unable to find market with id: 0';
    await expect(onMessage(kafkaMessage)).rejects.toThrowError(
      new Error(message),
    );
  });

  it('creates new asset', async () => {
    await MarketTable.create(testConstants.defaultMarket);
    await marketRefresher.updateMarkets();
    const transactionIndex: number = 0;

    const assetEvent: AssetCreateEventV1 = defaultAssetCreateEvent;
    const kafkaMessage: KafkaMessage = createKafkaMessageFromAssetEvent({
      assetEvent,
      transactionIndex,
      height: defaultHeight,
      time: defaultTime,
      txHash: defaultTxHash,
    });
    // Confirm there is no existing asset to or from the sender subaccount
    await expectNoExistingAssets();

    await onMessage(kafkaMessage);

    const newAssets: AssetFromDatabase[] = await AssetTable.findAll(
      {},
      [], {
        orderBy: [[AssetColumns.id, Ordering.ASC]],
      });
    expect(newAssets.length).toEqual(1);
    expectAssetMatchesEvent(assetEvent, newAssets[0]);
    expectTimingStats();
    const asset: AssetFromDatabase = assetRefresher.getAssetFromId('0');
    expect(asset).toBeDefined();
  });
});

function expectTimingStats() {
  expectTimingStat('create_asset');
}

function expectTimingStat(fnName: string) {
  expect(stats.timing).toHaveBeenCalledWith(
    `ender.${STATS_FUNCTION_NAME}.timing`,
    expect.any(Number),
    {
      className: 'AssetCreationHandler',
      eventType: 'AssetCreateEvent',
      fnName,
    },
  );
}

function expectAssetMatchesEvent(
  event: AssetCreateEventV1,
  asset: AssetFromDatabase,
) {
  expect(asset.id).toEqual(event.id.toString());
  expect(asset.hasMarket).toEqual(event.hasMarket);
  expect(asset.marketId).toEqual(event.marketId);
  expect(asset.symbol).toEqual(event.symbol);
  expect(asset.atomicResolution).toEqual(event.atomicResolution);
}

function createKafkaMessageFromAssetEvent({
  assetEvent,
  transactionIndex,
  height,
  time,
  txHash,
}: {
  assetEvent: AssetCreateEventV1 | undefined,
  transactionIndex: number,
  height: number,
  time: Timestamp,
  txHash: string,
}) {
  const events: IndexerTendermintEvent[] = [];
  let eventIndex: number = 0;
  if (assetEvent !== undefined) {
    events.push(
      createIndexerTendermintEvent(
        DydxIndexerSubtypes.ASSET,
        binaryToBase64String(
          AssetCreateEventV1.encode(assetEvent).finish(),
        ),
        transactionIndex,
        eventIndex,
      ),
    );
    eventIndex += 1;
  }

  const block: IndexerTendermintBlock = createIndexerTendermintBlock(
    height,
    time,
    events,
    [txHash],
  );

  const binaryBlock: Uint8Array = IndexerTendermintBlock.encode(block).finish();
  return createKafkaMessage(Buffer.from(binaryBlock));
}

async function expectNoExistingAssets() {
  // Confirm there is no existing asset
  const assets: AssetFromDatabase[] = await AssetTable.findAll(
    {},
    [], {
      orderBy: [[AssetColumns.id, Ordering.ASC]],
    });

  expect(assets.length).toEqual(0);
}
