package model

import "time"

type Transaction struct {
	ID        uint      `json:"ID"`
	UserID    uint      `json:"user_id"`
	BookID    uint      `json:"book_id"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type TransactionReq struct {
	BookID uint    `json:"book_id"`
	Amount float32 `json:"amount"`
}

type TransactionCancelReq struct {
	TransactionID uint `json:"transaction_id"`
}
