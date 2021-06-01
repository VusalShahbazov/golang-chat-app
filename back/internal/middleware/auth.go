package middleware

import (
	"back/internal/models"
	"back/internal/response"
	"back/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(ctx *gin.Context)  {
	bearToken := ctx.GetHeader("Authorization")
	var token string
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		token =  strArr[1]
	}else{
		response.BadRequest(ctx , "Token not provided")
		ctx.Abort()
		return
	}

	claims, err := jwt.ValidateToken(token)
	if err != nil {
		//mode errors
		response.Unauthorized(ctx , err.Error())
		ctx.Abort()
		return
	}

	var user models.User
	err = models.DB.Where("id = ? " , claims.UserId).Find(&user).Error
	if err != nil {
		response.UnexpectedError(ctx , nil)
		ctx.Abort()
		return
	}
	ctx.Set("user" ,user)
	ctx.Next()
}
