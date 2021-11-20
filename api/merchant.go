package api

import (
	"fmt"
	"majoominipos/models"
	"majoominipos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchatApi struct {
	merchantService service.MerchantService
}

func MerchantApiProvider(r service.MerchantService) MerchatApi {
	return MerchatApi{merchantService: r}
}

func (m *MerchatApi) Create(c *gin.Context) {
	var merchant models.Merchants
	err := c.BindJSON(&merchant)
	fmt.Println("API err , ", merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	newMerchant, err := m.merchantService.Registration(merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": newMerchant})
}

func (m *MerchatApi) Login(c *gin.Context) {
	var merchantData models.Merchants
	err := c.BindJSON(&merchantData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	login, err := m.merchantService.Login(merchantData)
	if err != nil || login == false {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": login})
}
