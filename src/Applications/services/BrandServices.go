package services

import (
	"voucher-redeem-api/src/Domains/brands/entities"
)

type repository interface {
	InsertBrand(createBrand entities.CreateBrand) (*entities.Brand, error)
}

type Service struct {
	repo repository
}

func NewBrandService(repo repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateNewBrand(name string) (*entities.Brand, error) {
	createBrand, err := entities.NewCreateBrand(name)

	if err != nil {
		return nil, err
	}

	createdBrand, err := s.repo.InsertBrand(*createBrand)

	if err != nil {
		return nil, err
	}

	return createdBrand, nil
}
