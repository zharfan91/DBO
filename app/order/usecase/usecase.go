package usecase

import (
	app "dbo/app/order"
	"dbo/model"
)

type DBOUsecase struct {
	repo app.IDBORepository
}

func NewOrderUsecase(repo app.IDBORepository) app.IDBOUsecase {
	return &DBOUsecase{repo: repo}
}
func (u *DBOUsecase) CreateOrder(request model.DBOrder) model.Respon {

	err := u.repo.CreateOrder(request)
	if err != nil {
		res := model.Respon{
			Status:  400,
			Message: "Failed to insert record to a database",
		}
		return res
	}
	res := model.Respon{
		Status:  200,
		Message: "the request was successful",
	}
	return res
}
func (u *DBOUsecase) GetOrder(pagination model.Pagination) (*model.Pagination, error) {
	result, err := u.repo.GetOrder(pagination)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *DBOUsecase) GetOrderDetail(orderID string) (*model.DBOrder, error) {
	result, err := u.repo.GetOrderDetail(orderID)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *DBOUsecase) DeleteOrder(orderID string) model.Respon {
	res := u.repo.DeleteOrder(orderID)
	if !res {
		resp := model.Respon{
			Status:  400,
			Message: "We couldn't delete your pages. Try again",
		}
		return resp
	}
	resp := model.Respon{
		Status:  200,
		Message: "the request was successful",
	}
	return resp
}
func (u *DBOUsecase) UpdateOrder(request model.DBOrder, orderID string) model.Respon {

	err := u.repo.UpdateOrder(request, orderID)
	if err != nil {
		res := model.Respon{
			Status:  400,
			Message: "Failed to insert record to a database",
		}
		return res
	}
	res := model.Respon{
		Status:  200,
		Message: "the request was successful",
	}
	return res
}
