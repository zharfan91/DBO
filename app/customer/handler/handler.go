package handler

import (
	app "dbo/app/customer"
	"dbo/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DBOHandler struct {
	usecase app.IDBOUsecase
}

func NewDBOHandler(usecase app.IDBOUsecase) app.IDBOHandler {
	return &DBOHandler{usecase: usecase}
}

// @Summary Creat Customer
// @Tags Customer
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param request body model.Customer false "request"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/customer [post]
func (h DBOHandler) CreateCustomer(c *gin.Context) {
	var request model.DBOCustomer
	err := c.Bind(&request)
	if err != nil {
		log.Println("err handler login:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		return
	}

	res := h.usecase.CreateCustomer(request)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
	return
}

// @Summary Get All Customer
// @Tags Customer
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/customer [get]
func (h DBOHandler) GetCustomer(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	pagination := model.Pagination{
		Page:  page,
		Limit: limit,
	}
	res, err := h.usecase.GetCustomer(pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "the request was successful",
		"data":    res,
	})
	return
}

// @Summary Get Detail Customer
// @Tags Customer
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param Customer_ID path string true "Customer ID"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/customer/{Customer_ID} [get]
func (h DBOHandler) GetCustomerDetail(c *gin.Context) {
	customerID := c.Param("id")
	res, err := h.usecase.GetCustomerDetail(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "the request was successful",
		"data":    res,
	})
}

// @Summary Delete Customer
// @Tags Customer
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param Customer_ID path string true "Customer ID"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/customer/{Customer_ID} [delete]
func (h DBOHandler) DeleteCustomer(c *gin.Context) {
	customerID := c.Param("id")
	res := h.usecase.DeleteCustomer(customerID)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
	return
}

// @Summary Update Customer
// @Tags Customer
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param Customer_ID path string true "Customer ID"
// @Param request body model.Customer false "request"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/customer/{Customer_ID} [put]
func (h DBOHandler) UpdateCustomer(c *gin.Context) {
	var request model.DBOCustomer
	customerID := c.Param("id")
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		return
	}

	res := h.usecase.UpdateCustomer(request, customerID)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
	return
}
