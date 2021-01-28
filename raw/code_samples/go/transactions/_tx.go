type TxByPrice struct{ Transactions }

func (s TxByPrice) Less(i, j int) bool {
	return s.Transactions[i].data.Price.Cmp(s.Transactions[j].data.Price) > 0
}

type TxByPriceAndNonce struct{ Transactions }

func (s TxByPriceAndNonce) Less(i, j int) bool {
	// we can ignore the error here. Sorting shouldn't care about validness
	ifrom, _ := s.Transactions[i].From()
	jfrom, _ := s.Transactions[j].From()
	// favour nonce if they are from the same recipient
	if ifrom == jfrom {
		return s.Transactions[i].data.AccountNonce < s.Transactions[j].data.AccountNonce
	}
	return s.Transactions[i].data.Price.Cmp(s.Transactions[j].data.Price) > 0
}
