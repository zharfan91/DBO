package app

import "dbo/model"

type IDBORepository interface {
	CreateCustomer(request model.DBOCustomer) error
	GetCustomer(model.Pagination) (*model.Pagination, error)
	GetCustomerDetail(string) (*model.Customer, error)
	DeleteCustomer(customerID string) bool
	UpdateCustomer(request model.DBOCustomer, customerID string) error
}
