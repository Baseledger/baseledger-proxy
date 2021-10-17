package types

import (
	"errors"
	"fmt"
	"time"

	"github.com/badoux/checkmail"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/token"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uuid.UUID `json:"id" sql:"id"`
	CreatedAt time.Time `sql:"not null;default:now()" json:"created_at,omitempty"`
	Email     string    `json:"email" validate:"required" sql:"email"`
	Password  string    `json:"password" validate:"required" sql:"password"`
}

func (u *User) Create() bool {
	pwdHash, err := generatePasswordHash(u.Password)
	if err != nil {
		logger.Errorf("errors while creating new entry %v\n", err.Error())
		return false
	}
	err = validateEmail(u.Email)
	if err != nil {
		logger.Errorf("errors while creating new entry %v\n", err.Error())
		return false
	}
	existingUser := getUserByEmail(u.Email)
	if existingUser != nil {
		logger.Error("user already exists %v\n")
		return false
	}
	u.Password = pwdHash
	if dbutil.Db.GetConn().NewRecord(u) {
		result := dbutil.Db.GetConn().Create(&u)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			logger.Errorf("errors while creating new entry %v\n", errors)
			return false
		}
		return rowsAffected > 0
	}

	return false
}

func (u *User) Login() (string, error) {
	existingUser := getUserByEmail(u.Email)
	if existingUser == nil {
		errorMsg := fmt.Sprintf("user with email %v is not registered", u.Email)
		logger.Error(errorMsg)
		return "", errors.New(errorMsg)
	}

	passwordMatch := checkPasswordHash(u.Password, existingUser.Password)
	if !passwordMatch {
		errorMsg := fmt.Sprintf("passwords not matching for user %v", u.Email)
		logger.Error(errorMsg)
		return "", errors.New(errorMsg)
	}

	return token.GetToken(u.Email)
}

func getUserByEmail(email string) *User {
	db := dbutil.Db.GetConn()
	var user User
	res := db.First(&user, "email = ?", email)

	if res.Error != nil {
		logger.Infof("User with email %v not found", email)
		return nil
	}

	return &user
}

func validateEmail(email string) error {
	return checkmail.ValidateFormat(email)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logger.Errorf("passwords not matchin error %v\n", err.Error())
	}
	return err == nil
}

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
