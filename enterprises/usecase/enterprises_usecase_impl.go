package usecase

import (
	"singleaf/enterprises"
	"singleaf/enterprises/common"
	"singleaf/enterprises/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// EnterprisesUsecaseImpl use for get a repo connection
type EnterprisesUsecaseImpl struct {
	enterprisesRepo enterprises.EnterprisesRepo
}

// CreateUsecase use for get connection from repository
func CreateUsecase(enterprisesRepo enterprises.EnterprisesRepo) enterprises.EnterprisesUsecase {
	return &EnterprisesUsecaseImpl{enterprisesRepo}
}

// CheckMail use for business logic when a new account want to regsiter he mail
// this function will check the mail is already register or nah
func (call *EnterprisesUsecaseImpl) CheckMail(enterprises *models.Enterprises) bool {
	return call.enterprisesRepo.CheckMail(enterprises)
}

// Login use for business logic when use trying to loggin in
func (call *EnterprisesUsecaseImpl) Login(enterprises *models.Enterprises) (*models.Enterprises, error) {
	model, err := call.enterprisesRepo.Login(enterprises)
	if err != nil {
		return nil, errors.New("Email not registered")
	}

	err = common.VerifyPassword(model.EnterpriseHash, enterprises.EnterpriseHash)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("Error invalid password")
	}

	return model, nil
}

// Create use for business logic when a new account is create
func (call *EnterprisesUsecaseImpl) Create(enterprises *models.Enterprises) (*models.Enterprises, error) {
	enterprises, err := common.Encrypt(enterprises)
	if err != nil {
		return nil, err
	}

	status := call.enterprisesRepo.CheckMail(enterprises)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	return call.enterprisesRepo.Create(enterprises)
}

// Update use for business logic when update a account
func (call *EnterprisesUsecaseImpl) Update(enterprises *models.Enterprises) (*models.Enterprises, error) {

	status := call.enterprisesRepo.CheckMail(enterprises)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	enterprises, err := common.Encrypt(enterprises)
	if err != nil {
		return nil, err
	}

	return call.enterprisesRepo.Update(enterprises)
}

// FindAll use for business logic when you want to print all enterprises account to the List
func (call *EnterprisesUsecaseImpl) FindAll() ([]*models.EnterprisesWrapper, error) {
	return call.enterprisesRepo.FindAll()
}

// FindByID use for business logic when you want to find account by id
func (call *EnterprisesUsecaseImpl) FindByID(id int) (*models.EnterprisesWrapper, error) {
	return call.enterprisesRepo.FindByID(id)
}

func (call *EnterprisesUsecaseImpl) CheckAdmin(email string, biztag string) (*models.EnterprisesWrapper, error) {
	return call.enterprisesRepo.CheckAdmin(email, biztag)
}

func (call *EnterprisesUsecaseImpl) MyDomains(email string) ([]*models.EnterprisesWrapper, error) {
	return call.enterprisesRepo.MyDomains(email)
}
// Delete use for delete use
func (call *EnterprisesUsecaseImpl) Delete(id int) error {
	return call.enterprisesRepo.Delete(id)
}

// UpdatePhoto use for update photo profile
func (call *EnterprisesUsecaseImpl) UpdatePhoto(enterprises *models.Enterprises) error {
	return call.enterprisesRepo.UpdatePhoto(enterprises)
}
