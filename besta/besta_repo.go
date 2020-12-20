package besta

import "singleaf/besta/models"

// BestaRepo is a interface for function repository
type BestaRepo interface {
	Login(besta *models.Besta) (*models.Besta, error)
	CheckMail(besta *models.Besta) bool
	Create(besta *models.Besta) (*models.Besta, error)
	FindAll() ([]*models.BestaWrapper, error)
	FindByID(id int) (*models.BestaWrapper, error)
	//FindAllUserSubs(id string) ([]*models.AllUserBesta, error)
	FindAllUS(id int) ([]*models.AllUserBesta, error)
	Update(besta *models.Besta) (*models.Besta, error)
	Delete(id int) error
	UpdatePhoto(besta *models.Besta) error
}
