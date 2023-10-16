package handler

import (
	app "dbo/app/auth"
	"dbo/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DBOHandler struct {
	usecase app.IDBOUsecase
}

func NewAuthHandler(usecase app.IDBOUsecase) app.IDBOHandler {
	return &DBOHandler{usecase: usecase}
}

// @Summary Login
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body model.Auth false "request"
// @Success   200  {object} model.Respon
// @Failure   400  {object} model.Respon
// @Router /login [post]
func (h DBOHandler) Login(c *gin.Context) {
	var request model.Auth
	err := c.Bind(&request)
	if err != nil {
		log.Println("err handler login:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		return
	}
	res, err := h.usecase.Login(request)
	if err != nil {
		log.Println("err handler login:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "the request was successful",
		"token":   res,
	})
	log.Println("login", res)
}
