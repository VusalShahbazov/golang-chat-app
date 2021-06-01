package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"os"
	"reflect"
	"regexp"
	"strings"
)
type ValidationResponse  map[string][]string

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func GenerateValidationResponse(c *gin.Context, err error) {
	if e, ok := err.(validator.ValidationErrors); ok {
		r := ValidationResponse{}
		for _, v := range e {
			key := ToSnakeCase(v.Field())
			r[key] = append(r[key], v.Error())
		}
		ValidationError(c, &r)
		return
	}
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		ValidationError(c, gin.H{
			ToSnakeCase(e.Field): []string{fmt.Sprintf("Error %s", e.Error())},
		})
		return
	}
	if os.Getenv("ENV") == "production" {
		ErrorResponse(c, gin.H{"error": "Bad request", "type": reflect.TypeOf(err).String()}, 400)
	} else {
		ErrorResponse(c, gin.H{"error": err.Error(), "type": reflect.TypeOf(err).String()}, 400)
	}
}

func ValidationErrorByKey(c *gin.Context , key string , value []string)  {
	ValidationError(c , gin.H{key : value})
}
