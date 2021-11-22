package api

import (
	"majoominipos/middleware"
	"majoominipos/models"
	"majoominipos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ServiceProduct service.ProductsServices
}

func ProductApiProvider(p service.ProductsServices) ProductAPI {
	return ProductAPI{ServiceProduct: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	id, _ := middleware.ClaimsToken(c)
	id_merchant := int(id["Id"].(float64))
	products := p.ServiceProduct.FindAll(id_merchant)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var product models.Products
	err := c.BindJSON(&product)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	Id_Merchant, err := middleware.ClaimsToken(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	product.Id_Merchant = int(Id_Merchant["Id"].(float64))

	createProduct, e := p.ServiceProduct.Save(product)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": createProduct})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var UpdatedField models.Products
	err := c.BindJSON(&UpdatedField)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "format data is wrong"})
		return
	}

	update, err := p.ServiceProduct.Save(UpdatedField)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "format data is wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": update})
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	del := p.ServiceProduct.Delete(id)
	if del != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": del.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
