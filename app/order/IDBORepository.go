package app

import "dbo/model"

type IDBORepository interface {
	CreateOrder(request model.DBOrder) error
	GetOrder(model.Pagination) (*model.Pagination, error)
	GetOrderDetail(string) (*model.DBOrder, error)
	DeleteOrder(customerID string) bool
	UpdateOrder(request model.DBOrder, customerID string) error
}
