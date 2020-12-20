package common

import (
	"singleaf/besta/models"
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
func Encrypt(besta *models.Besta) (*models.Besta, error) {
	hashedPassword, err := Hash(besta.Useremail)
	if err != nil {
		return nil, err
	}
	besta.Useremail = string(hashedPassword)
	return besta, nil
}

// Validate use for validate register, login, and update
func Validate(action string, besta *models.Besta) error {
	switch strings.ToLower(action) {
	case "register":
		if besta.Fullname == "" {
			return errors.New("Required User Fullname")
		}
		if besta.Useremail == "" {
			return errors.New("Required User Email")
		}
		if besta.UserID == "" {
			return errors.New("Required User ID")
		}
		if besta.Service == "" {
			return errors.New("Required Service Name")
		}
		if besta.ServiceID == "" {
			return errors.New("Required Service ID")
		}
				
		if err := checkmail.ValidateFormat(besta.Useremail); err != nil {
			return errors.New("Invalid email address")
		}
		return nil
	case "login":
		// if besta.Fullname == "" {
		// 	return errors.New("Required User Fullname")
		// }
		if besta.Useremail == "" {
			return errors.New("Required User Email")
		}
		// if besta.UserID == "" {
		// 	return errors.New("Required User ID")
		// }
		// if besta.Service == "" {
		// 	return errors.New("Required Service Name")
		// }
		// if besta.ServiceID == "" {
		// 	return errors.New("Required Service ID")
		// }
		if err := checkmail.ValidateFormat(besta.Useremail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "update":
		if besta.Fullname == "" {
			return errors.New("Required User Fullname")
		}
		if besta.Useremail == "" {
			return errors.New("Required User Email")
		}
		if besta.UserID == "" {
			return errors.New("Required User ID")
		}
		if besta.Service == "" {
			return errors.New("Required Service Name")
		}
		if besta.ServiceID == "" {
			return errors.New("Required Service ID")
		}
		if err := checkmail.ValidateFormat(besta.Useremail); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		if besta.Fullname == "" {
			return errors.New("Required User Fullname")
		}
		if besta.Useremail == "" {
			return errors.New("Required User Email")
		}
		if besta.UserID == "" {
			return errors.New("Required User ID")
		}
		if besta.Service == "" {
			return errors.New("Required Service Name")
		}
		if besta.ServiceID == "" {
			return errors.New("Required Service ID")
		}
		if err := checkmail.ValidateFormat(besta.Useremail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
