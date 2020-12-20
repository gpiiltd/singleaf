package subscriber

import (
	//"singleaf/ustbe"
	"singleaf/ustbe/common"
	"singleaf/ustbe/models"

	//"github.com/jinzhu/gorm"

	"github.com/jinzhu/gorm"
	//"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// UstbeRepositoryImpl is use sharing connection
type Subscriber struct {
	database *gorm.DB
}


// FindAll use when you want to show all ustbe data
func (call *Subscriber) FindAll() ([]*models.UstbeWrapper, error) {
	subsList := make([]*models.UstbeWrapper, 0)

	err := call.database.Table("ustbe").Find(&subsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return subsList, nil
}

func FindAllUS(id int) (int, error) {
	return 100, nil
}
// FindAll use when you want to show all ustbe data
// func FindAllUS(id int) ([]*models.AllUserUstbe, error) {
// 	var call *Subscriber{} //= new(Subscriber)
// 	subsList := make([]*models.AllUserUstbe, 0)

// 	//err := db.Table("ustbe").Where("userid = ?", id).Find(&subsList).Error
// 	err := call.database.Table("ustbe").Find(&subsList).Error
// 	//db.Not("service = ?", "jinzhu").First(&subsList)
// 	if err != nil {
// 		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
// 		return nil, err
// 	}
// 	return subsList, nil
// }

