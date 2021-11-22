package repository

import (
	"majoominipos/models"

	"github.com/jinzhu/gorm"
)

type OutletRepository struct {
	DB *gorm.DB
}

func OutletRepositoryProvider(dbProv *gorm.DB) OutletRepository {
	return OutletRepository{DB: dbProv}
}

func (o *OutletRepository) FindAll(Id_merchant string) ([]models.Outlets, error) {
	var outlet []models.Outlets
	data := o.DB.Where("id_merchant = ?", Id_merchant).Find(&outlet)
	return outlet, data.Error
}

func (o *OutletRepository) Create(Outlet models.Outlets) (models.Outlets, error) {
	simpan := o.DB.Save(&Outlet)
	if simpan.Error != nil {
		return models.Outlets{}, nil
	}
	return Outlet, nil
}
