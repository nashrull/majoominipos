package routes

import (
	"majoominipos/api"
	"majoominipos/middleware"
	"majoominipos/repository"
	"majoominipos/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func OutletRoutes(r *gin.Engine, db *gorm.DB) {
	outletRepo := repository.OutletRepositoryProvider(db)
	OutletService := service.OutletServiceProvider(outletRepo)
	outletApi := api.OutletApiProvider(OutletService)

	outletRoute := r.Group("outlet")
	outletRoute.Use(middleware.ValidateMerchant)
	outletRoute.POST("", outletApi.Create)
}
