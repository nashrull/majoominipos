package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"majoominipos/product"
	"majoominipos/routes"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/majoopos?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return db
}

func initProductAPI(db *gorm.DB) product.ProductAPI {
	productRepository := product.ProvideProductRepostiory(db)
	productService := product.ProvideProductService(productRepository)
	productAPI := product.ProvideProductAPI(productService)
	return productAPI
}

func main() {
	db := initDB()
	defer db.Close()

	// productAPI := initProductAPI(db)

	r := gin.Default()
	routes.ProductsRoute(r, db)
	routes.MerchantRoutes(r, db)
	routes.OutletRoutes(r, db)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
