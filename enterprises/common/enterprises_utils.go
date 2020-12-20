package common

import (
	"singleaf/enterprises/models"
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
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword use for password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Encrypt use for encypt
func Encrypt(enterprises *models.Enterprises) (*models.Enterprises, error) {
	hashedPassword, err := Hash(enterprises.EnterpriseHash)
	if err != nil {
		return nil, err
	}
	enterprises.EnterpriseHash = string(hashedPassword)
	return enterprises, nil
}

// Validate use for validate register, login, and update
func Validate(action string, enterprises *models.Enterprises) error {
	switch strings.ToLower(action) {
	case "register":
		if enterprises.EnterpriseName == "" {
			return errors.New("Required Name")
		}
		if enterprises.EnterpriseEmail == "" {
			return errors.New("Required Email")
		}
		
		if enterprises.CompanyTag == "" {
			return errors.New("Required Tag")
		}
		if enterprises.Description == "" {
			return errors.New("Required Desription")
		}
		if enterprises.Address == "" {
			return errors.New("Required Address")
		}
		
		if enterprises.VerifyID == "" {
			return errors.New("Required VerifyID")
		}
		// if err := checkmail.ValidateFormat(enterprises.EnterpriseEmail); err != nil {
		// 	return errors.New("Invalid email address")
		// }
		return nil
	case "login":
		if enterprises.EnterpriseEmail == "" {
			return errors.New("Required Email")
		}
		if enterprises.EnterpriseHash == "" {
			return errors.New("Required Password")
		}
		if err := checkmail.ValidateFormat(enterprises.EnterpriseEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "update":
		if enterprises.EnterpriseName == "" {
			return errors.New("Required Name")
		}
		if enterprises.EnterpriseEmail == "" {
			return errors.New("Required Email")
		}
		if enterprises.NoTlpn == "" {
			return errors.New("Required Phone Number")
		}
		if enterprises.EnterpriseHash == "" {
			return errors.New("Required Password")
		}
		if enterprises.Description == "" {
			return errors.New("Required Gender")
		}
		if enterprises.Address == "" {
			return errors.New("Required Address")
		}
		if err := checkmail.ValidateFormat(enterprises.EnterpriseEmail); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		if enterprises.EnterpriseName == "" {
			return errors.New("Required Name")
		}
		if enterprises.EnterpriseEmail == "" {
			return errors.New("Required Email")
		}
		if enterprises.NoTlpn == "" {
			return errors.New("Required Phone Number")
		}
		if enterprises.EnterpriseHash == "" {
			return errors.New("Required Password")
		}
		if enterprises.Description == "" {
			return errors.New("Required Gender")
		}
		if enterprises.Address == "" {
			return errors.New("Required Address")
		}
		if err := checkmail.ValidateFormat(enterprises.EnterpriseEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
