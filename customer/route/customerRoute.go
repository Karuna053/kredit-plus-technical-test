package route

import (
	"kredit-plus/customer/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(route *gin.Engine) {
	routes := route.Group("/api/customer")
	{
		// Customer
		routes.POST("/create", controller.CreateCustomer)
	}
}
