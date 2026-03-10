package routes

import (
	"building-safety/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/customers", controllers.GetCustomers)
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers/:id", controllers.GetCustomer)
	r.PUT("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)

	return r

}
