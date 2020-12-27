package main

import (
	"github.com/gin-gonic/gin"
	"restGo/simple_rest"
)

func main() {
	memoryStorage := simple_rest.NewMemoryStorage()
	handler := simple_rest.NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee:id", handler.GetEmployee)
	router.PUT("/employee:id", handler.UpdateEmployee)
	router.DELETE("/employee:id", handler.DeleteEmployee)

	router.Run()

}
