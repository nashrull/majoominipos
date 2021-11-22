package middleware

import (
	"fmt"
	"majoominipos/models"
	"strings"

	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//jwt service
type JWTService interface {
	GenerateToken(data models.Merchants) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Id       int
	Username string
	Status   string
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "askhdjadwj",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(data models.Merchants) (string, error) {
	claims := &authCustomClaims{
		data.Id,
		data.Username,
		data.Status,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	return t, err
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

func ClaimsToken(c *gin.Context) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	tokenString := strings.Join(c.Request.Header["Token"], ",")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return result, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("Claims >", claims)

	if ok && token.Valid {
		result["Id"] = claims["Id"]
		result["Status"] = claims["Status"]
		result["Username"] = claims["Username"]
		return result, nil
	}
	return result, nil
}
