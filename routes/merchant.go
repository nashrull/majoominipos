package routes

import (
	"majoominipos/api"
	"majoominipos/repository"
	"majoominipos/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MerchantRoutes(r *gin.Engine, db *gorm.DB) {
	productRepo := repository.MerchantRepositoryProvider(db)
	productService := service.MerchantServiceProvider(productRepo)
	productApi := api.MerchantApiProvider(productService)

	merchantRoute := r.Group("merchant")
	merchantRoute.POST("", productApi.Create)
	merchantRoute.POST("login", productApi.Login)
}
