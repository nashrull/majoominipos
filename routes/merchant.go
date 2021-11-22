package routes

import (
	"majoominipos/api"
	"majoominipos/middleware"
	"majoominipos/repository"
	"majoominipos/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MerchantRoutes(r *gin.Engine, db *gorm.DB) {
	merchantRepo := repository.MerchantRepositoryProvider(db)
	merchantService := service.MerchantServiceProvider(merchantRepo)
	merchantApi := api.MerchantApiProvider(merchantService)

	merchantRoute := r.Group("merchant")
	merchantRoute.POST("", merchantApi.Create)
	merchantRoute.GET("", merchantApi.GetAll)
	merchantRoute.POST("login", merchantApi.Login)
	merchantRoute.Use(middleware.ValidateMerchant)
	merchantRoute.DELETE(":id", merchantApi.Delete)
	merchantRoute.PUT("", merchantApi.Update)
}
