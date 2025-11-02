package models

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRes struct {
	Message string `json:"message" binding:"required"`
}
