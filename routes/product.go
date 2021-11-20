package routes

import (
	"majoominipos/api"
	"majoominipos/repository"
	"majoominipos/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ProductsRoute(r *gin.Engine, db *gorm.DB) {
	productRepository := repository.ProductRepositoryProvider(db)
	productService := service.ProviderProductService(productRepository)
	productAPI := api.ProductApiProvider(productService)

	product := r.Group("product/")
	product.GET("", productAPI.FindAll)
	product.POST("", productAPI.Create)

}
