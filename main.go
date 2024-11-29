package main

import (
	"log"
	"os"

	"github.com/Amierza/employee-management/cmd"
	"github.com/Amierza/employee-management/config"
	"github.com/Amierza/employee-management/controller"
	"github.com/Amierza/employee-management/middleware"
	"github.com/Amierza/employee-management/repository"
	"github.com/Amierza/employee-management/routes"
	"github.com/Amierza/employee-management/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Command(db)
		return
	}

	var (
		employeeRepository repository.EmployeeRepository = repository.NewEmployeeRepository(db)
		employeeService    service.EmployeeService       = service.NewEmployeeService(employeeRepository)
		employeeController controller.EmployeeController = controller.NewEmployeeController(employeeService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.Employee(server, employeeController)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
