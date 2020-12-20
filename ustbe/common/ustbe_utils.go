package common

import (
	"singleaf/ustbe/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	//"github.com/badoux/checkmail"
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
func Encrypt(ustbe *models.Ustbe) (*models.Ustbe, error) {
	hashedPassword, err := Hash(ustbe.Useremail)
	if err != nil {
		return nil, err
	}
	ustbe.Useremail = string(hashedPassword)
	return ustbe, nil
}

// Validate use for validate register, login, and update
func Validate(action string, ustbe *models.Ustbe) error {
	switch strings.ToLower(action) {
	case "register":
		if ustbe.UserID == "" {
			return errors.New("Required User ID")
		}
		if ustbe.CompanyTag == "" {
			return errors.New("Required Publisher Tag")
		}
		if ustbe.CompanyID == "" {
			return errors.New("Required Publisher ID")
		}
		if ustbe.Description == "" {
			return errors.New("Required Publisher Name")
		}
				
		// if err := checkmail.ValidateFormat(ustbe.Useremail); err != nil {
		// 	return errors.New("Invalid email address")
		// }
		return nil
	case "login":
		// if ustbe.Fullname == "" {
		// 	return errors.New("Required User Fullname")
		// }
		// if ustbe.Useremail == "" {
		// 	return errors.New("Required User Email")
		// }
		if ustbe.UserID == "" {
			return errors.New("Required User ID")
		}
		if ustbe.CompanyTag == "" {
			return errors.New("Required Service Tag")
		}
		if ustbe.CompanyID == "" {
			return errors.New("Required Service ID")
		}
		// if err := checkmail.ValidateFormat(ustbe.Useremail); err != nil {
		// 	return errors.New("Invalid Email")
		// }
		return nil
	case "update":
		
		if ustbe.UserID == "" {
			return errors.New("Required User ID")
		}
		if ustbe.CompanyTag == "" {
			return errors.New("Required Service Tag")
		}
		if ustbe.CompanyID == "" {
			return errors.New("Required Service ID")
		}
		// if err := checkmail.ValidateFormat(ustbe.Useremail); err != nil {
		// 	return errors.New("Invalid Email")
		// }

		return nil
	default:
		
		if ustbe.UserID == "" {
			return errors.New("Required User ID")
		}
		if ustbe.CompanyTag == "" {
			return errors.New("Required Service Tag")
		}
		if ustbe.CompanyID == "" {
			return errors.New("Required Service ID")
		}
		// if err := checkmail.ValidateFormat(ustbe.Useremail); err != nil {
		// 	return errors.New("Invalid Email")
		// }
		return nil
	}
}
