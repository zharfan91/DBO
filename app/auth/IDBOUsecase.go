package app

import "dbo/model"

type IDBOUsecase interface {
	Login(model.Auth) (string, error)
}
