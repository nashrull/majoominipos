package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateMerchant(c *gin.Context) {
	token := strings.Join(c.Request.Header["Token"], ",")
	validToken, err := JWTAuthService().ValidateToken(token)
	if err != nil {
		fmt.Println("Error token =>", err.Error())
		c.Abort()
		return
	}
	if validToken.Valid == false {
		c.Abort()
		return
	}

}
