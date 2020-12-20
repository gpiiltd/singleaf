package handler

import (
	//"singleaf/auth"
	"singleaf/middlewares"
	"singleaf/besta"
	"singleaf/besta/common"
	"singleaf/besta/models"
	"encoding/json"
	"fmt"
	//"log"
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

// BestaHandler struct use for get funcntion in business logic
type BestaHandler struct {
	bestaUsecase besta.BestaUsecase
}

// CreateHandler use for handling request
func CreateHandler(r *mux.Router, usecase besta.BestaUsecase) {
	BestaHandler := BestaHandler{usecase}

	//r.HandleFunc("/besta/findallusersubs{id}", BestaHandler.FindAllUserSubs).Methods(http.MethodGet)

	//r.HandleFunc("/login", BestaHandler.loginUser).Methods(http.MethodPost)
	//r.HandleFunc("/register", BestaHandler.createUser).Methods(http.MethodPost)findallusersubs

	// make a new subrouter when you want to grouping where path want to be protect and FindAllUserSubs
	authorized := r.NewRoute().Subrouter()
	authorized.Use(middlewares.SetMiddlewareAuthentication)
	authorized.HandleFunc("/besta", BestaHandler.findAll).Methods(http.MethodGet)
	authorized.HandleFunc("/besta/user/{id}", BestaHandler.UserFindSub).Methods(http.MethodGet)
	authorized.HandleFunc("/besta/register", BestaHandler.createBesta).Methods(http.MethodPost)
	authorized.HandleFunc("/besta/{id}", BestaHandler.findByID).Methods(http.MethodGet)
	//authorized.HandleFunc("/besta/findallusersubs{id}", BestaHandler.FindAllUserSubs).Methods(http.MethodGet)
	authorized.HandleFunc("/besta/{id}", BestaHandler.updateSubs).Methods(http.MethodPut)
	authorized.HandleFunc("/besta/{id}", BestaHandler.deleteSubs).Methods(http.MethodDelete)

	// make a new subrouter extends a authorized path
	uploadRequest := authorized.PathPrefix("/besta").Subrouter()
	uploadRequest.Use(middlewares.FileSizeLimiter) // use middleware to limited size when upload file
	uploadRequest.HandleFunc("/photo/{id}", BestaHandler.handlingPhoto).Methods(http.MethodPost)
}

// func (call *BestaHandler) loginSubs(w http.ResponseWriter, r *http.Request) {

// 	besta := new(models.Besta)
// 	//besta.Name = "King Starr"

// 	fmt.Println("%+v", besta)

// 	//fmt.Println("%+v", r)

// 	if err := json.NewDecoder(r.Body).Decode(&besta); err != nil {
// 		// w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser", "Error when trying to login, error is =>", err)
// 		// common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
// 		// return

// 		r.ParseForm()


// 		besta.Useremail = r.FormValue("email") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
// 		//besta.Password = r.FormValue("password") //string `gorm:"column:password;size:100;not null;" json:"password"`
// 	}

	

// 	err := common.Validate("login", besta)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		// common.LogError("loginUser 1", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	besta, err = call.bestaUsecase.Login(besta)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		// common.LogError("loginUser 2", "Error when trying to login, error is =>", err)
// 		common.Response(w, common.Message(false, err.Error(), nil))
// 		return
// 	}

// 	// token, err := auth.CreateToken(besta)
// 	// if err != nil {
// 	// 	common.LogError("loginUser 3", "Error when trying to generate token, error is =>", err)
// 	// 	common.Response(w, common.Message(false, "Opps.. something when wrong", nil))
// 	// 	return
// 	// }

// 	// allsubs := new(models.BestaWrapper)
// 	// copier.Copy(&allsubs, &besta)

// 	// finalResponse := map[string]interface{}{"Apps": "status", "User": allsubs, "data": map[string]interface{}{"token": token}}

// 	// common.Response(w, common.Message(true, "Success", finalResponse))

// 	common.Response(w, common.Message(true, "Success", map[string]interface{}{"token": "kingstarrtoken"}))
// }

func (call *BestaHandler) createBesta(w http.ResponseWriter, r *http.Request) {

	besta := new(models.Besta)

	
	if err := json.NewDecoder(r.Body).Decode(&besta); err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		//common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		//return

		r.ParseForm()

		

		besta.Fullname = r.FormValue("fullname") //"Kingstarr Alistring" // `gorm:"column:name;size:255;not null" json:"name"`
		besta.Useremail = r.FormValue("useremail") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
		besta.UserID = r.FormValue("userid") //string `gorm:"column:notlpn;not null" json:"no_tlpn"`
		besta.Service = r.FormValue("servicename") //string `gorm:"column:password;size:100;not null;" json:"password"`
		besta.ServiceID =  r.FormValue("serviceid") //string `gorm:"column:gender;size:15;not null" json:"gender"`
	}
		
	err := common.Validate("register", besta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	// val, _ := auth.TokenValid(r)
	// if val.Role != "admin" && besta.Role == "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	response, err := call.bestaUsecase.Create(besta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), response))
		return
	}

	common.Response(w, common.Message(true, "Register successfully", response))
	return
}

