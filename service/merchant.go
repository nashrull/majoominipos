package service

import (
	"errors"
	"majoominipos/helpers"
	"majoominipos/models"
	"majoominipos/repository"
)

type MerchantService struct {
	MerchantRepository repository.MerchantRepository
}

func MerchantServiceProvider(p repository.MerchantRepository) MerchantService {
	return MerchantService{MerchantRepository: p}
}

func (p *MerchantService) Registration(merchant models.Merchants) (models.Merchants, error) {
	newMerchant := merchant
	newMerchant.Password, _ = helpers.HashPassword(merchant.Password)
	_, err := p.MerchantRepository.Save(newMerchant)
	if err != nil {
		return models.Merchants{}, err
	}
	return newMerchant, nil
}

func (p *MerchantService) Login(merchant models.Merchants) (models.Merchants, error) {
	merchantDB, err := p.MerchantRepository.FindByUsername(merchant.Username)
	if err != nil {
		return models.Merchants{}, errors.New("data tidak ditemukan salah")
	}
	cekPassword := helpers.CheckPasswordHash(merchant.Password, merchantDB.Password)
	if !cekPassword {
		return models.Merchants{}, errors.New("password salah")
	}
	return merchantDB, nil
}

func (p *MerchantService) Delete(id int) error {
	return p.MerchantRepository.Delete(id)
}

func (p *MerchantService) Update(merchant models.Merchants) (models.Merchants, error) {
	// find merchant
	var updateMerchant models.Merchants
	updateMerchant, err := p.MerchantRepository.FindByUsername(merchant.Username)
	if err != nil {
		return updateMerchant, err
	}
	updateMerchant.Password, _ = helpers.HashPassword(merchant.Password)
	updateMerchant.Username = merchant.Username
	updateMerchant.Status = merchant.Status
	updateMerchant.Nama = merchant.Nama
	return p.MerchantRepository.Save(updateMerchant)
}

func (p *MerchantService) GetAll() []models.Merchants {
	var ListOfMerchant []models.Merchants
	ListOfMerchant = p.MerchantRepository.GetAll()
	return ListOfMerchant
}
