package common

import (
	"singleaf/apps/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

// LogError use logging error
func LogError(function string, str string, err error) {
	log.Printf("[%v] %v : %v", function, str, err)
}

// Message wrapper a message response
func Message(status bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message, "data": data}
}

// Response for print a response
func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

// Hash use for hashing password
func Hash(dataparam string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(dataparam), bcrypt.DefaultCost)
}

//VerifyPassword use for password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Encrypt use for encypt
func Encrypt(apps *models.Apps) (*models.Apps, error) {
	hashedData, err := Hash(apps.Role)
	if err != nil {
		return nil, err
	}
	apps.Role = string(hashedData)
	return apps, nil
}

// Validate use for validate register, login, and update
func Validate(action string, apps *models.Apps) error {
	switch strings.ToLower(action) {
	case "register":
		if apps.Name == "" {
			return errors.New("Required Name")
		}
		if apps.Email == "" {
			return errors.New("Required Email")
		}
		if apps.Description == "" {
			return errors.New("Required Description")
		}

		if err := checkmail.ValidateFormat(apps.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if apps.Email == "" {
			return errors.New("Required Email")
		}
		if apps.Name == "" {
			return errors.New("Required Password")
		}
		if apps.Description == "" {
			return errors.New("Required Description")
		}
		if err := checkmail.ValidateFormat(apps.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "update":
		if apps.Name == "" {
			return errors.New("Required Name")
		}
		if apps.Email == "" {
			return errors.New("Required Email")
		}
		if apps.Description == "" {
			return errors.New("Required Description")
		}
		if err := checkmail.ValidateFormat(apps.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		if apps.Name == "" {
			return errors.New("Required Name")
		}
		if apps.Email == "" {
			return errors.New("Required Email")
		}
		if apps.Description == "" {
			return errors.New("Required Description")
		}

		if err := checkmail.ValidateFormat(apps.Email); err != nil {
			return errors.New("Invalid Email")
		}
		
		return nil
	}
}
