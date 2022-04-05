package service

import (
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type DormOwnerService interface {
	GetDormOwner(id string) model.DormOwner
}

func DormOwnerServiceHandler(dormOwnerRepository repository.DormOwnerRepository) DormOwnerService {
	return &dormOwnerService{
		dormOwnerRepository: dormOwnerRepository,
	}
}

type dormOwnerService struct {
	dormOwnerRepository repository.DormOwnerRepository
}

func (dormOwnerService *dormOwnerService) GetDormOwner(id string) model.DormOwner {
	dormOwner, err := dormOwnerService.dormOwnerRepository.FindDormOwnerByID(id)

	if err != nil {
		panic(err)
	}

	return dormOwner
}
