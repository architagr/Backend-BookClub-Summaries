package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"inventory.com/inventory_gateway/internal/controller"
	"inventory.com/inventory_gateway/internal/gateway"
	"inventory.com/inventory_gateway/internal/handler/ginhandler"
)

var (
	categoryGatewayAddr = "http://localhost:8081" // Example address, adjust as needed
	categoryControler   *controller.CategoryController
)

func init() {
	categoryControler = controller.NewCategoryController(gateway.NewCategoryGateway(categoryGatewayAddr))
}
func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	ginhandler.RegisterCategoryRoutes(engine, categoryControler)
	if err := engine.Run(":8083"); err != nil {
		log.Fatalf("[server] Failed to start server: %v", err)
	}
}
