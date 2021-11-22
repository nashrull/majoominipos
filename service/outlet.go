package service

import (
	"majoominipos/models"
	"majoominipos/repository"
)

type OutletService struct {
	OutletRepository repository.OutletRepository
}

func OutletServiceProvider(p repository.OutletRepository) OutletService {
	return OutletService{OutletRepository: p}
}

func (o OutletService) CreateOutlet(outlet models.Outlets) (models.Outlets, error) {
	return o.OutletRepository.Create(outlet)
}
