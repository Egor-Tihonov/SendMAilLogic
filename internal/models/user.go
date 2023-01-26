package models

type EmailRequest struct {
	Email string `json:"email"`
}

type UserMail struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
