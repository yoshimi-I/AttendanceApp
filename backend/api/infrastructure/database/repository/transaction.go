package repository

import (
	"gorm.io/gorm"
	"work-management-app/domain/repository"
)

type TransactionImpl struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewTransaction(db *gorm.DB) repository.Transaction {
	return &TransactionImpl{
		db: db,
		tx: nil,
	}
}

func (t *TransactionImpl) Tx() interface{} {
	return t.tx
}

func (t *TransactionImpl) TxBegin() (repository.Transaction, error) {
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}
	t.tx = tx
	return t, nil
}

func (t *TransactionImpl) TxCommit() error {
	if t.tx != nil {
		return t.tx.Commit().Error // トランザクションをコミット
	}
	return nil
}

func (t *TransactionImpl) TxRollback() error {
	if t.tx != nil {
		return t.tx.Rollback().Error // トランザクションをロールバック
	}
	return nil
}

// ConvertOrm トランザクションをormに整形
func ConvertOrm(tx repository.Transaction) *gorm.DB {
	_, ok := tx.Tx().(*gorm.DB)
	if !ok {
		return nil
	}
	return tx.Tx().(*gorm.DB)
}