func (call *BestaHandler) findAll(w http.ResponseWriter, r *http.Request) {

	// val, err := auth.TokenValid(r)
	// if err != nil || val.Role != "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	result, err := call.bestaUsecase.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

func (call *BestaHandler) UserFindSub(w http.ResponseWriter, r *http.Request) {

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

	result, err := call.bestaUsecase.FindAllUS(IDSubs)
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
// func (call *BestaHandler) FindAllUserSubs(w http.ResponseWriter, r *http.Request) {

// 	// val, err := auth.TokenValid(r)
// 	// if err != nil || val.Role != "admin" {
// 	// 	w.WriteHeader(http.StatusForbidden)
// 	// 	common.Response(w, common.Message(false, "Access denied", nil))
// 	// 	return
// 	// }

	
// 	result := []*models.AllUserBesta{}

// 	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
// 		return
// 	}

// 	log.Println(IDSubs)

// 	result, err = call.bestaUsecase.FindAllUserSubs("g")
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		common.Response(w, common.Message(false, "result empty", nil))
// 		return
// 	}
// 	common.Response(w, common.Message(true, "Success", result))
// }


func (call *BestaHandler) findByID(w http.ResponseWriter, r *http.Request) {
	besta := new(models.BestaWrapper)

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	//allsubs, err := call.bestaUsecase.FindByID(IDSubs)
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

	besta, err = call.bestaUsecase.FindByID(IDSubs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", besta))
}

func (call *BestaHandler) updateSubs(w http.ResponseWriter, r *http.Request) {
	besta := new(models.Besta)

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&besta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	allsubs, err := call.bestaUsecase.FindByID(IDSubs)
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

	besta.ID = allsubs.ID
	err = common.Validate("update", besta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	besta, err = call.bestaUsecase.Update(besta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Update successfully", besta))
	return
}

func (call *BestaHandler) deleteSubs(w http.ResponseWriter, r *http.Request) {
	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	//allsubs, err := call.bestaUsecase.FindByID(IDSubs)
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

	err = call.bestaUsecase.Delete(IDSubs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. something when wrong", nil))
		return
	}

	common.Response(w, common.Message(true, "Success delete account", nil))
}

func (call *BestaHandler) handlingPhoto(w http.ResponseWriter, r *http.Request) {

	besta := new(models.Besta)

	IDSubs, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	//val, _ := auth.TokenValid(r)
	allsubs, err := call.bestaUsecase.FindByID(IDSubs)
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
	copier.Copy(&besta, &allsubs)
	err = call.bestaUsecase.UpdatePhoto(besta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success file was uploaded", filePath))
}

func (call *BestaHandler) handleFile(r *http.Request, besta *models.BestaWrapper, key string) (string, error) {
	file, handler, err := r.FormFile(key)
	if err != nil {
		common.LogError("handleFile", "Error When Handle File, error is => ", err)
		return "", err
	}

	defer file.Close()

	image := viper.GetString("file.path") + "/" + besta.Photo
	if _, err := os.Stat(image); !os.IsNotExist(err) && !strings.Contains(besta.Photo, "default.jpeg") {
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
