package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"inventory.com/catalog/internal/controller"
	"inventory.com/catalog/internal/handler/ginhandler"
	"inventory.com/catalog/internal/repository/memory"
)

var (
	categoryRepo    *memory.Category
	subCategoryRepo *memory.SubCategory
	productRepo     *memory.Product
)

var (
	categoryCtrl    *controller.CategoryController
	subCategoryCtrl *controller.SubCategoryController
	productCtrl     *controller.ProductController
)

func init() {
	initRepos()
	initControllers()
}

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	ginhandler.InitCategoryHandler(engine, categoryCtrl)
	ginhandler.InitSubCategoryHandler(engine, subCategoryCtrl)
	ginhandler.InitProductHandler(engine, productCtrl)
	if err := engine.Run(":8081"); err != nil {
		log.Fatalf("[server] Failed to start server: %v", err)
	}

}

func initRepos() {
	categoryRepo = memory.NewCategory()
	subCategoryRepo = memory.NewSubCategory()
	productRepo = memory.NewProduct()
}
func initControllers() {
	categoryCtrl = controller.NewCategoryController(categoryRepo)
	subCategoryCtrl = controller.NewSubCategoryController(subCategoryRepo, categoryCtrl)
	productCtrl = controller.NewProductController(productRepo, subCategoryCtrl)
}
