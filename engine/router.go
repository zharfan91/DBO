package engine

import (
	customerHandler "dbo/app/customer/handler"
	customerRepo "dbo/app/customer/repository"
	customerUsecase "dbo/app/customer/usecase"

	authHandler "dbo/app/auth/handler"
	authRepo "dbo/app/auth/repository"
	authUsecase "dbo/app/auth/usecase"

	orderHandler "dbo/app/order/handler"
	orderRepo "dbo/app/order/repository"
	orderUsecase "dbo/app/order/usecase"

	"dbo/middleware"

	_ "dbo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.Use(gin.Recovery())

	// jwtService := middleware.NewJWTService()

	RepositoryCustomer := customerRepo.NewDBORepository(db)
	UsecaseCustomer := customerUsecase.NewDBOUsecase(RepositoryCustomer)
	HandlerCustomer := customerHandler.NewDBOHandler(UsecaseCustomer)

	RepositoryAuth := authRepo.NewAuthRepository(db)
	UsecaseAuth := authUsecase.NewAuthUsecase(RepositoryAuth)
	HandlerAuth := authHandler.NewAuthHandler(UsecaseAuth)

	RepositoryOrder := orderRepo.NewOrderRepository(db)
	UsecaseOrder := orderUsecase.NewOrderUsecase(RepositoryOrder)
	HandlerOrder := orderHandler.NewOrderHandler(UsecaseOrder)

	// Routes
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/login", HandlerAuth.Login)

	dbo := r.Group("/dbo", middleware.AuthorizeUserMiddleware())
	{
		//customer
		dbo.POST("/customer", HandlerCustomer.CreateCustomer)
		dbo.GET("/customer", HandlerCustomer.GetCustomer)
		dbo.GET("/customer/:id", HandlerCustomer.GetCustomerDetail)
		dbo.DELETE("/customer/:id", HandlerCustomer.DeleteCustomer)
		dbo.PUT("/customer/:id", HandlerCustomer.UpdateCustomer)

		//order
		dbo.POST("/order", HandlerOrder.CreateOrder)
		dbo.GET("/order", HandlerOrder.GetOrder)
		dbo.GET("/order/:id", HandlerOrder.GetOrderDetail)
		dbo.DELETE("/order/:id", HandlerOrder.DeleteOrder)
		dbo.PUT("/order/:id", HandlerOrder.UpdateOrder)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
