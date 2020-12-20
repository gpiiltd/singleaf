package subscriber

import (
	//"singleaf/besta"
	"singleaf/besta/common"
	"singleaf/besta/models"

	//"github.com/jinzhu/gorm"

	"github.com/jinzhu/gorm"
	//"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// BestaRepositoryImpl is use sharing connection
type Subscriber struct {
	database *gorm.DB
}


// FindAll use when you want to show all besta data
func (call *Subscriber) FindAll() ([]*models.BestaWrapper, error) {
	subsList := make([]*models.BestaWrapper, 0)

	err := call.database.Table("besta").Find(&subsList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return subsList, nil
}

func FindAllUS(id int) (int, error) {
	return 100, nil
}
// FindAll use when you want to show all besta data
// func FindAllUS(id int) ([]*models.AllUserBesta, error) {
// 	var call *Subscriber{} //= new(Subscriber)
// 	subsList := make([]*models.AllUserBesta, 0)

// 	//err := db.Table("besta").Where("userid = ?", id).Find(&subsList).Error
// 	err := call.database.Table("besta").Find(&subsList).Error
// 	//db.Not("service = ?", "jinzhu").First(&subsList)
// 	if err != nil {
// 		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
// 		return nil, err
// 	}
// 	return subsList, nil
// }

