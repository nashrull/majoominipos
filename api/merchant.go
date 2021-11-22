package api

import (
	"majoominipos/middleware"
	"majoominipos/models"
	"majoominipos/service"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{"msg ": err.Error()})
		return
	}
	login, err := m.merchantService.Login(merchantData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// generate token
	token, err := middleware.JWTAuthService().GenerateToken(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "token": token})
}

func (m *MerchatApi) Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "success"})
}

func (m *MerchatApi) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	del := m.merchantService.Delete(id)
	if del != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": del.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (m *MerchatApi) GetAll(c *gin.Context) {
	var listMerchant []models.Merchants
	listMerchant = m.merchantService.GetAll()
	c.JSON(http.StatusOK, gin.H{"data": listMerchant})
}

func (m *MerchatApi) Update(c *gin.Context) {
	var merchant models.Merchants
	err := c.BindJSON(&merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	updatedeData, err := m.merchantService.Update(merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedeData})
}
