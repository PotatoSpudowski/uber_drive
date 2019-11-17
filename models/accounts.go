package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "../utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserID uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("driverid = ?", account.DriverID).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "DriverID already in use by another account."), false
	}
	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}) {
	
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account")
	}

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
}