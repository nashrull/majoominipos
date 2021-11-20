package repository

import (
	"majoominipos/models"

	"github.com/jinzhu/gorm"
)

type MerchantRepository struct {
	DB *gorm.DB
}

func MerchantRepositoryProvider(dbProv *gorm.DB) MerchantRepository {
	return MerchantRepository{DB: dbProv}
}

func (p *MerchantRepository) FindByUsername(username string) (models.Merchants, error) {
	var merchant models.Merchants
	m := p.DB.Where("username = ?", username).Find(&merchant)
	if m.Error != nil {
		return models.Merchants{}, m.Error
	}
	return merchant, nil
}

func (p *MerchantRepository) Save(merchant models.Merchants) (models.Merchants, error) {
	merchantSave := p.DB.Save(&merchant)
	if merchantSave.Error != nil {
		return merchant, merchantSave.Error
	}
	return merchant, nil
}
