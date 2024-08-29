package models

import "gorm.io/gorm"

type Transaction struct {
	ID                uint   `gorm:"primaryKey"`
	TransactionHash   string `gorm:"unique"`
	TransactionStatus int
	BlockHash         string
	BlockNumber       uint64
	From              string
	To                string
	ContractAddress   string
	LogsCount         int
	Input             string
	Value             string
}

// UserTransaction represents a mapping between a user and their requested transaction hashes.
type UserTransaction struct {
	gorm.Model
	Username string `gorm:"index"`
	TxHash   string
}
