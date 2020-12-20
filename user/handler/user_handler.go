package handler

import (
	"singleaf/auth"
	"singleaf/middlewares"
	"singleaf/user"
	"singleaf/user/common"
	"singleaf/user/models"
	//"singleaf/subscriptions"
	//"singleaf/subscriptions/subscriber"
	//"bytes"
	"encoding/json"
	"fmt"
	"io"
	//"io/ioutil"
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

// UserHandler struct use for get funcntion in business logic
type UserHandler struct {
	userUsecase user.UserUsecase
}

// CreateHandler use for handling request
func CreateHandler(r *mux.Router, usecase user.UserUsecase) {
	userHandler := UserHandler{usecase}

	r.HandleFunc("/login", userHandler.loginUser).Methods(http.MethodPost)
	r.HandleFunc("/register", userHandler.createUser).Methods(http.MethodPost)

	// make a new subrouter when you want to grouping where path want to be protect and nah
	authorized := r.NewRoute().Subrouter()
	authorized.Use(middlewares.SetMiddlewareAuthentication)
	authorized.HandleFunc("/user", userHandler.findAll).Methods(http.MethodGet)
	authorized.HandleFunc("/user/{id}", userHandler.findByID).Methods(http.MethodGet)
	authorized.HandleFunc("/user/{id}", userHandler.updateUser).Methods(http.MethodPut)
	authorized.HandleFunc("/user/{id}", userHandler.deleteUser).Methods(http.MethodDelete)

	// make a new subrouter extends a authorized path
	uploadRequest := authorized.PathPrefix("/user").Subrouter()
	uploadRequest.Use(middlewares.FileSizeLimiter) // use middleware to limited size when upload file
	uploadRequest.HandleFunc("/photo/{id}", userHandler.handlingPhoto).Methods(http.MethodPost)
}

func (call *UserHandler) loginUser(w http.ResponseWriter, r *http.Request) {

	user := new(models.User)
	//user.Name = "King Starr"

	fmt.Println("%+v", user)

	//fmt.Println("%+v", r)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// common.LogError("loginUser", "Error when trying to login, error is =>", err)
		// common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		// return

		r.ParseForm()


		user.Email = r.FormValue("email") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
		user.Password = r.FormValue("password") //string `gorm:"column:password;size:100;not null;" json:"password"`
	}

	

	err := common.Validate("login", user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// common.LogError("loginUser 1", "Error when trying to login, error is =>", err)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	user, err = call.userUsecase.Login(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// common.LogError("loginUser 2", "Error when trying to login, error is =>", err)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		common.LogError("loginUser 3", "Error when trying to generate token, error is =>", err)
		common.Response(w, common.Message(false, "Opps.. something when wrong", nil))
		return
	}

	users := new(models.UserWrapper)
	copier.Copy(&users, &user)

	//var subscriptionInfo []byte //map[string]interface{}{}
	//var msg interface{}
	fmt.Println("Starting to consume intrernal api...")

	//var s = strconv.FormatUint(uint64(user.ID),10)

	// var serverPath = "http://localhost:5799/subscriptions/user/"


	// request, _ := http.NewRequest("GET", serverPath+s, bytes.NewBuffer(subscriptionInfo)) //, bytes.NewBuffer(jsonValue))
	// request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Authorization", "Bearer "+token)
	// client := &http.Client{}
	// response, err := client.Do(request)
	
	// if err != nil {
 //    	fmt.Printf("The HTTP request failed with error %s\n", err)
 //    	//subscriptionInfo = string(subscriptionInfo)
	// } else {
 //    	data, _ := ioutil.ReadAll(response.Body)
    	
	// 	err = json.Unmarshal(data, &msg)
	// 	if err != nil {
	// 		//http.Error(w, err.Error(), 500)
	// 		//return
	// 	}
 //    	fmt.Println(string(data))
 //    	//subscriptionInfo = []byte(msg)
	// }

	

	//subscriptionInfo,_ := subscriber.FindAllUS(1) //call new subscriptionrepo link //map[string]interface{}{"Name": "Proprietor"}

	finalResponse := map[string]interface{}{"User": users, "data": map[string]interface{}{"token": token}} //"Subscriptions": msg, 

	common.Response(w, common.Message(true, "Success", finalResponse))
}

