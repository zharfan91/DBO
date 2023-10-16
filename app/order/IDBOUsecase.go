package app

import "dbo/model"

type IDBOUsecase interface {
	CreateOrder(model.DBOrder) model.Respon
	GetOrder(model.Pagination) (*model.Pagination, error)
	GetOrderDetail(string) (*model.DBOrder, error)
	DeleteOrder(customerID string) model.Respon
	UpdateOrder(model.DBOrder, string) model.Respon
}
