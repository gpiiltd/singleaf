package usecase

import (
	"singleaf/besta"
	"singleaf/besta/common"
	"singleaf/besta/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// BestaUsecaseImpl use for get a repo connection
type BestaUsecaseImpl struct {
	bestaRepo besta.BestaRepo
}

// CreateUsecase use for get connection from repository
func CreateUsecase(bestaRepo besta.BestaRepo) besta.BestaUsecase {
	return &BestaUsecaseImpl{bestaRepo}
}

// CheckMail use for business logic when a new account want to regsiter he mail
// this function will check the mail is already register or nah
func (call *BestaUsecaseImpl) CheckMail(besta *models.Besta) bool {
	return call.bestaRepo.CheckMail(besta)
}

// Login use for business logic when use trying to loggin in
func (call *BestaUsecaseImpl) Login(besta *models.Besta) (*models.Besta, error) {
	model, err := call.bestaRepo.Login(besta)
	if err != nil {
		return nil, errors.New("Email not registered")
	}

	err = common.VerifyPassword(model.Useremail, besta.Useremail)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("Error invalid password")
	}

	return model, nil
}

// Create use for business logic when a new account is create
func (call *BestaUsecaseImpl) Create(besta *models.Besta) (*models.Besta, error) {
	besta, err := common.Encrypt(besta)
	if err != nil {
		return nil, err
	}

	status := call.bestaRepo.CheckMail(besta)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	return call.bestaRepo.Create(besta)
}

// Update use for business logic when update a account
func (call *BestaUsecaseImpl) Update(besta *models.Besta) (*models.Besta, error) {

	status := call.bestaRepo.CheckMail(besta)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	besta, err := common.Encrypt(besta)
	if err != nil {
		return nil, err
	}

	return call.bestaRepo.Update(besta)
}

// FindAll use for business logic when you want to print all besta account to the List
func (call *BestaUsecaseImpl) FindAll() ([]*models.BestaWrapper, error) {
	return call.bestaRepo.FindAll()
}

// FindAll use for business logic when you want to print all besta account to the List
func (call *BestaUsecaseImpl) FindAllUS(id int) ([]*models.AllUserBesta, error) {
	return call.bestaRepo.FindAllUS(id)
}


// FindAllUserSubs use for business logic when you want to print all besta account to the List
// func (call *BestaUsecaseImpl) FindAllUserSubs(id string) ([]*models.AllUserBesta, error) {
// 	return call.bestaRepo.FindAllUserSubs(id)
// }

// FindByID use for business logic when you want to find account by id
func (call *BestaUsecaseImpl) FindByID(id int) (*models.BestaWrapper, error) {
	return call.bestaRepo.FindByID(id)
}

// Delete use for delete use
func (call *BestaUsecaseImpl) Delete(id int) error {
	return call.bestaRepo.Delete(id)
}

// UpdatePhoto use for update photo profile
func (call *BestaUsecaseImpl) UpdatePhoto(besta *models.Besta) error {
	return call.bestaRepo.UpdatePhoto(besta)
}
