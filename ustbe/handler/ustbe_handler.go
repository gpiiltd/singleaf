package handler

import (
	//"singleaf/auth"
	"singleaf/middlewares"
	"singleaf/ustbe"
	"singleaf/ustbe/common"
	"singleaf/ustbe/models"
	"encoding/json"
	"fmt"
	"log"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// UstbeHandler struct use for get funcntion in business logic
type UstbeHandler struct {
	ustbeUsecase ustbe.UstbeUsecase
}

// CreateHandler use for handling request
func CreateHandler(r *mux.Router, usecase ustbe.UstbeUsecase) {
	UstbeHandler := UstbeHandler{usecase}

	//r.HandleFunc("/ustbe/findallusersubs{id}", UstbeHandler.FindAllUserSubs).Methods(http.MethodGet)

	//r.HandleFunc("/login", UstbeHandler.loginUser).Methods(http.MethodPost)
	//r.HandleFunc("/register", UstbeHandler.createUser).Methods(http.MethodPost)findallusersubs

	// make a new subrouter when you want to grouping where path want to be protect and FindAllUserSubs
	authorized := r.NewRoute().Subrouter()
	authorized.Use(middlewares.SetMiddlewareAuthentication)
	authorized.HandleFunc("/ustbe", UstbeHandler.findAll).Methods(http.MethodGet)
	authorized.HandleFunc("/ustbe/user/{id}", UstbeHandler.UserFindSub).Methods(http.MethodGet)
	authorized.HandleFunc("/ustbe/publisher/{id}", UstbeHandler.PublisherFindSub).Methods(http.MethodGet)
	authorized.HandleFunc("/ustbe/register", UstbeHandler.createUstbe).Methods(http.MethodPost)
	authorized.HandleFunc("/ustbe/{id}", UstbeHandler.findByID).Methods(http.MethodGet)
	//authorized.HandleFunc("/ustbe/findallusersubs{id}", UstbeHandler.FindAllUserSubs).Methods(http.MethodGet)
	authorized.HandleFunc("/ustbe/{id}", UstbeHandler.updateSubs).Methods(http.MethodPut)
	authorized.HandleFunc("/ustbe/{id}", UstbeHandler.deleteSubs).Methods(http.MethodDelete)

	// make a new subrouter extends a authorized path
	uploadRequest := authorized.PathPrefix("/ustbe").Subrouter()
	uploadRequest.Use(middlewares.FileSizeLimiter) // use middleware to limited size when upload file
	uploadRequest.HandleFunc("/photo/{id}", UstbeHandler.handlingPhoto).Methods(http.MethodPost)
}

// func (call *UstbeHandler) loginSubs(w http.ResponseWriter, r *http.Request) {

// 	ustbe := new(models.Ustbe)
// 	//ustbe.Name = "King Starr"

// 	fmt.Println("%+v", ustbe)

// 	//fmt.Println("%+v", r)

// 	if err := json.NewDecoder(r.Body).Decode(&ustbe); err != nil {
// 		// w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser", "Error when trying to login, error is =>", err)
// 		// common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
// 		// return

// 		r.ParseForm()


// 		ustbe.Useremail = r.FormValue("email") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
// 		//ustbe.Password = r.FormValue("password") //string `gorm:"column:password;size:100;not null;" json:"password"`
// 	}

	

// 	err := common.Validate("login", ustbe)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser 1", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	ustbe, err = call.ustbeUsecase.Login(ustbe)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		// common.LogError("loginUser 2", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	// token, err := auth.CreateToken(ustbe)
// 	// if err != nil {
// 	// 	common.LogError("loginUser 3", "Error when trying to generate token, error is =>", err)
// 	// 	common.Response(w, common.Message(false, "Opps.. something when wrong", nil))
// 	// 	return
// 	// }

// 	// allsubs := new(models.UstbeWrapper)
// 	// copier.Copy(&allsubs, &ustbe)

// 	// finalResponse := map[string]interface{}{"Apps": "status", "User": allsubs, "data": map[string]interface{}{"token": token}}

// 	// common.Response(w, common.Message(true, "Success", finalResponse))

// 	common.Response(w, common.Message(true, "Success", map[string]interface{}{"token": "kingstarrtoken"}))
// }

func (call *UstbeHandler) createUstbe(w http.ResponseWriter, r *http.Request) {

	ustbe := new(models.Ustbe)

	
	//if err := json.NewDecoder(r.Body).Decode(&ustbe); err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		//common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		//return
		log.Println(r.Form)
		r.ParseForm()

		

		//ustbe.Fullname = r.FormValue("fullname") //"Kingstarr Alistring" // `gorm:"column:name;size:255;not null" json:"name"`
		//ustbe.Useremail = r.FormValue("useremail") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
		
		ustbe.UserID = r.FormValue("userid") //string `gorm:"column:notlpn;not null" json:"no_tlpn"`
		ustbe.CompanyTag = r.FormValue("companytag") //string `gorm:"column:password;size:100;not null;" json:"password"`
		ustbe.CompanyID =  r.FormValue("companyid") //string `gorm:"column:gender;size:15;not null" json:"gender"`
		ustbe.Description =  r.FormValue("companyname")
		ustbe.RecordID = ustbe.CompanyTag+ustbe.UserID
	//}
		
	err := common.Validate("register", ustbe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		log.Println(err.Error())
		log.Println("king")
		return
	}

	// val, _ := auth.TokenValid(r)
	// if val.Role != "admin" && ustbe.Role == "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	response, err := call.ustbeUsecase.Create(ustbe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), response))
		return
	}

	common.Response(w, common.Message(true, "Subscribed successfully", response))
	return
}

