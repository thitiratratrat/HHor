package service

import (
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type DormOwnerService interface {
	GetDorms(id string) []model.Dorm
}

func DormOwnerServiceHandler(dormRepository repository.DormRepository) DormOwnerService {
	return &dormOwnerService{
		dormRepository: dormRepository,
	}
}

type dormOwnerService struct {
	dormRepository repository.DormRepository
}

func (dormOwnerService *dormOwnerService) GetDorms(id string) []model.Dorm {
	return dormOwnerService.dormRepository.FindDormOwnerDorms(id)
}
