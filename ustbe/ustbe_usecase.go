package ustbe

import "singleaf/ustbe/models"

// UstbeUsecase is a interface for function business logic
type UstbeUsecase interface {
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
