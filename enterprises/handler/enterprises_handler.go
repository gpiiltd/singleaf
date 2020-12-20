package handler

import (
	"singleaf/auth"
	"singleaf/middlewares"
	"singleaf/enterprises"
	"singleaf/enterprises/common"
	"singleaf/enterprises/models"
	//"singleaf/driver"
	//"singleaf/subscriptions/subscriber"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"log"

	"github.com/jinzhu/copier"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// EnterprisesHandler struct use for get funcntion in business logic
type EnterprisesHandler struct {
	enterprisesUsecase enterprises.EnterprisesUsecase
}

// CreateHandler use for handling request
func CreateHandler(r *mux.Router, usecase enterprises.EnterprisesUsecase) {
	enterprisesHandler := EnterprisesHandler{usecase}

	
	// make a new subrouter when you want to grouping where path want to be protect and nah
	authorized := r.NewRoute().Subrouter()
	authorized.Use(middlewares.SetMiddlewareAuthentication)
	authorized.HandleFunc("/enterprise", enterprisesHandler.findAll).Methods(http.MethodGet)
	authorized.HandleFunc("/enterprise/register", enterprisesHandler.createEnterprise).Methods(http.MethodPost)
	authorized.HandleFunc("/enterprise/{id}", enterprisesHandler.findByID).Methods(http.MethodGet)
	authorized.HandleFunc("/enterprise/checkadmin/{email}", enterprisesHandler.checkAdmin).Methods(http.MethodGet)
	authorized.HandleFunc("/enterprise/mydomains/{email}", enterprisesHandler.myDomains).Methods(http.MethodGet)
	authorized.HandleFunc("/enterprise/{id}", enterprisesHandler.updateEnterprise).Methods(http.MethodPut)
	authorized.HandleFunc("/enterprise/{id}", enterprisesHandler.deleteEnterprise).Methods(http.MethodDelete)

	// make a new subrouter extends a authorized path
	uploadRequest := authorized.PathPrefix("/enterprise").Subrouter()
	uploadRequest.Use(middlewares.FileSizeLimiter) // use middleware to limited size when upload file
	uploadRequest.HandleFunc("/photo/{id}", enterprisesHandler.handlingPhoto).Methods(http.MethodPost)
}

// func (call *EnterprisesHandler) loginEnterprise(w http.ResponseWriter, r *http.Request) {

// 	enterprises := new(models.Enterprises)
// 	//enterprises.EnterpriseName = "King Starr"

// 	fmt.Println("%+v", enterprises)

// 	//fmt.Println("%+v", r)

// 	if err := json.NewDecoder(r.Body).Decode(&enterprises); err != nil {
// 		// w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser", "Error when trying to login, error is =>", err)
// 		// common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
// 		// return

// 		r.ParseForm()


// 		enterprises.EnterpriseEmail = r.FormValue("email") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
// 		enterprises.EnterpriseHash = r.FormValue("password") //string `gorm:"column:password;size:100;not null;" json:"password"`
// 	}

	

// 	err := common.Validate("login", enterprises)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser 1", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	enterprises, err = call.enterprisesUsecase.Login(enterprises)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		// common.LogError("loginUser 2", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	// token, err := auth.CreateToken(enterprises)
// 	// if err != nil {
// 	// 	common.LogError("loginUser 3", "Error when trying to generate token, error is =>", err)
// 	// 	common.Response(w, common.Message(false, "Opps.. something when wrong", nil))
// 	// 	return
// 	// }

// 	enterprisess := new(models.EnterprisesWrapper)
// 	copier.Copy(&enterprisess, &enterprises)

// 	//get subscription information

// 	subscriptionInfo := "" //subscriber.FindByID(enterprises) //call new subscriptionrepo link //map[string]interface{}{"Name": "Proprietor"}

// 	finalResponse := map[string]interface{}{"Subscriptions": subscriptionInfo, "Enterprises": enterprisess, "data": map[string]interface{}{"token": "token"}}

// 	common.Response(w, common.Message(true, "Success", finalResponse))
// }

func (call *EnterprisesHandler) createEnterprise(w http.ResponseWriter, r *http.Request) {

	enterprises := new(models.Enterprises)

	if err := json.NewDecoder(r.Body).Decode(&enterprises); err != nil {

		r.ParseForm()




		enterprises.EnterpriseName = r.FormValue("enterprisename") //"Kingstarr Alistring" // `gorm:"column:name;size:255;not null" json:"name"`
		enterprises.EnterpriseEmail = r.FormValue("enterpriseemail") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
		enterprises.CompanyTag = r.FormValue("companytag") //string `gorm:"column:notlpn;not null" json:"no_tlpn"`
		enterprises.EnterpriseHash = r.FormValue("enterpriseemail") //string `gorm:"column:password;size:100;not null;" json:"password"`
		enterprises.Description =  "Mixed Enterprise" //string `gorm:"column:gender;size:15;not null" json:"gender"`
		enterprises.Address =   r.FormValue("address") //string `gorm:"column:address;size:300;not null" json:"address"`
		enterprises.Role =    "Primary Enterprise"
		enterprises.EnterpriseLogo ="default.jpeg"
		enterprises.VerifyID = r.FormValue("verifyid")
		//enterprises.CompanyTag = "ksi"
	}

	err := common.Validate("register", enterprises)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	// db := driver.Config()

	// defer func() {
	// 	err := db.Close()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()

	// //driver.CreatePgDb(enterprises.CompanyTag)

	// autocreatedb := db.Exec("create database " + enterprises.CompanyTag)

	//log.Println("Output: %q\n", autocreatedb)

	response, err := call.enterprisesUsecase.Create(enterprises)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(true, err.Error(), response))
		return
	}

	common.Response(w, common.Message(true, "Register successfully", response))
	return
}

