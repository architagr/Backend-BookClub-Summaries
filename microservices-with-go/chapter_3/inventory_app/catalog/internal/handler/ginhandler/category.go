package ginhandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"inventory.com/catalog/pkg/model"
)

// handleError is a helper to send consistent error responses
func handleError(ctx *gin.Context, status int, msg string) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"error": msg,
	})
}

type ICategoryController interface {
	Create(ctx context.Context, data *model.Category) (*model.Category, error)
	Update(ctx context.Context, id model.CategoryID, data *model.Category) error
	Get(ctx context.Context, id model.CategoryID) (*model.Category, error)
	GetAll(ctx context.Context) ([]*model.Category, error)
	Delete(ctx context.Context, id model.CategoryID) (*model.Category, error)
}

type categoryHandler struct {
	ctrl ICategoryController
}

func (handler *categoryHandler) post(ctx *gin.Context) {
	var data model.Category
	if err := ctx.ShouldBindJSON(&data); err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}

	created, err := handler.ctrl.Create(ctx.Request.Context(), &data)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to create category")
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (handler *categoryHandler) update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid category ID")
		return
	}

	var data model.Category
	if err := ctx.ShouldBindJSON(&data); err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}
	data.ID = model.CategoryID(id)
	err = handler.ctrl.Update(ctx.Request.Context(), data.ID, &data)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to update category")
		return
	}
	ctx.JSON(http.StatusAccepted, data)
}

func (handler *categoryHandler) getAll(ctx *gin.Context) {
	all, err := handler.ctrl.GetAll(ctx.Request.Context())
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to retrieve categories")
		return
	}
	ctx.JSON(http.StatusOK, all)
}

func (handler *categoryHandler) get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid category ID")
		return
	}

	category, err := handler.ctrl.Get(ctx.Request.Context(), model.CategoryID(id))
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to retrieve category")
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (handler *categoryHandler) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid category ID")
		return
	}

	_, err = handler.ctrl.Delete(ctx.Request.Context(), model.CategoryID(id))
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to delete category")
		return
	}
	ctx.JSON(http.StatusNoContent, struct{}{})
}

func InitCategoryHandler(engine *gin.Engine, ctrl ICategoryController) {
	handler := &categoryHandler{ctrl: ctrl}

	categoryRouterGroup := engine.Group("/categories")
	{
		categoryRouterGroup.POST("", handler.post)
		categoryRouterGroup.PUT("/:id", handler.update)
		categoryRouterGroup.GET("", handler.getAll)
		categoryRouterGroup.GET("/:id", handler.get)
		categoryRouterGroup.DELETE("/:id", handler.delete)
	}
}
