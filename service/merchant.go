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

func (p *MerchantService) Login(merchant models.Merchants) (bool, error) {
	merchantDB, err := p.MerchantRepository.FindByUsername(merchant.Username)
	if err != nil {
		return false, err
	}
	cekPassword := helpers.CheckPasswordHash(merchant.Password, merchantDB.Password)
	if !cekPassword {
		return false, errors.New("password salah")
	}
	return true, nil
}
