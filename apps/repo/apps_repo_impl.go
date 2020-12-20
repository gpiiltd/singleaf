package repo

import (
	"singleaf/apps"
	"singleaf/apps/common"
	"singleaf/apps/models"

	"github.com/jinzhu/gorm"
)

// AppsRepositoryImpl is use sharing connection
type AppsRepositoryImpl struct {
	database *gorm.DB
}

// CreateRepo use for getting connection
func CreateRepo(db *gorm.DB) apps.AppsRepo {
	return &AppsRepositoryImpl{db}
}

// Login for check a valid email when apps login in
func (call *AppsRepositoryImpl) Login(apps *models.Apps) (*models.Apps, error) {
	allapps := new(models.Apps)

	err := call.database.Table("apps").Where("email = ?", apps.Email).Take(&allapps).Error
	if err != nil {
		// common.LogError("Login", "Error when trying geeting email address, error is =>", err)
		return nil, err
	}
	return allapps, nil
}

// CheckMail use for validate email is already use or nah
func (call *AppsRepositoryImpl) CheckMail(apps *models.Apps) bool {
	apps_s := new(models.AppsWrapper)
	err := call.database.Raw("SELECT * FROM \"apps\" WHERE  email = ? LIMIT 1", apps.Email).Scan(apps_s).Error
	if err != nil {
		return true // email already registered
	}
	return false // email not registered
}

// Create use for create a new account apps
func (call *AppsRepositoryImpl) Create(apps *models.Apps) (*models.Apps, error) {
	err := call.database.Table("apps").Create(&apps).Error
	if err != nil {
		common.LogError("Create", "Error wehn trying to create use error is =>", err)
		return nil, err
	}
	return apps, nil
}

// FindByID use for search apps by id
func (call *AppsRepositoryImpl) FindByID(id int) (*models.AppsWrapper, error) {
	apps := new(models.AppsWrapper)

	err := call.database.Table("apps").Where("id = ?", id).First(&apps).Error
	if err != nil {
		common.LogError("FindByID", "Error when trying to find data by id, error is =>", err)
		return nil, err
	}

	return apps, nil
}

// FindAll use when you want to show all apps data
func (call *AppsRepositoryImpl) FindAll() ([]*models.AppsWrapper, error) {
	appsList := make([]*models.AppsWrapper, 0)

	err := call.database.Table("apps").Find(&appsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return appsList, nil
}

// Update use for update apps data
func (call *AppsRepositoryImpl) Update(apps *models.Apps) (*models.Apps, error) {

	allapps := new(models.Apps)

	err := call.database.Table("apps").Where("id = ?", apps.ID).First(&allapps).Update(&apps).Error
	if err != nil {
		common.LogError("Update", "Error when trying to update apps data, error is =>", err)
		return nil, err
	}
	return allapps, nil
}

// Delete use for delete account
func (call *AppsRepositoryImpl) Delete(id int) error {
	err := call.database.Where("id = ?", id).Delete(&models.Apps{}).Error
	if err != nil {
		common.LogError("Delete", "Error when trying to delete apps, error is =>", err)
		return err
	}
	return nil
}

// UpdatePhoto use for update a picture
func (call *AppsRepositoryImpl) UpdatePhoto(apps *models.Apps) error {
	allapps := new(models.Photo)
	err := call.database.Table("apps").Where(" id = ?", apps.ID).First(&allapps).Update(apps).Error
	if err != nil {
		common.LogError("UpdatePhoto", "Error when trying to update photo profile, error is =>", err)
		return err
	}
	return nil
}
