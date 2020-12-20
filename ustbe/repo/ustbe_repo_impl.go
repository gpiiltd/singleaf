package repo

import (
	"singleaf/ustbe"
	"singleaf/ustbe/common"
	"singleaf/ustbe/models"

	"github.com/jinzhu/gorm"
)

// UstbeRepositoryImpl is use sharing connection
type UstbeRepositoryImpl struct {
	database *gorm.DB
}

// CreateRepo use for getting connection
func CreateRepo(db *gorm.DB) ustbe.UstbeRepo {
	return &UstbeRepositoryImpl{db}
}

// Login for check a valid email when ustbe login in
func (call *UstbeRepositoryImpl) Login(ustbe *models.Ustbe) (*models.Ustbe, error) {
	allsubs := new(models.Ustbe)

	err := call.database.Table("ustbe").Where("email = ?", ustbe.Useremail).Take(&allsubs).Error
	if err != nil {
		// common.LogError("Login", "Error when trying geeting email address, error is =>", err)
		return nil, err
	}
	return allsubs, nil
}

// CheckMail use for validate email is already use or nah
func (call *UstbeRepositoryImpl) CheckMail(ustbe *models.Ustbe) bool {
	allsubs := new(models.UstbeWrapper)
	err := call.database.Raw("SELECT * FROM \"ustbe\" WHERE  useremail = ? LIMIT 1", ustbe.Useremail).Scan(allsubs).Error
	if err != nil {
		return true // email already registered
	}
	return false // email not registered
}

// Create use for create a new account ustbe
func (call *UstbeRepositoryImpl) Create(ustbe *models.Ustbe) (*models.Ustbe, error) {
	err := call.database.Table("ustbe").Create(&ustbe).Error
	if err != nil {
		common.LogError("Create", "Error wehn trying to create use error is =>", err)
		return nil, err
	}
	return ustbe, nil
}

// FindByID use for search ustbe by id
func (call *UstbeRepositoryImpl) FindByID(id int) (*models.UstbeWrapper, error) {
	ustbe := new(models.UstbeWrapper)

	err := call.database.Table("ustbe").Where("id = ?", id).First(&ustbe).Error
	if err != nil {
		common.LogError("FindByID", "Error when trying to find data by id, error is =>", err)
		return nil, err
	}

	return ustbe, nil
}

// FindAll use when you want to show all ustbe data
func (call *UstbeRepositoryImpl) FindAll() ([]*models.UstbeWrapper, error) {
	subsList := make([]*models.UstbeWrapper, 0)

	err := call.database.Table("ustbe").Find(&subsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return subsList, nil
}


// FindAll use when you want to show all ustbe data
func (call *UstbeRepositoryImpl) FindAllUS(id int) ([]*models.AllUserUstbe, error) {
	subsList := make([]*models.AllUserUstbe, 0)

	err := call.database.Table("ustbe").Where("userid = ?", id).Find(&subsList).Error
	//err := call.database.Table("ustbe").Find(&subsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return subsList, nil
}

// FindAll use when you want to show all ustbe data
// func (call *UstbeRepositoryImpl) FindAllUserSubs(id string) ([]*models.AllUserUstbe, error) {
// 	subsList := make([]*models.AllUserUstbe, 0)

// 	err := call.database.Table("ustbe").Where("userid = ?", id).Find(&subsList).Error
// 	//err := call.database.Table("ustbe").Where("id = ?", id).First(&ustbe).Error
// 	//db.Where("name <> ?", "jinzhu").Find(&users)
// 	if err != nil {
// 		common.LogError("FindAllUserSubs", "Error when trying to get all ada, error is =>", err)
// 		return nil, err
// 	}
// 	return subsList, nil


// 	// result := db.First(&user)
// 	// result.RowsAffected // returns found records count
// 	// result.Error        // returns error

// 	// // check error ErrRecordNotFound
// 	// errors.Is(result.Error, gorm.ErrRecordNotFound)
// }

// Update use for update ustbe data
func (call *UstbeRepositoryImpl) Update(ustbe *models.Ustbe) (*models.Ustbe, error) {

	allsubs := new(models.Ustbe)

	err := call.database.Table("ustbe").Where("id = ?", ustbe.ID).First(&allsubs).Update(&ustbe).Error
	if err != nil {
		common.LogError("Update", "Error when trying to update ustbe data, error is =>", err)
		return nil, err
	}
	return allsubs, nil
}

// Delete use for delete account
func (call *UstbeRepositoryImpl) Delete(id int) error {
	err := call.database.Where("id = ?", id).Delete(&models.Ustbe{}).Error
	if err != nil {
		common.LogError("Delete", "Error when trying to delete ustbe, error is =>", err)
		return err
	}
	return nil
}

// UpdatePhoto use for update a picture
func (call *UstbeRepositoryImpl) UpdatePhoto(ustbe *models.Ustbe) error {
	allsubs := new(models.Photo)
	err := call.database.Table("ustbe").Where(" id = ?", ustbe.ID).First(&allsubs).Update(ustbe).Error
	if err != nil {
		common.LogError("UpdatePhoto", "Error when trying to update photo profile, error is =>", err)
		return err
	}
	return nil
}
