	/* //approach 1 	*/
	transactions := self.eth.TxPool().GetTransactions()
	sort.Sort(types.TxByNonce{transactions})


	//approach 2 
	transactions := self.eth.TxPool().GetTransactions()
	sort.Sort(types.TxByPriceAndNonce{transactions})

	/* // approach 3 	*/

	// commit transactions for this run.
	txPerOwner := make(map[common.Address]types.Transactions)
	// Sort transactions by owner
	for _, tx := range self.eth.TxPool().GetTransactions() {
		from, _ := tx.From() // we can ignore the sender error
		txPerOwner[from] = append(txPerOwner[from], tx)
	}
	var (
		singleTxOwner types.Transactions
		multiTxOwner  types.Transactions
	)
	// Categorise transactions by
	// 1. 1 owner tx per block
	// 2. multi txs owner per block
	for _, txs := range txPerOwner {
		if len(txs) == 1 {
			singleTxOwner = append(singleTxOwner, txs[0])
		} else {
			multiTxOwner = append(multiTxOwner, txs...)
		}
	}
	sort.Sort(types.TxByPrice{singleTxOwner})
	sort.Sort(types.TxByNonce{multiTxOwner})
	transactions := append(singleTxOwner, multiTxOwner...)
