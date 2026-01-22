package dto

import "time"

type AddTransactionReq struct {
	Amount   float64 `json:"amount" binding:"required"`
	Account  int     `json:"account" binding:"required"`
	Currency int     `json:"currency" binding:"required"`
	Type     int     `json:"type" binding:"required"`
	Category int     `json:"category"`
	Notes    string  `json:"notes" `
}

type AddTransactionRes struct {
	Message string `json:"message" binding:"required"`
}

type TotalByCategory struct {
	Category string  `json:"Category"`
	Total    float64 `json:"Total"`
	Color    string  `json:"Color"`
}

type TotalsByCategoryRes struct {
	Message string            `json:"message"`
	Totals  []TotalByCategory `json:"totals"`
}

type TransactionByDetail struct {
	Id          int       `db:"id"`
	Category    string    `db:"category"`
	Subcategory string    `db:"subcategory"`
	Amount      float64   `db:"amount"`
	Currency    string    `db:"currency"`
	Notes       string    `db:"notes"`
	Date        time.Time `db:"created_at"`
	Color       string    `db:"color"`
	Type        int       `db:"type"`
}

type GetTransactionsByDetailRes struct {
	Message      string                `json:"message"`
	Transactions []TransactionByDetail `json:"transactions"`
	TotalPages   int                   `json:"total_pages"`
}
