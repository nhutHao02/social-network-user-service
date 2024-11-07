package model

type SignUpRequest struct {
	Email    string `json:"email" binding:"required" db:"Email"`
	Password string `json:"password" binding:"min=6" db:"Password"`
}

type SignUpReponse struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" db:"Email"`
	Password string `json:"password" binding:"min=6" db:"Password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}
