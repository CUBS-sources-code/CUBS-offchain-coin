package repository

import (
	"time"
)

type Transaction struct {
	ID        int `gorm:"primarykey"`
	Sender    string
	Receiver  string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionRepository interface {
	GetAll() ([]Transaction, error)
	GetById(int) (*Transaction, error)
	GetBySender(string) ([]Transaction, error)
	GetByReceiver(string) ([]Transaction, error)
	Create(string, string, int) (*Transaction, error)
}
