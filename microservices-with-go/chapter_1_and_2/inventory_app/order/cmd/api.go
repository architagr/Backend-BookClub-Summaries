package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"inventory.com/order/internal/controller"
	"inventory.com/order/internal/handler/ginhandler"
	"inventory.com/order/internal/repository/memory"
)

var repo *memory.Order
var ctrl *controller.OrderController

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	ginhandler.RegisterOrderRoutes(engine, ctrl)
	if err := engine.Run(":8082"); err != nil {
		log.Fatalf("[server] Failed to start server: %v", err)
	}
}

func initRepository() {
	repo = memory.New()
}
func initController() {
	ctrl = controller.NewOrderController(repo)
}
func init() {
	initRepository()
	initController()
}
