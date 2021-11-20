package service

import (
	"majoominipos/models"
	"majoominipos/repository"
)

type ProductsServices struct {
	Repositori repository.ProductRepository
}

func ProviderProductService(p repository.ProductRepository) ProductsServices {
	return ProductsServices{Repositori: p}
}

func (p *ProductsServices) FindAll() []models.Products {
	return p.Repositori.FindAll()
}

func (p *ProductsServices) FindByID(id int) models.Products {
	return p.Repositori.FindByID(id)
}

func (p *ProductsServices) Save(product models.Products) (models.Products, error) {
	return p.Repositori.Save(product)
}
