package api

import (
	"majoominipos/middleware"
	"majoominipos/models"
	"majoominipos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutletApi struct {
	OutletService service.OutletService
}

func OutletApiProvider(p service.OutletService) OutletApi {
	return OutletApi{OutletService: p}
}

func (o *OutletApi) Create(c *gin.Context) {
	var Outlet models.Outlets
	merchant, err := middleware.ClaimsToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	Outlet.Id_merchant = int(merchant["Id"].(float64))
	err = c.BindJSON(&Outlet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	newOutlet, err := o.OutletService.CreateOutlet(Outlet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "data": newOutlet})
}
