
package jwt

import (
	"back/internal/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId uint64 `json:"user_id"`
}

func getTtl() time.Duration {
	return 15 * time.Hour
}

func getSecretKey() string {
	return os.Getenv("JWT_SECRET")
}

func GenerateTokenFromUser(user *models.User) (string, error) {
	claims := &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(getTtl()).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getSecretKey()))
}

func ValidateToken(signedToken string) (claims *TokenClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getSecretKey()), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return

}

func GetUserFromToken(token string) (*models.User , error)  {
	claims , err := ValidateToken(token)
	if err != nil  {
		return  nil, err
	}
	var user models.User
	err = models.DB.Where("id = ? " , claims.UserId).Find(&user).Error
	if err != nil {
		return  nil, err

	}


	return  &user , nil
}