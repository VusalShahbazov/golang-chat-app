package helpers

import (
	"back/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetOrDefault(value , default_value string) string {
	trim_value := strings.Trim(value , " ")
	if trim_value != ""  {
		return trim_value
	}
	return default_value
}


func GetUser(ctx *gin.Context) (*models.User, error) {
	stUser, ok := ctx.Get("user")
	if !ok {
		return nil, errors.New("User not exists")
	}
	user, ok := stUser.(models.User)
	if !ok {
		return nil, errors.New("User key is not valid")
	}
	return &user, nil
}