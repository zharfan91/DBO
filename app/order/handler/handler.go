package handler

import (
	app "dbo/app/order"
	"dbo/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DBOHandler struct {
	usecase app.IDBOUsecase
}

func NewOrderHandler(usecase app.IDBOUsecase) app.IDBOHandler {
	return &DBOHandler{usecase: usecase}
}

// @Summary Creat Order
// @Tags Order
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param request body model.Order false "request"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/order [post]
func (h DBOHandler) CreateOrder(c *gin.Context) {
	var request model.DBOrder
	err := c.Bind(&request)
	if err != nil {
		log.Println("err handler login:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		return
	}

	res := h.usecase.CreateOrder(request)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
	return
}

// @Summary Get All Order
// @Tags Order
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/order [get]
func (h DBOHandler) GetOrder(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	pagination := model.Pagination{
		Page:  page,
		Limit: limit,
	}
	res, err := h.usecase.GetOrder(pagination)
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

// @Summary Get Detail Order
// @Tags Order
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param Order_ID path string true "Order ID"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/order/{Order_ID} [get]
func (h DBOHandler) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("id")
	res, err := h.usecase.GetOrderDetail(orderID)
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

// @Summary Delete Order
// @Tags Order
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param Order_ID path string true "Order ID"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/order/{Order_ID} [delete]
func (h DBOHandler) DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	res := h.usecase.DeleteOrder(orderID)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
	return
}

// @Summary Update Order
// @Tags Order
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param Order_ID path string true "Order ID"
// @Param request body model.Order false "request"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /dbo/Order/{Order_ID} [put]
func (h DBOHandler) UpdateOrder(c *gin.Context) {
	var request model.DBOrder
	orderID := c.Param("id")
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		return
	}

	res := h.usecase.UpdateOrder(request, orderID)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
	return
}
