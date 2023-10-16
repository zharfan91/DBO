package usecase

import (
	app "dbo/app/customer"
	"dbo/model"

	"golang.org/x/crypto/bcrypt"
)

type DBOUsecase struct {
	repo app.IDBORepository
}

func NewDBOUsecase(repo app.IDBORepository) app.IDBOUsecase {
	return &DBOUsecase{repo: repo}
}
func (u *DBOUsecase) CreateCustomer(request model.DBOCustomer) model.Respon {

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		res := model.Respon{
			Status:  400,
			Message: "Failed to Generate to hash",
		}
		return res
	}
	request.Password = string(bytes)
	err = u.repo.CreateCustomer(request)
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
func (u *DBOUsecase) GetCustomer(pagination model.Pagination) (*model.Pagination, error) {
	result, err := u.repo.GetCustomer(pagination)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *DBOUsecase) GetCustomerDetail(customerID string) (*model.Customer, error) {
	result, err := u.repo.GetCustomerDetail(customerID)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *DBOUsecase) DeleteCustomer(customerID string) model.Respon {
	res := u.repo.DeleteCustomer(customerID)
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
func (u *DBOUsecase) UpdateCustomer(request model.DBOCustomer, customerID string) model.Respon {

	err := u.repo.UpdateCustomer(request, customerID)
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
