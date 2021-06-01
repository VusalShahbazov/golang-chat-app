package models


type User struct {
	Model
	LastName string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email string `json:"email"`
	Password string `json:"-"`
	Avatar *string `json:"avatar"`
}
