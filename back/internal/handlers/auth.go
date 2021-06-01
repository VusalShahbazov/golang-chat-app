package handlers

import (
	"back/internal/models"
	"back/internal/requests"
	"back/internal/response"
	"back/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(context *gin.Context) {
	var json requests.Login
	err := context.ShouldBindJSON(&json)
	if err != nil {
		response.GenerateValidationResponse(context, err)
		return
	}
	var user models.User
	err = models.DB.Where("email = ? ", json.Email).Find(&user).Error
	if err != nil {
		response.BadRequest(context, "Invalid Login or Password")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		response.BadRequest(context, "Invalid Login or Password")
		return
	}
	responseWithToken(context, &user)
}
func Register(context *gin.Context) {
	var json requests.Register

	err := context.ShouldBindJSON(&json)
	if err != nil {
		response.GenerateValidationResponse(context, err)
		return
	}

	var res struct {
		exists bool
	}

	err = models.DB.Raw("select exists(select 1 from users where email = ? ) as found ", json.Email).Scan(&res.exists).Error

	if err != nil {
		response.InternalServerError(context, nil)
		return
	}

	if res.exists {
		response.ValidationErrorByKey(context, "email", []string{"Email Must be unique"})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10)

	user := models.User{
		LastName:  json.LastName,
		FirstName: json.FirstName,
		Email:     json.Email,
		Password:  string(password),
	}

	models.DB.Create(&user)

	responseWithToken(context, &user)
}
func Logout(context *gin.Context) {
	_, exists := context.Get("user")
	if exists {
		context.Set("user", nil)
		response.SuccessResponse(context, gin.H{"message": "Successful logged out"})
		return
	}
	response.UnexpectedError(context, nil)
}
func Me(context *gin.Context) {
	get, exists := context.Get("user")
	if exists {
		response.SuccessResponse(context, gin.H{"user": get})
		return
	}
	response.UnexpectedError(context, nil)
}

func responseWithToken(ctx *gin.Context, user *models.User) {
	token, err := jwt.GenerateTokenFromUser(user)
	if err != nil {
		response.ErrorResponse(ctx, gin.H{"message": err.Error()}, 400)
		return
	}
	response.SuccessResponse(ctx, gin.H{"token": token, "user": user})
}
