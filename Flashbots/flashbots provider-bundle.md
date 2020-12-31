## ethers-provider-flashbots-bundle

Contains the `FlashbotsBundleProvider` ethers.js provider to provide high-level access to eth_sendBundle rpc endpoint.

Flashbots-enabled relays and miners will expose a single jsonrpc endpoint: `eth_sendBundle`. Since this is a brand-new, non-standard endpoint, ethers.js and other libraries do not natively support these requests (like `getTransactionCount`). In order to interact with `eth_sendBundle`, you will also need access to another full-featured endpoint for nonce-calculation, gas estimation, and transaction status.

This library is not a fully functional ethers.js implementation, just a simple provider class, designed to interact with your existing ethers.js v5 module.

You can pass in a generic ethers.js provider to the flashbots provider in the constructor:

```
const NETWORK_INFO = {chainId: 1, ensAddress: '', name: 'mainnet'}

// Standard json rpc provider directly from ethers.js
const provider = new providers.JsonRpcProvider({url: ETHEREUM_URL}, NETWORK_INFO)

// flashbots provider requires passing in a standard provider
const flashbotsProvider = new FlashbotsBundleProvider(provider, {url: FLASHBOTS_RELAY_URL}, NETWORK_INFO)
```

The flashbotsProvider provides the sendBundle function:

```
flashbotsProvider.sendBundle(bundledTransactions: Array<FlashbotsBundleTransaction | FlashbotsBundleRawTransaction>, targetBlockNumber: number)
    => Promise<FlashbotsTransactionResponse>
```

## Example

```
// Using the map below ships two different bundles, targeting the next two blocks
const blockNumber = await provider.getBlockNumber()
const bundlePromises = _.map([blockNumber + 1, blockNumber + 2], targetBlockNumber =>
  this.flashbotsProvider.sendBundle(
    [
      {
        signedTransaction: SIGNED_ORACLE_UPDATE_FROM_PENDING_POOL
      },
      {
        signer: wallet,
        transaction: transaction
      }
    ],
    targetBlockNumber
  )
)
```

## bundledTransactions

A Flashbots bundle consists of one or more transactions in strict order to be relayed to the miner directly. While the miner requires signed transaction, `sendBundle()` can receive a mix of pre-signed transaction and `TransactionRequest` + `Signer` (wallet) objects

These bundles can pay the miner either via gas fees _OR_ via `block.coinbase.transfer(minerReward)`

## targetBlockNumber

The only block number for which the bundle is to be considered valid. If you would like more than one block to be targeted, submit multiple rpc calls targeting each specific block. This value should be higher than the value of getBlockNumber(). Submitting a bundle with a target block number of the current block, or earlier, is a no-op.

## FlashbotsTransactionResponse

A high-level object which contains metadata available at transaction submission time, as well as the following functions which can wait, track, and simulate the bundle's behavior

- receipts() - Returns promise of an array of transaction receipts corresponding to the transaction hashes that were relayed as part of the bundle. Will not wait for block to be mined; could return incomplete information
- wait() - Returns a promise which will wait for target block number to be reched _OR_ one of the transactions to become invalid due to nonce-issues (including, but not limited to, one of the transactions from you bundle being included too early). Returns the wait resolution as a status enum
- simulate() - Returns a promise of the transaction simulation, once the proper block height has been reached. Use this function to troubleshoot failing bundles and verify miner profitability
