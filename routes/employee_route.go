package routes

import (
	"github.com/Amierza/employee-management/controller"
	"github.com/gin-gonic/gin"
)

func Employee(route *gin.Engine, employeeController controller.EmployeeController) {
	routes := route.Group("/api/employee")
	{
		routes.POST("/create-employee", employeeController.CreateEmployee)
		routes.GET("/get-all-employee", employeeController.GetAllEmployee)
		routes.POST("/update-employee", employeeController.UpdateProfile)
		routes.POST("/delete-employee", employeeController.DeleteEmployee)
	}
}
