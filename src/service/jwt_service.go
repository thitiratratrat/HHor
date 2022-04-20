package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(id string, role Role) string
	ValidateToken(token string) (*jwt.Token, error)
	GetClaims(authHeader string) jwt.MapClaims
}

type Role string

const (
	Student   Role = "Student"
	DormOwner Role = "DormOwner"
)

type authClaims struct {
	Id   string `json:"id"`
	Role Role   `json:"role"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func JWTServiceHandler() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "HHOR",
	}
}

func getSecretKey() string {
	return os.Getenv("SECRET_KEY")
}

func (service *jwtService) GenerateToken(id string, role Role) string {
	claims := &authClaims{
		id,
		role,
		jwt.StandardClaims{
			Issuer:   service.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(service.secretKey))

	if err != nil {
		panic(err)
	}

	return signedToken
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})

}

func (service *jwtService) GetClaims(authHeader string) jwt.MapClaims {
	const BEARER_SCHEMA = "Bearer "
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := service.ValidateToken(tokenString)

	return token.Claims.(jwt.MapClaims)
}
