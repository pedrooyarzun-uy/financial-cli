package models

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRes struct {
	Message string `json:"message" binding:"required"`
	Auth    string `json:"auth" binding:"required"`
}
