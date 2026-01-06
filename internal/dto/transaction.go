package dto

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
}

type TotalsByCategoryRes struct {
	Message string            `json:"message"`
	Totals  []TotalByCategory `json:"totals"`
}
