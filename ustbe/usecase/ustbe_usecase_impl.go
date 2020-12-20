package usecase

import (
	"singleaf/ustbe"
	"singleaf/ustbe/common"
	"singleaf/ustbe/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UstbeUsecaseImpl use for get a repo connection
type UstbeUsecaseImpl struct {
	ustbeRepo ustbe.UstbeRepo
}

// CreateUsecase use for get connection from repository
func CreateUsecase(ustbeRepo ustbe.UstbeRepo) ustbe.UstbeUsecase {
	return &UstbeUsecaseImpl{ustbeRepo}
}

// CheckMail use for business logic when a new account want to regsiter he mail
// this function will check the mail is already register or nah
func (call *UstbeUsecaseImpl) CheckMail(ustbe *models.Ustbe) bool {
	return call.ustbeRepo.CheckMail(ustbe)
}

// Login use for business logic when use trying to loggin in
func (call *UstbeUsecaseImpl) Login(ustbe *models.Ustbe) (*models.Ustbe, error) {
	model, err := call.ustbeRepo.Login(ustbe)
	if err != nil {
		return nil, errors.New("Email not registered")
	}

	err = common.VerifyPassword(model.Useremail, ustbe.Useremail)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("Error invalid password")
	}

	return model, nil
}

// Create use for business logic when a new account is create
func (call *UstbeUsecaseImpl) Create(ustbe *models.Ustbe) (*models.Ustbe, error) {
	ustbe, err := common.Encrypt(ustbe)
	if err != nil {
		return nil, err
	}

	status := call.ustbeRepo.CheckMail(ustbe)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	return call.ustbeRepo.Create(ustbe)
}

// Update use for business logic when update a account
func (call *UstbeUsecaseImpl) Update(ustbe *models.Ustbe) (*models.Ustbe, error) {

	status := call.ustbeRepo.CheckMail(ustbe)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	ustbe, err := common.Encrypt(ustbe)
	if err != nil {
		return nil, err
	}

	return call.ustbeRepo.Update(ustbe)
}

// FindAll use for business logic when you want to print all ustbe account to the List
func (call *UstbeUsecaseImpl) FindAll() ([]*models.UstbeWrapper, error) {
	return call.ustbeRepo.FindAll()
}

// FindAll use for business logic when you want to print all ustbe account to the List
func (call *UstbeUsecaseImpl) FindAllUS(id int) ([]*models.AllUserUstbe, error) {
	return call.ustbeRepo.FindAllUS(id)
}


// FindAllUserSubs use for business logic when you want to print all ustbe account to the List
// func (call *UstbeUsecaseImpl) FindAllUserSubs(id string) ([]*models.AllUserUstbe, error) {
// 	return call.ustbeRepo.FindAllUserSubs(id)
// }

// FindByID use for business logic when you want to find account by id
func (call *UstbeUsecaseImpl) FindByID(id int) (*models.UstbeWrapper, error) {
	return call.ustbeRepo.FindByID(id)
}

// Delete use for delete use
func (call *UstbeUsecaseImpl) Delete(id int) error {
	return call.ustbeRepo.Delete(id)
}

// UpdatePhoto use for update photo profile
func (call *UstbeUsecaseImpl) UpdatePhoto(ustbe *models.Ustbe) error {
	return call.ustbeRepo.UpdatePhoto(ustbe)
}
