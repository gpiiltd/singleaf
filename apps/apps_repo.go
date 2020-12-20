package apps

import "singleaf/apps/models"

// UserRepo is a interface for function repository
type AppsRepo interface {
	Login(apps *models.Apps) (*models.Apps, error)
	CheckMail(apps *models.Apps) bool
	Create(apps *models.Apps) (*models.Apps, error)
	FindAll() ([]*models.AppsWrapper, error)
	FindByID(id int) (*models.AppsWrapper, error)
	Update(apps *models.Apps) (*models.Apps, error)
	Delete(id int) error
	UpdatePhoto(apps *models.Apps) error
}
