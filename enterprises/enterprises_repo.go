package enterprises

import "singleaf/enterprises/models"

// EnterprisesRepo is a interface for function repository
type EnterprisesRepo interface {
	Login(enterprises *models.Enterprises) (*models.Enterprises, error)
	CheckMail(enterprises *models.Enterprises) bool
	Create(enterprises *models.Enterprises) (*models.Enterprises, error)
	FindAll() ([]*models.EnterprisesWrapper, error)
	FindByID(id int) (*models.EnterprisesWrapper, error)
	CheckAdmin(email string, biztag string) (*models.EnterprisesWrapper, error)
	MyDomains(email string) ([]*models.EnterprisesWrapper, error)
	Update(enterprises *models.Enterprises) (*models.Enterprises, error)
	Delete(id int) error
	UpdatePhoto(enterprises *models.Enterprises) error
}
