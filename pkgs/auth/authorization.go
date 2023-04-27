package auth

import (
	"strings"
	"math/rand"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	UserID	primitive.ObjectID
	jwt.StandardClaims
}

func NewClaims(userid primitive.ObjectID) Claims {
	return Claims{userid, jwt.StandardClaims{}}
}

var JWTKey []byte

func GenerateJWTKey() {
	JWTKey = make([]byte, 32)
	rand.Read(JWTKey)
}

func GenerateTokenStringFromClaims(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetTokenString(c *gin.Context) string {
	authString := c.Request.Header.Get("Authorization")
	tokenStrings := strings.Split(authString, "Bearer ")
	if len(tokenStrings) != 2 || tokenStrings[0] != "" {
		return ""
	}

	return tokenStrings[1]
}

func ParseTokenString(tokenString string) (Claims, error)  {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Claims{}, errors.New("signature invalid")
		} else {
			return Claims{}, errors.New("can't parse token string")
		}
	}

	if !token.Valid {
		return Claims{}, errors.New("token invalid")
	}

	return *claims, nil
}