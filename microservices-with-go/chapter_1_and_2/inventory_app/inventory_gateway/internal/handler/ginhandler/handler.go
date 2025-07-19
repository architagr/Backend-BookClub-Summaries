package ginhandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"inventory.com/catalog/pkg/model"
)

type ICategoryControler interface {
	Create(ctx context.Context, data *model.Category) (*model.Category, error)
	Update(ctx context.Context, id model.CategoryID, data *model.Category) (*model.Category, error)
	// Get(ctx context.Context, id model.CategoryID) (*model.Category, error)
}

type CategoryHandler struct {
	controller ICategoryControler
}

func NewCategoryHandler(controller ICategoryControler) *CategoryHandler {
	return &CategoryHandler{controller: controller}
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var data *model.Category
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category data"})
		return
	}

	data, err := h.controller.Create(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

func (h *CategoryHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var data *model.Category
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category data"})
		return
	}

	data, err = h.controller.Update(ctx, model.CategoryID(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

func RegisterCategoryRoutes(engine *gin.Engine, ctrl ICategoryControler) {
	handler := NewCategoryHandler(ctrl)
	categoryRouter := engine.Group("/categories")
	{
		categoryRouter.POST("/", handler.Create)
		categoryRouter.PUT("/:id", handler.Update)
	}
}
