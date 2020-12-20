package ustbe

import "singleaf/ustbe/models"

// UstbeRepo is a interface for function repository
type UstbeRepo interface {
	Login(ustbe *models.Ustbe) (*models.Ustbe, error)
	CheckMail(ustbe *models.Ustbe) bool
	Create(ustbe *models.Ustbe) (*models.Ustbe, error)
	FindAll() ([]*models.UstbeWrapper, error)
	FindByID(id int) (*models.UstbeWrapper, error)
	//FindAllUserSubs(id string) ([]*models.AllUserUstbe, error)
	FindAllUS(id int) ([]*models.AllUserUstbe, error)
	Update(ustbe *models.Ustbe) (*models.Ustbe, error)
	Delete(id int) error
	UpdatePhoto(ustbe *models.Ustbe) error
}
