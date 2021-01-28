

```cpp
// mainLoop is responsible for transaction relaying, status querying,
// missing block fetching and reorg logic handling.
func (pool *TxPool) mainLoop() {
	defer pool.wg.Done()

	// missingBlock represents a missing block with a batch of
	// locally created transaction hashes it contains.
	type missingBlock struct {
		number   uint64
		failure  int
		txHashes []common.Hash
	}
	// includedBlock represents a mined block with a batch of
	// locally created transactions it contains.
	type includedBlock struct {
		number uint64
		txs    []*types.Transaction
	}
	// sentTransaction represents a transaction which has been sent
	// via relay backend and some query statistic for better request
	// scheduling.
	type sentTransaction struct {
		tx         *types.Transaction
		unknown    int
		queryUntil time.Time
	}
```


``cpp

    Send first message
    NOTE !!
    It's been tested that f2pool.com does not respond with json error to wrong
    access message (which is needed to autodetect stratum mode).
    IT DOES NOT RESPOND AT ALL !!
    Due to this we need to set a timeout (arbitrary set to 1 second) and
    if no response within that time consider the tentative login failed
    and switch to next stratum mode test
```

source: https://github.com/ethereum-mining/ethminer/blob/47ae149ee734a82c453d23ef46fed78da109b73e/libpoolprotocols/stratum/EthStratumClient.cpp#L650


```cpp
 EthStratumClient::processResponse(Json::Value& responseObject)

    // Store jsonrpc version to test against
    int _rpcVer = responseObject.isMember("jsonrpc") ? 2 : 1;

    bool _isNotification = false;  // Whether or not this message is a reply to previous request or
                                   // is a broadcast notification
    bool _isSuccess = false;       // Whether or not this is a succesful or failed response (implies
                                   // _isNotification = false)
    string _errReason = "";        // Content of the error reason
    string _method = "";           // The method of the notification (or request from pool)
    unsigned _id = 0;  // This SHOULD be the same id as the request it is responding to (known
                       // exception is ethermine.org using 999)
```
source: https://github.com/ethereum-mining/ethminer/blob/47ae149ee734a82c453d23ef46fed78da109b73e/libpoolprotocols/stratum/EthStratumClient.cpp#L721


```cpp
    Process response for each stratum flavour :
    ETHEREUMSTRATUM2 response to mining.hello
    ETHEREUMSTRATUM  response to mining.subscribe
    ETHPROXY         response to eth_submitLogin
    STRATUM          response to mining.subscribe


    check if the block number is in a valid range
    A year has ~31536000 seconds
    50 years have ~1576800000
    assuming a (very fast) blocktime of 10s:
    ==> in 50 years we get 157680000 (=0x9660180) blocks

    if (m_current.block > 0x9660180)

    // Only some eth-proxy compatible implementations carry the block number
    // namely ethermine.org
    m_current.block = -1;
    if (m_conn->StratumMode() == EthStratumClient::ETHPROXY &&
    jPrm.size() > prmIdx &&
    jPrm.get(Json::Value::ArrayIndex(prmIdx), "").asString().substr(0, 2) ==
    "0x")
```

source:https://github.com/ethereum-mining/ethminer/blob/47ae149ee734a82c453d23ef46fed78da109b73e/libpoolprotocols/stratum/EthStratumClient.cpp#L1344

```c

// coinmine.pl fix
int l = sShareTarget.length();
if (l < 66)
sShareTarget = "0x" + string(66 - l, '0') + sShareTarget.substr(2);


    // There is no stratum method to submit the hashrate so we use the rpc variant.
    // Note !!
    // id = 6 is also the id used by ethermine.org and nanopool to push new jobs
    // thus we will be in trouble if we want to check the result of hashrate submission
    // actually change the id from 6 to 9
    jReq["jsonrpc"] = "2.0";
    if (!m_conn->Workername().empty())
    jReq["worker"] = m_conn->Workername();
    jReq["method"] = "eth_submitHashrate";



// Transact invokes the (paid) contract method with params as input values.
func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	// Otherwise pack up the parameters and invoke the contract
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	// todo(rjl493456442) check the method is payable or not,
	// reject invalid transaction at the first place
	return c.transact(opts, &c.address, input)
}

// RawTransact initiates a transaction with the given raw calldata as the input.
// It's usually used to initiates transaction for invoking **Fallback** function.
func (c *BoundContract) RawTransact(opts *TransactOpts, calldata []byte) (*types.Transaction, error) {
	// todo(rjl493456442) check the method is payable or not,
	// reject invalid transaction at the first place
	return c.transact(opts, &c.address, calldata)
}
// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (c *BoundContract) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	// todo(rjl493456442) check the payable fallback or receive is defined
	// or not, reject invalid transaction at the first place
	return c.transact(opts, &c.address, nil)
}
// transact executes an actual transaction invocation, first deriving any missing
// authorization fields, and then scheduling the transaction for execution.
func (c *BoundContract) transact(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	var err error
```
