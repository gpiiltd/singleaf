package usecase

import (
	"singleaf/apps"
	"singleaf/apps/common"
	"singleaf/apps/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// AppsUsecaseImpl use for get a repo connection
type AppsUsecaseImpl struct {
	appsRepo apps.AppsRepo
}

// CreateUsecase use for get connection from repository
func CreateUsecase(appsRepo apps.AppsRepo) apps.AppsUsecase {
	return &AppsUsecaseImpl{appsRepo}
}

// CheckMail use for business logic when a new account want to regsiter he mail
// this function will check the mail is already register or nah
func (call *AppsUsecaseImpl) CheckMail(apps *models.Apps) bool {
	return call.appsRepo.CheckMail(apps)
}

// Login use for business logic when use trying to loggin in
func (call *AppsUsecaseImpl) Login(apps *models.Apps) (*models.Apps, error) {
	model, err := call.appsRepo.Login(apps)
	if err != nil {
		return nil, errors.New("Email not registered")
	}

	err = common.VerifyPassword(model.Role, apps.Role)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("Error invalid data")
	}

	return model, nil
}

// Create use for business logic when a new account is create
func (call *AppsUsecaseImpl) Create(apps *models.Apps) (*models.Apps, error) {
	apps, err := common.Encrypt(apps)
	if err != nil {
		return nil, err
	}

	status := call.appsRepo.CheckMail(apps)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	return call.appsRepo.Create(apps)
}

// Update use for business logic when update a account
func (call *AppsUsecaseImpl) Update(apps *models.Apps) (*models.Apps, error) {

	status := call.appsRepo.CheckMail(apps)
	if !status {
		return nil, errors.New("Opps.. sorry email already use other account")
	}

	apps, err := common.Encrypt(apps)
	if err != nil {
		return nil, err
	}

	return call.appsRepo.Update(apps)
}

// FindAll use for business logic when you want to print all apps account to the List
func (call *AppsUsecaseImpl) FindAll() ([]*models.AppsWrapper, error) {
	return call.appsRepo.FindAll()
}

// FindByID use for business logic when you want to find account by id
func (call *AppsUsecaseImpl) FindByID(id int) (*models.AppsWrapper, error) {
	return call.appsRepo.FindByID(id)
}

// Delete use for delete use
func (call *AppsUsecaseImpl) Delete(id int) error {
	return call.appsRepo.Delete(id)
}

func (call *AppsUsecaseImpl) UpdatePhoto(apps *models.Apps) error {
	return call.appsRepo.UpdatePhoto(apps)
}
