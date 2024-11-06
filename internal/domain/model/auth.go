package model

type SignUpRequest struct {
	Email    string `json:"email" validator:"require" db:"Email"`
	Password string `json:"password" validator:"min=6" db:"Password"`
}

type SignUpReponse struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}
