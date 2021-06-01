package requests

type Register struct {
	LastName string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}