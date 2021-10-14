package models

import (
	"time"

	"github.com/badoux/checkmail"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
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

func validateEmail(email string) error {
	return checkmail.ValidateFormat(email)
}

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