func stringify(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func (call *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {

	user := new(models.User)

	//user.Name = "King Starr"
	//r.ParseForm()

	//fmt.Println(r.Form)




	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		//common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		//return

		r.ParseForm()




		user.Name = r.FormValue("name") //"Kingstarr Alistring" // `gorm:"column:name;size:255;not null" json:"name"`
		user.Email = r.FormValue("email") //string `gorm:"column:email;size:100;not null;unique" json:"email"`
		user.NoTlpn = r.FormValue("no_tlpn") //string `gorm:"column:notlpn;not null" json:"no_tlpn"`
		user.Password = r.FormValue("password") //string `gorm:"column:password;size:100;not null;" json:"password"`
		user.Gender =  r.FormValue("gender") //string `gorm:"column:gender;size:15;not null" json:"gender"`
		user.Address =   r.FormValue("address") //string `gorm:"column:address;size:300;not null" json:"address"`
		user.Role =    "user"
		user.Photo ="default.jpeg"
	}

	err := common.Validate("register", user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	// val, _ := auth.TokenValid(r)
	// if val.Role != "admin" && user.Role == "admin" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	common.Response(w, common.Message(false, "Access denied", nil))
	// 	return
	// }

	response, err := call.userUsecase.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), response))
		return
	}

	common.Response(w, common.Message(true, "Register successfully", response))
	return
}

func (call *UserHandler) findAll(w http.ResponseWriter, r *http.Request) {

	val, err := auth.TokenValid(r)
	if err != nil || val.Role != "admin" {
		w.WriteHeader(http.StatusForbidden)
		common.Response(w, common.Message(false, "Access denied", nil))
		return
	}

	result, err := call.userUsecase.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		common.Response(w, common.Message(false, "result empty", nil))
		return
	}
	common.Response(w, common.Message(true, "Success", result))
}

func (call *UserHandler) findByID(w http.ResponseWriter, r *http.Request) {
	user := new(models.UserWrapper)

	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.userUsecase.FindByID(IDUser)
	if val.Role != "admin" {
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

	user, err = call.userUsecase.FindByID(IDUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success", user))
}

func (call *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)

	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.userUsecase.FindByID(IDUser)
	if val.Role != "admin" {
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

	user.ID = users.ID
	err = common.Validate("update", user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	user, err = call.userUsecase.Update(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Update successfully", user))
	return
}

func (call *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.userUsecase.FindByID(IDUser)
	if val.Role != "admin" {
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

	err = call.userUsecase.Delete(IDUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Oops.. something when wrong", nil))
		return
	}

	common.Response(w, common.Message(true, "Success delete account", nil))
}

func (call *UserHandler) handlingPhoto(w http.ResponseWriter, r *http.Request) {

	user := new(models.User)

	IDUser, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error(), nil))
		return
	}

	val, _ := auth.TokenValid(r)
	users, err := call.userUsecase.FindByID(IDUser)
	if val.Role != "admin" {
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

	users.Photo = filePath
	copier.Copy(&user, &users)
	err = call.userUsecase.UpdatePhoto(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, err.Error(), nil))
		return
	}

	common.Response(w, common.Message(true, "Success file was uploaded", filePath))
}

func (call *UserHandler) handleFile(r *http.Request, user *models.UserWrapper, key string) (string, error) {
	file, handler, err := r.FormFile(key)
	if err != nil {
		common.LogError("handleFile", "Error When Handle File, error is => ", err)
		return "", err
	}

	defer file.Close()

	image := viper.GetString("file.path") + "/" + user.Photo
	if _, err := os.Stat(image); !os.IsNotExist(err) && !strings.Contains(user.Photo, "default.jpeg") {
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
