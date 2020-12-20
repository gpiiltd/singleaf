package repo

import (
	"singleaf/enterprises"
	"singleaf/enterprises/common"
	"singleaf/enterprises/models"
	"singleaf/driver"
	"log"

	"github.com/jinzhu/gorm"
)

// EnterprisesRepositoryImpl is use sharing connection
type EnterprisesRepositoryImpl struct {
	database *gorm.DB
}

// CreateRepo use for getting connection
func CreateRepo(db *gorm.DB) enterprises.EnterprisesRepo {
	return &EnterprisesRepositoryImpl{db}
}

// Login for check a valid email when enterprises login in
func (call *EnterprisesRepositoryImpl) Login(enterprises *models.Enterprises) (*models.Enterprises, error) {
	users := new(models.Enterprises)

	err := call.database.Table("enterprises").Where("email = ?", enterprises.EnterpriseEmail).Take(&users).Error
	if err != nil {
		// common.LogError("Login", "Error when trying geeting email address, error is =>", err)
		return nil, err
	}
	return users, nil
}

// CheckMail use for validate email is already use or nah
func (call *EnterprisesRepositoryImpl) CheckMail(enterprises *models.Enterprises) bool {
	users := new(models.EnterprisesWrapper)
	err := call.database.Raw("SELECT * FROM \"enterprises\" WHERE  enterpriseemail = ? LIMIT 1", enterprises.EnterpriseEmail).Scan(users).Error
	if err != nil {
		return true // email already registered
	}
	return false // email not registered
}

// Create use for create a new account enterprises
func (call *EnterprisesRepositoryImpl) Create(enterprises *models.Enterprises) (*models.Enterprises, error) {

	//driver.CreatePgDb(enterprises.CompanyTag)

	autocreatedb := call.database.Exec("create database " + enterprises.CompanyTag)
	log.Println("Output: %q\n", autocreatedb)

	db := driver.ConfigSchema(enterprises.CompanyTag)

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	autocreateschema := db.Exec("create schema " + enterprises.CompanyTag)
	log.Println("Output: %q\n", autocreateschema)


	err := call.database.Table("enterprises").Create(&enterprises).Error
	if err != nil {
		common.LogError("Create", "Error wehn trying to create use error is =>", err)
		return nil, err
	}
	return enterprises, nil
}

// FindByID use for search enterprises by id
func (call *EnterprisesRepositoryImpl) MyDomains(email string) ([]*models.EnterprisesWrapper, error) {
	//enterprises := new(models.EnterprisesWrapper)
	//db.where("mail = ? AND tag >= ?", "starr", "gpitest").Find(&enter)
	enterprises := make([]*models.EnterprisesWrapper, 0)

	err := call.database.Table("enterprises").Where("enterpriseemail = ?", email).Find(&enterprises).Error
	if err != nil {
		common.LogError("CheckMail", "Error when trying to find data by email, error is =>", err)
		return nil, err
	}

	return enterprises, nil
}

func (call *EnterprisesRepositoryImpl) CheckAdmin(email string, biztag string) (*models.EnterprisesWrapper, error) {
	enterprises := new(models.EnterprisesWrapper)
	//db.where("mail = ? AND tag >= ?", "starr", "gpitest").Find(&enter)

	err := call.database.Table("enterprises").Where("enterpriseemail = ? AND companytag = ?", email, biztag).First(&enterprises).Error
	if err != nil {
		common.LogError("CheckMail", "Error when trying to find data by email, error is =>", err)
		return nil, err
	}

	return enterprises, nil
}
func (call *EnterprisesRepositoryImpl) FindByID(id int) (*models.EnterprisesWrapper, error) {
	enterprises := new(models.EnterprisesWrapper)

	err := call.database.Table("enterprises").Where("id = ?", id).First(&enterprises).Error
	if err != nil {
		common.LogError("FindByID", "Error when trying to find data by id, error is =>", err)
		return nil, err
	}

	return enterprises, nil
}

// FindAll use when you want to show all enterprises data
func (call *EnterprisesRepositoryImpl) FindAll() ([]*models.EnterprisesWrapper, error) {
	userList := make([]*models.EnterprisesWrapper, 0)

	err := call.database.Table("enterprises").Find(&userList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return userList, nil
}

// Update use for update enterprises data
func (call *EnterprisesRepositoryImpl) Update(enterprises *models.Enterprises) (*models.Enterprises, error) {

	users := new(models.Enterprises)

	err := call.database.Table("enterprises").Where("id = ?", enterprises.ID).First(&users).Update(&enterprises).Error
	if err != nil {
		common.LogError("Update", "Error when trying to update enterprises data, error is =>", err)
		return nil, err
	}
	return users, nil
}

// Delete use for delete account
func (call *EnterprisesRepositoryImpl) Delete(id int) error {
	err := call.database.Where("id = ?", id).Delete(&models.Enterprises{}).Error
	if err != nil {
		common.LogError("Delete", "Error when trying to delete enterprises, error is =>", err)
		return err
	}
	return nil
}

// UpdatePhoto use for update a picture
func (call *EnterprisesRepositoryImpl) UpdatePhoto(enterprises *models.Enterprises) error {
	users := new(models.EnterpriseLogo)
	err := call.database.Table("enterprises").Where(" id = ?", enterprises.ID).First(&users).Update(enterprises).Error
	if err != nil {
		common.LogError("UpdatePhoto", "Error when trying to update photo profile, error is =>", err)
		return err
	}
	return nil
}
