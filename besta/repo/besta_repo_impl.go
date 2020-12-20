package repo

import (
	"singleaf/besta"
	"singleaf/besta/common"
	"singleaf/besta/models"

	"github.com/jinzhu/gorm"
)

// BestaRepositoryImpl is use sharing connection
type BestaRepositoryImpl struct {
	database *gorm.DB
}

// CreateRepo use for getting connection
func CreateRepo(db *gorm.DB) besta.BestaRepo {
	return &BestaRepositoryImpl{db}
}

// Login for check a valid email when besta login in
func (call *BestaRepositoryImpl) Login(besta *models.Besta) (*models.Besta, error) {
	allsubs := new(models.Besta)

	err := call.database.Table("besta").Where("email = ?", besta.Useremail).Take(&allsubs).Error
	if err != nil {
		// common.LogError("Login", "Error when trying geeting email address, error is =>", err)
		return nil, err
	}
	return allsubs, nil
}

// CheckMail use for validate email is already use or nah
func (call *BestaRepositoryImpl) CheckMail(besta *models.Besta) bool {
	allsubs := new(models.BestaWrapper)
	err := call.database.Raw("SELECT * FROM \"besta\" WHERE  useremail = ? LIMIT 1", besta.Useremail).Scan(allsubs).Error
	if err != nil {
		return true // email already registered
	}
	return false // email not registered
}

// Create use for create a new account besta
func (call *BestaRepositoryImpl) Create(besta *models.Besta) (*models.Besta, error) {
	err := call.database.Table("besta").Create(&besta).Error
	if err != nil {
		common.LogError("Create", "Error wehn trying to create use error is =>", err)
		return nil, err
	}
	return besta, nil
}

// FindByID use for search besta by id
func (call *BestaRepositoryImpl) FindByID(id int) (*models.BestaWrapper, error) {
	besta := new(models.BestaWrapper)

	err := call.database.Table("besta").Where("id = ?", id).First(&besta).Error
	if err != nil {
		common.LogError("FindByID", "Error when trying to find data by id, error is =>", err)
		return nil, err
	}

	return besta, nil
}

// FindAll use when you want to show all besta data
func (call *BestaRepositoryImpl) FindAll() ([]*models.BestaWrapper, error) {
	subsList := make([]*models.BestaWrapper, 0)

	err := call.database.Table("besta").Find(&subsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return subsList, nil
}


// FindAll use when you want to show all besta data
func (call *BestaRepositoryImpl) FindAllUS(id int) ([]*models.AllUserBesta, error) {
	subsList := make([]*models.AllUserBesta, 0)

	err := call.database.Table("besta").Where("userid = ?", id).Find(&subsList).Error
	//err := call.database.Table("besta").Find(&subsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return subsList, nil
}

// FindAll use when you want to show all besta data
// func (call *BestaRepositoryImpl) FindAllUserSubs(id string) ([]*models.AllUserBesta, error) {
// 	subsList := make([]*models.AllUserBesta, 0)

// 	err := call.database.Table("besta").Where("userid = ?", id).Find(&subsList).Error
// 	//err := call.database.Table("besta").Where("id = ?", id).First(&besta).Error
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

// Update use for update besta data
func (call *BestaRepositoryImpl) Update(besta *models.Besta) (*models.Besta, error) {

	allsubs := new(models.Besta)

	err := call.database.Table("besta").Where("id = ?", besta.ID).First(&allsubs).Update(&besta).Error
	if err != nil {
		common.LogError("Update", "Error when trying to update besta data, error is =>", err)
		return nil, err
	}
	return allsubs, nil
}

// Delete use for delete account
func (call *BestaRepositoryImpl) Delete(id int) error {
	err := call.database.Where("id = ?", id).Delete(&models.Besta{}).Error
	if err != nil {
		common.LogError("Delete", "Error when trying to delete besta, error is =>", err)
		return err
	}
	return nil
}

// UpdatePhoto use for update a picture
func (call *BestaRepositoryImpl) UpdatePhoto(besta *models.Besta) error {
	allsubs := new(models.Photo)
	err := call.database.Table("besta").Where(" id = ?", besta.ID).First(&allsubs).Update(besta).Error
	if err != nil {
		common.LogError("UpdatePhoto", "Error when trying to update photo profile, error is =>", err)
		return err
	}
	return nil
}
