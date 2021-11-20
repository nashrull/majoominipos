package api

import (
	"majoominipos/models"
	"majoominipos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ServiceProduct service.ProductsServices
}

func ProductApiProvider(p service.ProductsServices) ProductAPI {
	return ProductAPI{ServiceProduct: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ServiceProduct.FindAll()
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var product models.Products
	err := c.BindJSON(&product)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	createProduct, e := p.ServiceProduct.Save(product)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": createProduct})
}
