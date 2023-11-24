package repository

type Transaction interface {
	Tx() interface{}
	TxBegin() (Transaction, error)
	TxCommit() error
	TxRollback() error
}
