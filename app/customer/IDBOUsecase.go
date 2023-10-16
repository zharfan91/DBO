package app

import "dbo/model"

type IDBOUsecase interface {
	CreateCustomer(model.DBOCustomer) model.Respon
	GetCustomer(model.Pagination) (*model.Pagination, error)
	GetCustomerDetail(string) (*model.Customer, error)
	DeleteCustomer(customerID string) model.Respon
	UpdateCustomer(model.DBOCustomer, string) model.Respon
}
