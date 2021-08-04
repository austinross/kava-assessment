package main

import (
	"container/heap"
	"fmt"
)

type Transaction struct {
	hash      string
	gas       float64
	feePerGas float64
	signature string
}

type Mempool []*Transaction

func (mp Mempool) Len() int { return len(mp) }

func (mp Mempool) Less(i, j int) bool {
	return (mp[i].gas * mp[i].feePerGas) < (mp[j].gas * mp[j].feePerGas) //Compare gas fees
}

func (mp Mempool) Swap(i, j int) { mp[i], mp[j] = mp[j], mp[i] }

func (mp *Mempool) Pop() interface{} {
	old := *mp
	n := len(old)
	transaction := old[n-1]
	old[n-1] = nil // avoid memory leak
	*mp = old[0 : n-1]
	return transaction
}

func (mp *Mempool) Push(x interface{}) {
	transaction := x.(Transaction)
	*mp = append(*mp, &transaction)
}

func (mp *Mempool) ToStrings() []string {
	var strings []string
	for range *mp {
		transaction := heap.Pop(mp).(*Transaction)
		strings = append(strings, fmt.Sprintf("TxHash=%s Gas=%d FeePerGas=%f Signature=%s",
			transaction.hash, int(transaction.gas), transaction.feePerGas, transaction.signature))
	}

	return strings
}
