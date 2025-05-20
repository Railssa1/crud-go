package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	errors_api "github.com/Railssa1/crud-go/src/config/errors"
	"github.com/Railssa1/crud-go/src/config/logger"
)

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *errors_api.ApiErrors
}

func (ud *UserDomain) CreateUser() *errors_api.ApiErrors {
	logger.Info("Init CreateUser domain")

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
