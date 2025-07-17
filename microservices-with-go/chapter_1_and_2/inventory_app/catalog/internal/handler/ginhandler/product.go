package ginhandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"inventory.com/catalog/pkg/model"
)

type IProductController interface {
	Create(ctx context.Context, data *model.ProductBasic) (*model.ProductBasic, error)
	Update(ctx context.Context, id model.ProductID, data *model.ProductBasic) error
	Get(ctx context.Context, id model.ProductID) (*model.ProductInformation, error)
	GetAll(ctx context.Context) ([]*model.ProductInformation, error)
	Delete(ctx context.Context, id model.ProductID) (*model.ProductBasic, error)
}
type productHandler struct {
	ctrl IProductController
}

func (handler *productHandler) post(ctx *gin.Context) {
	var data model.ProductBasic
	if err := ctx.ShouldBindJSON(&data); err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}
	created, err := handler.ctrl.Create(ctx.Request.Context(), &data)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to create product")
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (handler *productHandler) update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid product ID")
		return
	}

	var data model.ProductBasic
	if err := ctx.ShouldBindJSON(&data); err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}
	data.ID = model.ProductID(id)
	err = handler.ctrl.Update(ctx.Request.Context(), data.ID, &data)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to update product")
		return
	}
	ctx.JSON(http.StatusAccepted, data)
}

func (handler *productHandler) getAll(ctx *gin.Context) {
	all, err := handler.ctrl.GetAll(ctx.Request.Context())
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to retrieve products")
		return
	}
	ctx.JSON(http.StatusOK, all)
}

func (handler *productHandler) get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid product ID")
		return
	}
	product, err := handler.ctrl.Get(ctx.Request.Context(), model.ProductID(id))
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to retrieve product")
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (handler *productHandler) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid product ID")
		return
	}
	_, err = handler.ctrl.Delete(ctx.Request.Context(), model.ProductID(id))
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to delete product")
		return
	}
	ctx.JSON(http.StatusNoContent, struct{}{})
}

func InitProductHandler(engine *gin.Engine, ctrl IProductController) {
	handler := &productHandler{ctrl: ctrl}
	router := engine.Group("/products")
	router.POST("", handler.post)
	router.PUT(":id", handler.update)
	router.GET("", handler.getAll)
	router.GET(":id", handler.get)
	router.DELETE(":id", handler.delete)
}
