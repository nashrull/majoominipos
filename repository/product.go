package repository

import (
	"majoominipos/models"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProductRepositoryProvider(dbProv *gorm.DB) ProductRepository {
	return ProductRepository{DB: dbProv}
}

func (p *ProductRepository) FindAll() []models.Products {
	var products []models.Products
	p.DB.Find(&products)
	return products
}

func (p *ProductRepository) FindByID(id int) models.Products {
	var products models.Products
	p.DB.First(&products, id)
	return products
}

func (p *ProductRepository) Save(product models.Products) (models.Products, error) {
	productSave := p.DB.Save(&product)
	if productSave.Error != nil {
		return models.Products{}, productSave.Error
	}
	return product, nil
}

func (p *ProductRepository) Delete(product ProductRepository) {
	p.DB.Delete(&product)
}
