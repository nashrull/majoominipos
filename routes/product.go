package routes

import (
	"majoominipos/api"
	"majoominipos/middleware"
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
	product.Use(middleware.ValidateMerchant)
	product.GET("", productAPI.FindAll)
	product.POST("", productAPI.Create)
	product.PUT("", productAPI.Update)
	product.DELETE(":id", productAPI.Delete)
}