func (call *EnterprisesHandler) findAll(w http.ResponseWriter, r *http.Request) {

	// val, err := auth.TokenValid(r)
	// if err != nil || val.Role != "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	result, err := call.enterprisesUsecase.FindAll()
	if err != nil {
		//w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}

	fmt.Println("*********************************************")
	fmt.Println(result)
	fmt.Println("*********************************************")

	
	fmt.Println(len(result))
	if len(result) == 0 {
    	//w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(true, "Empty Query Response", result))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

func (call *EnterprisesHandler) myDomains(w http.ResponseWriter, r *http.Request) {
	//enterprises := new(models.EnterprisesWrapper)
	enterprises := make([]*models.EnterprisesWrapper, 0)

	//IDUser, err := strconv.Atoi(mux.Vars(r)["email"])
	auvars := mux.Vars(r)
  	email := auvars["email"]
  	
  	log.Println(email+" ___king___")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request: unknown email address! ", nil))
		return
	}

	enterprises, err := call.enterprisesUsecase.MyDomains(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", enterprises))
}

func (call *EnterprisesHandler) checkAdmin(w http.ResponseWriter, r *http.Request) {
	enterprises := new(models.EnterprisesWrapper)

	//IDUser, err := strconv.Atoi(mux.Vars(r)["email"])
	auvars := mux.Vars(r)
  	email := auvars["email"]


  	//if err := json.NewDecoder(r.Body).Decode(&enterprises); err != nil {

		r.ParseForm()

		bizTag := r.FormValue("companytag") //string `gorm:"column:notlpn;not null" json:"no_tlpn"`
		
	//}
  	
  	log.Println(email+" ___king___")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request: unknown email address! ", nil))
		return
	}

	enterprises, err := call.enterprisesUsecase.CheckAdmin(email, bizTag)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", enterprises))
}


func (call *EnterprisesHandler) findByID(w http.ResponseWriter, r *http.Request) {
	enterprises := new(models.EnterprisesWrapper)

	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	//users, err := call.enterprisesUsecase.FindByID(IDUser)
	// if val.Role != "admin" {
	// 	if err != nil || val.ID != users.ID {
	// 		w.WriteHeader(http.StatusForbidden)
	// 		common.Response(w, common.Message(false, "Access denied", nil))
	// 		return
	// 	}
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	common.Response(w, common.Message(false, err.Error(), nil))
	// 	return
	// }

	enterprises, err = call.enterprisesUsecase.FindByID(IDUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", enterprises))
}

func (call *EnterprisesHandler) updateEnterprise(w http.ResponseWriter, r *http.Request) {
	enterprises := new(models.Enterprises)

	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&enterprises)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.enterprisesUsecase.FindByID(IDUser)
	if val.Role != "Proprietor" {
		if err != nil || val.ID != users.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	enterprises.ID = users.ID
	err = common.Validate("update", enterprises)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	enterprises, err = call.enterprisesUsecase.Update(enterprises)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Update successfully", enterprises))
	return
}

func (call *EnterprisesHandler) deleteEnterprise(w http.ResponseWriter, r *http.Request) {
	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.enterprisesUsecase.FindByID(IDUser)
	if val.Role != "Platform Admin" {
		if err != nil || val.ID != users.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	err = call.enterprisesUsecase.Delete(IDUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. something when wrong", nil))
		return
	}

	common.Response(w, common.Message(true, "Success delete account", nil))
}

func (call *EnterprisesHandler) handlingPhoto(w http.ResponseWriter, r *http.Request) {

	enterprises := new(models.Enterprises)

	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.enterprisesUsecase.FindByID(IDUser)
	if val.Role != "Proprietor" {
		if err != nil || val.ID != users.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	filePath, err := call.handleFile(r, users, "photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. error when trying handling file, please  contact you support. error is : "+err.Error(), nil))
		return
	}

	users.EnterpriseLogo = filePath
	copier.Copy(&enterprises, &users)
	err = call.enterprisesUsecase.UpdatePhoto(enterprises)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success file was uploaded", filePath))
}

func (call *EnterprisesHandler) handleFile(r *http.Request, enterprises *models.EnterprisesWrapper, key string) (string, error) {
	file, handler, err := r.FormFile(key)
	if err != nil {
		common.LogError("handleFile", "Error When Handle File, error is => ", err)
		return "", err
	}

	defer file.Close()

	image := viper.GetString("publisher.path") + "/" + enterprises.EnterpriseLogo
	if _, err := os.Stat(image); !os.IsNotExist(err) && !strings.Contains(enterprises.EnterpriseLogo, "default.jpeg") {
		os.Remove(image)
	}

	fileNameSlice := strings.Split(handler.Filename, ".")
	fileName := fmt.Sprintf("%v_%v_%v.%v", key, fileNameSlice[0], time.Now().Format("21504052006"), fileNameSlice[len(fileNameSlice)-1])

	filePath := filepath.Join(viper.GetString("publisher.path"), fileName)

	targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		common.LogError("handleFile", "Error when open file with, error is => ", err)
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		common.LogError("handleFile", "Error when copy file, error is => ", err)
		return "", err
	}
	return fileName, nil
}
