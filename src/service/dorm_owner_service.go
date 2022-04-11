package service

import (
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"github.com/thitiratratrat/hhor/src/utils"
)

type DormOwnerService interface {
	GetDormOwner(id string) model.DormOwner
	UpdateDormOwner(id string, updateDormOwnerDTO dto.UpdateDormOwnerDTO) model.DormOwner
	UpdateDormOwnerPictures(id string, profilePictureUrl string, bankQrUrl string) model.DormOwner
	UpdateBankAccount(id string, updateBankAccountDTO dto.UpdateBankAccountDTO) model.DormOwner
	DeleteBankAccount(id string) model.DormOwner
}

func DormOwnerServiceHandler(dormOwnerRepository repository.DormOwnerRepository, encryptor utils.Encryptor) DormOwnerService {
	return &dormOwnerService{
		dormOwnerRepository: dormOwnerRepository,
		encryptor:           encryptor,
	}
}

type dormOwnerService struct {
	dormOwnerRepository repository.DormOwnerRepository
	encryptor           utils.Encryptor
}

func (dormOwnerService *dormOwnerService) GetDormOwner(id string) model.DormOwner {
	dormOwner, err := dormOwnerService.dormOwnerRepository.FindDormOwnerByID(id)

	if err != nil {
		panic(err)
	}

	if dormOwner.BankAccount.AccountName == nil && dormOwner.BankAccount.BankQrUrl == nil {
		dormOwner.BankAccount = nil
	} else {
		if dormOwner.BankAccount.AccountName != nil {
			decryptedAccountName, _ := dormOwnerService.encryptor.Decrypt(*dormOwner.BankAccount.AccountName)
			dormOwner.BankAccount.AccountName = &decryptedAccountName
		}

		if dormOwner.BankAccount.AccountNumber != nil {
			decryptedAccontNumber, _ := dormOwnerService.encryptor.Decrypt(*dormOwner.BankAccount.AccountNumber)
			dormOwner.BankAccount.AccountNumber = &decryptedAccontNumber
		}
	}

	return dormOwner
}

func (dormOwnerService *dormOwnerService) UpdateDormOwner(id string, updateDormOwnerDTO dto.UpdateDormOwnerDTO) model.DormOwner {
	dormOwner := mapUpdateDormOwner(updateDormOwnerDTO)
	_, err := dormOwnerService.dormOwnerRepository.UpdateDormOwner(id, dormOwner)

	if err != nil {
		panic(err)
	}

	return dormOwnerService.GetDormOwner(id)
}

func (dormOwnerService *dormOwnerService) UpdateDormOwnerPictures(id string, profilePictureUrl string, bankQrUrl string) model.DormOwner {
	dormOwner := mapUpdateDormOwnerPicture(profilePictureUrl, bankQrUrl)
	_, err := dormOwnerService.dormOwnerRepository.UpdateDormOwner(id, dormOwner)

	if err != nil {
		panic(err)
	}

	return dormOwnerService.GetDormOwner(id)
}

func (dormOwnerService *dormOwnerService) UpdateBankAccount(id string, updateBankAccountDTO dto.UpdateBankAccountDTO) model.DormOwner {
	dormOwner := dormOwnerService.mapUpdateBankAccount(updateBankAccountDTO)
	_, err := dormOwnerService.dormOwnerRepository.UpdateDormOwner(id, dormOwner)

	if err != nil {
		panic(err)
	}

	return dormOwnerService.GetDormOwner(id)
}

func (dormOwnerService *dormOwnerService) DeleteBankAccount(id string) model.DormOwner {
	_, err := dormOwnerService.dormOwnerRepository.DeleteBankAccount(id)

	if err != nil {
		panic(err)
	}

	return dormOwnerService.GetDormOwner(id)
}

func mapUpdateDormOwner(updateDormOwnerDTO dto.UpdateDormOwnerDTO) model.DormOwner {
	return model.DormOwner{
		Firstname:   updateDormOwnerDTO.Firstname,
		Lastname:    updateDormOwnerDTO.Lastname,
		LineID:      updateDormOwnerDTO.LineID,
		PhoneNumber: updateDormOwnerDTO.PhoneNumber,
	}
}

func mapUpdateDormOwnerPicture(profilePictureUrl string, bankQrUrl string) model.DormOwner {
	var dormOwner model.DormOwner

	if len(profilePictureUrl) > 0 {
		dormOwner.PictureUrl = &profilePictureUrl
	}

	if len(bankQrUrl) > 0 {
		dormOwner.BankAccount = &model.BankAccount{
			BankQrUrl: &bankQrUrl,
		}

	}

	return dormOwner
}

func (dormOwnerService *dormOwnerService) mapUpdateBankAccount(updateBankAccountDTO dto.UpdateBankAccountDTO) model.DormOwner {
	encryptedAccountName, err := dormOwnerService.encryptor.Encrypt(updateBankAccountDTO.AccountName)

	if err != nil {
		panic(err)
	}

	encryptedAccountNumber, err := dormOwnerService.encryptor.Encrypt(updateBankAccountDTO.AccountNumber)

	if err != nil {
		panic(err)
	}

	return model.DormOwner{
		BankAccount: &model.BankAccount{
			Bank:          &updateBankAccountDTO.Bank,
			AccountName:   &encryptedAccountName,
			AccountNumber: &encryptedAccountNumber,
		},
	}
}