func (call *UstbeHandler) findAll(w http.ResponseWriter, r *http.Request) {

	// val, err := auth.TokenValid(r)
	// if err != nil || val.Role != "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	result, err := call.ustbeUsecase.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

func (call *UstbeHandler) UserFindSub(w http.ResponseWriter, r *http.Request) {

	// val, err := auth.TokenValid(r)
	// if err != nil || val.Role != "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }
	//id := "1"

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	result, err := call.ustbeUsecase.FindAllUS(IDSubs)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}

	if len(result) == 0 {
		//w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

func (call *UstbeHandler) PublisherFindSub(w http.ResponseWriter, r *http.Request) {

	// val, err := auth.TokenValid(r)
	// if err != nil || val.Role != "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }
	//id := "1"

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	result, err := call.ustbeUsecase.FindAllUS(IDSubs)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}

	if len(result) == 0 {
		//w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

// func (call *UstbeHandler) FindAllUserSubs(w http.ResponseWriter, r *http.Request) {

// 	// val, err := auth.TokenValid(r)
// 	// if err != nil || val.Role != "admin" {
// 	// 	w.WriteHeader(http.StatusForbidden)
// 	// 	common.Response(w, common.Message(false, "Access denied", nil))
// 	// 	return
// 	// }

	
// 	result := []*models.AllUserUstbe{}

// 	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
// 		return
// 	}

// 	log.Println(IDSubs)

// 	result, err = call.ustbeUsecase.FindAllUserSubs("g")
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		common.Response(w, common.Message(false, "result empty", nil))
// 		return
// 	}
// 	common.Response(w, common.Message(true, "Success", result))
// }


func (call *UstbeHandler) findByID(w http.ResponseWriter, r *http.Request) {
	ustbe := new(models.UstbeWrapper)

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	//allsubs, err := call.ustbeUsecase.FindByID(IDSubs)
	// if val.Role != "Platform Admin" {
	// 	if err != nil || val.ID != allsubs.ID {
	// 		w.WriteHeader(http.StatusForbidden)
	// 		common.Response(w, common.Message(false, "Access denied", nil))
	// 		return
	// 	}
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	common.Response(w, common.Message(false, err.Error(), nil))
	// 	return
	// }

	ustbe, err = call.ustbeUsecase.FindByID(IDSubs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", ustbe))
}

func (call *UstbeHandler) updateSubs(w http.ResponseWriter, r *http.Request) {
	ustbe := new(models.Ustbe)

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&ustbe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	allsubs, err := call.ustbeUsecase.FindByID(IDSubs)
	// if val.Role != "Platform Admin" {
	// 	if err != nil || val.ID != allsubs.ID {
	// 		w.WriteHeader(http.StatusForbidden)
	// 		common.Response(w, common.Message(false, "Access denied", nil))
	// 		return
	// 	}
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	common.Response(w, common.Message(false, err.Error(), nil))
	// 	return
	// }

	ustbe.ID = allsubs.ID
	err = common.Validate("update", ustbe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	ustbe, err = call.ustbeUsecase.Update(ustbe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Update successfully", ustbe))
	return
}

func (call *UstbeHandler) deleteSubs(w http.ResponseWriter, r *http.Request) {
	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	//allsubs, err := call.ustbeUsecase.FindByID(IDSubs)
	// if val.Role != "Platform Admin" {
	// 	if err != nil || val.ID != allsubs.ID {
	// 		w.WriteHeader(http.StatusForbidden)
	// 		common.Response(w, common.Message(false, "Access denied", nil))
	// 		return
	// 	}
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	common.Response(w, common.Message(false, err.Error(), nil))
	// 	return
	// }

	err = call.ustbeUsecase.Delete(IDSubs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. something when wrong", nil))
		return
	}

	common.Response(w, common.Message(true, "Success delete account", nil))
}

func (call *UstbeHandler) handlingPhoto(w http.ResponseWriter, r *http.Request) {

	ustbe := new(models.Ustbe)

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	allsubs, err := call.ustbeUsecase.FindByID(IDSubs)
	// if val.Role != "Platform Admin" {
	// 	if err != nil || val.ID != allsubs.ID {
	// 		w.WriteHeader(http.StatusForbidden)
	// 		common.Response(w, common.Message(false, "Access denied", nil))
	// 		return
	// 	}
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	common.Response(w, common.Message(false, err.Error(), nil))
	// 	return
	// }

	filePath, err := call.handleFile(r, allsubs, "photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. error when trying handling file, please  contact you support. error is : "+err.Error(), nil))
		return
	}

	allsubs.Photo = filePath
	copier.Copy(&ustbe, &allsubs)
	err = call.ustbeUsecase.UpdatePhoto(ustbe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success file was uploaded", filePath))
}

func (call *UstbeHandler) handleFile(r *http.Request, ustbe *models.UstbeWrapper, key string) (string, error) {
	file, handler, err := r.FormFile(key)
	if err != nil {
		common.LogError("handleFile", "Error When Handle File, error is => ", err)
		return "", err
	}

	defer file.Close()

	image := viper.GetString("file.path") + "/" + ustbe.Photo
	if _, err := os.Stat(image); !os.IsNotExist(err) && !strings.Contains(ustbe.Photo, "default.jpeg") {
		os.Remove(image)
	}

	fileNameSlice := strings.Split(handler.Filename, ".")
	fileName := fmt.Sprintf("%v_%v_%v.%v", key, fileNameSlice[0], time.Now().Format("21504052006"), fileNameSlice[len(fileNameSlice)-1])

	filePath := filepath.Join(viper.GetString("file.path"), fileName)

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
