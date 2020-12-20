package handler

import (
	"singleaf/auth"
	"singleaf/middlewares"
	"singleaf/apps"
	"singleaf/apps/common"
	"singleaf/apps/models"
	"encoding/json"
	"fmt"
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

// AppsHandler struct use for get funcntion in business logic
type AppsHandler struct {
	appsUsecase apps.AppsUsecase
}

// CreateHandler use for handling request
func CreateHandler(r *mux.Router, usecase apps.AppsUsecase) {
	appsHandler := AppsHandler{usecase}

	//r.HandleFunc("/login", appsHandler.loginUser).Methods(http.MethodPost)
	//r.HandleFunc("/register", appsHandler.createUser).Methods(http.MethodPost)

	// make a new subrouter when you want to grouping where path want to be protect and nah
	authorized := r.NewRoute().Subrouter()
	authorized.Use(middlewares.SetMiddlewareAuthentication)
	authorized.HandleFunc("/apps/register", appsHandler.createApps).Methods(http.MethodPost)
	authorized.HandleFunc("/apps", appsHandler.findAll).Methods(http.MethodGet)
	authorized.HandleFunc("/apps/{id}", appsHandler.findByID).Methods(http.MethodGet)
	authorized.HandleFunc("/apps/{id}", appsHandler.updateApps).Methods(http.MethodPut)
	authorized.HandleFunc("/apps/{id}", appsHandler.deleteApps).Methods(http.MethodDelete)

	// make a new subrouter extends a authorized path
	uploadRequest := authorized.PathPrefix("/apps").Subrouter()
	uploadRequest.Use(middlewares.FileSizeLimiter) // use middleware to limited size when upload file
	uploadRequest.HandleFunc("/photo/{id}", appsHandler.handlingPhoto).Methods(http.MethodPost)
}

// func (call *AppsHandler) loginApps(w http.ResponseWriter, r *http.Request) {

// 	apps := new(models.Apps)
// 	//user.Name = "King Starr"

// 	fmt.Println("%+v", apps)

// 	//fmt.Println("%+v", r)

// 	if err := json.NewDecoder(r.Body).Decode(&apps); err != nil {
// 		// w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser", "Error when trying to login, error is =>", err)
// 		// common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
// 		// return

// 		r.ParseForm()


// 		apps.Email = r.FormValue("email") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
// 		apps.Name = r.FormValue("password") //string `gorm:"column:password;size:100;not null;" json:"password"`
// 	}

	

// 	err := common.Validate("login", apps)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser 1", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	apps, err = call.appsUsecase.Login(apps)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		// common.LogError("loginUser 2", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	//token, err := auth.CreateToken(apps)
// 	// if err != nil {
// 	// 	common.LogError("loginApps 3", "Error when trying to generate token, error is =>", err)
// 	// 	common.Response(w, common.Message(false, "Opps.. something when wrong", nil))
// 	// 	return
// 	// }
// 	common.Response(w, common.Message(true, "Success", map[string]interface{}{"token": "kingstarrtoken"}))
// }

func (call *AppsHandler) createApps(w http.ResponseWriter, r *http.Request) {

	apps := new(models.Apps)

	//user.Name = "King Starr"
	//r.ParseForm()

	//fmt.Println(r.Form)




	if err := json.NewDecoder(r.Body).Decode(&apps); err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		//common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		//return

		r.ParseForm()




		apps.Name = r.FormValue("name")
		apps.Email = r.FormValue("email") 
		apps.Description = r.FormValue("description") 
		apps.Status = "1" 
		apps.Role =    "app"
		apps.Photo ="default.jpeg"
	}

	err := common.Validate("register", apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	// val, _ := auth.TokenValid(r)
	// if val.Role != "Platform Admin" && user.Role == "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	response, err := call.appsUsecase.Create(apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), response))
		return
	}

	common.Response(w, common.Message(true, "Register successfully", response))
	return
}

func (call *AppsHandler) findAll(w http.ResponseWriter, r *http.Request) {

	val, err := auth.TokenValid(r)
	if err != nil || val.Role != "Platform Admin" {
		w.WriteHeader(http.StatusForbidden)
		common.Response(w, common.Message(false, "Access denied", nil))
		return
	}

	result, err := call.appsUsecase.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

func (call *AppsHandler) findByID(w http.ResponseWriter, r *http.Request) {
	apps := new(models.AppsWrapper)

	IDApps, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	apps_s, err := call.appsUsecase.FindByID(IDApps)
	if val.Role != "Platform Admin" {
		if err != nil || val.ID != apps_s.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	apps, err = call.appsUsecase.FindByID(IDApps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", apps))
}

func (call *AppsHandler) updateApps(w http.ResponseWriter, r *http.Request) {
	apps := new(models.Apps)

	IDApps, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	apps_s, err := call.appsUsecase.FindByID(IDApps)
	if val.Role != "Platform Admin" {
		if err != nil || val.ID != apps_s.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	apps.ID = apps_s.ID
	err = common.Validate("update", apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	apps, err = call.appsUsecase.Update(apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Update successfully", apps))
	return
}

func (call *AppsHandler) deleteApps(w http.ResponseWriter, r *http.Request) {
	IDApps, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	apps_s, err := call.appsUsecase.FindByID(IDApps)
	if val.Role != "Platform Admin" {
		if err != nil || val.ID != apps_s.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	err = call.appsUsecase.Delete(IDApps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. something when wrong", nil))
		return
	}

	common.Response(w, common.Message(true, "Success delete account", nil))
}

func (call *AppsHandler) handlingPhoto(w http.ResponseWriter, r *http.Request) {

	apps := new(models.Apps)

	IDApps, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	apps_s, err := call.appsUsecase.FindByID(IDApps)
	if val.Role != "Platform Admin" {
		if err != nil || val.ID != apps_s.ID {
			w.WriteHeader(http.StatusForbidden)
			common.Response(w, common.Message(false, "Access denied", nil))
			return
		}
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	filePath, err := call.handleFile(r, apps_s, "photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. error when trying handling file, please  contact you support. error is : "+err.Error(), nil))
		return
	}

	apps_s.Photo = filePath
	copier.Copy(&apps, &apps_s)
	err = call.appsUsecase.UpdatePhoto(apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success file was uploaded", filePath))
}

func (call *AppsHandler) handleFile(r *http.Request, apps *models.AppsWrapper, key string) (string, error) {
	file, handler, err := r.FormFile(key)
	if err != nil {
		common.LogError("handleFile", "Error When Handle File, error is => ", err)
		return "", err
	}

	defer file.Close()

	image := viper.GetString("file.path") + "/" + apps.Photo
	if _, err := os.Stat(image); !os.IsNotExist(err) && !strings.Contains(apps.Photo, "default.jpeg") {
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
