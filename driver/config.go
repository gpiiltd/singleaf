package driver

import (
	"singleaf/user/models"
	appmodels "singleaf/apps/models"
	ustbemodels "singleaf/ustbe/models"
	bestamodels "singleaf/besta/models"
	enterprisemodels "singleaf/enterprises/models"
	//armsmodels "singleaf/arms/models"
	"fmt"
	"log"
	"net/url"

    "os/exec"
    "bytes"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func CreatePgDb(dbname string) {
    cmd := exec.Command("createdb", "-p", "5432", "-h", "127.0.0.1", "-U", viper.GetString("database.user"), "-P", viper.GetString("database.pass"), "-e", dbname)
    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Run(); err != nil {
        log.Printf("Error: %v", err)
    }
    log.Printf("Output: %q\n", out.String())
}

func Config() *gorm.DB {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.db_name")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")

	//createPgDb(dbName)

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	getConnection, err := gorm.Open("postgres", connStr)

	err = getConnection.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// common.InitTable(getConnection)

	getConnection.SingularTable(true)

	getConnection.Debug().AutoMigrate(
		&models.User{},
		//&models.Platformdate{},
	)

	getConnection.Debug().AutoMigrate(
		&appmodels.Apps{},
	)

	getConnection.Debug().AutoMigrate(
		&ustbemodels.Ustbe{}, //User_Subscribe_To_Business_Entity{},
		//&submodels.Besta{}, //Business_Entity_Subscribe_To_App{},

	)

	getConnection.Debug().AutoMigrate(
		&bestamodels.Besta{}, //User_Subscribe_To_Business_Entity{},
		//&bestamodels.Besta{}, //Business_Entity_Subscribe_To_App{},

	)

	getConnection.Debug().AutoMigrate(
		&enterprisemodels.Enterprises{},
		//&enterprisemodels.Subscription{},
		//&enterprisemodels.Sessions{},
	)

	// getConnection.Debug().AutoMigrate(
	// 	&enterprisemodels.Semester{},
	// )

	// getConnection.Debug().AutoMigrate(
	// 	&enterprisemodels.Sessions{},
	// )

	// getConnection.Debug().AutoMigrate(
	// 	&armsmodels.Arms{},
	// 	&armsmodels.CI{},
	// )

	return getConnection

}

func ConfigSchema(db string) *gorm.DB {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := db //viper.GetString("database.db_name")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")

	//createPgDb(dbName)

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	getConnection, err := gorm.Open("postgres", connStr)

	err = getConnection.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return getConnection

}