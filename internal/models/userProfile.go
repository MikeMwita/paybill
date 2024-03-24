package models

type UserProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive string `json:"isActive"`
}
