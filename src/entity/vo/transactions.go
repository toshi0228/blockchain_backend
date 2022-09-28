package vo

import "strings"

type Transactions string

func NewTransactions(sliceString []string) Transactions {

	txPoolStr := strings.Join(sliceString, ",")
	return Transactions(txPoolStr)

}

func (tx Transactions) Value() string {
	return string(tx)
}
