package repository

import (
	app "dbo/app/auth"
	"dbo/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBORepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) app.IDBORepository {
	return &DBORepository{db: db}
}
func (r *DBORepository) Login(request model.Auth) (model.DBOCustomer, error) {
	var res model.DBOCustomer
	r.db.First(&res, "username = ?", request.Username)
	match := CheckPasswordHash(request.Password, string(res.Password))
	if !match {
		return res, errors.New("Wrong user name or password ")
	}
	return res, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
