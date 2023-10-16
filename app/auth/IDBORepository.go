package app

import "dbo/model"

type IDBORepository interface {
	Login(request model.Auth) (model.DBOCustomer, error)
}
