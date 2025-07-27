package ginhandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"inventory.com/catalog/pkg/model"
)

type ISubCategoryController interface {
	Create(ctx context.Context, data *model.SubCategoryBasic) (*model.SubCategoryBasic, error)
	Update(ctx context.Context, id model.SubCategoryID, data *model.SubCategoryBasic) error
	Get(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryDetails, error)
	GetAll(ctx context.Context) ([]*model.SubCategoryDetails, error)
	Delete(ctx context.Context, id model.SubCategoryID) (*model.SubCategoryBasic, error)
}

type subCategoryHandler struct {
	ctrl ISubCategoryController
}

func (handler *subCategoryHandler) post(ctx *gin.Context) {
	var data model.SubCategoryBasic
	if err := ctx.ShouldBindJSON(&data); err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}
	created, err := handler.ctrl.Create(ctx.Request.Context(), &data)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to create subcategory")
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (handler *subCategoryHandler) update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid subcategory ID")
		return
	}
	var data model.SubCategoryBasic
	if err := ctx.ShouldBindJSON(&data); err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}
	data.BaseInfo.ID = model.SubCategoryID(id)
	err = handler.ctrl.Update(ctx.Request.Context(), data.BaseInfo.ID, &data)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to update subcategory")
		return
	}

	ctx.JSON(http.StatusAccepted, data)
}

func (handler *subCategoryHandler) getAll(ctx *gin.Context) {
	all, err := handler.ctrl.GetAll(ctx.Request.Context())
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to retrieve subcategories")
		return
	}
	ctx.JSON(http.StatusOK, all)
}

func (handler *subCategoryHandler) get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid subcategory ID")
		return
	}
	subCategory, err := handler.ctrl.Get(ctx.Request.Context(), model.SubCategoryID(id))
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to retrieve subcategory")
		return
	}
	ctx.JSON(http.StatusOK, subCategory)
}

func (handler *subCategoryHandler) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "invalid subcategory ID")
		return
	}
	_, err = handler.ctrl.Delete(ctx.Request.Context(), model.SubCategoryID(id))
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "failed to delete subcategory")
		return
	}
	ctx.JSON(http.StatusNoContent, struct{}{})
}

func InitSubCategoryHandler(engine *gin.Engine, ctrl ISubCategoryController) {
	handler := &subCategoryHandler{ctrl: ctrl}
	router := engine.Group("/subcategories")
	router.POST("", handler.post)
	router.PUT(":id", handler.update)
	router.GET("", handler.getAll)
	router.GET(":id", handler.get)
	router.DELETE(":id", handler.delete)
}
