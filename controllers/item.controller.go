package controllers

import (
	"net/http"

	"aesthetic/models"
	"aesthetic/services"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemService services.ItemService
}

func NewItemController(itemService services.ItemService) *ItemController {
	return &ItemController{itemService}
}

func (ic *ItemController) GetAllItem(ctx *gin.Context) {
	itemCode := ctx.Query("item_code")
	items, err := ic.itemService.GetAllItem(itemCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"items": items}})
}

func (ic *ItemController) CreateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	itemResponse, err := ic.itemService.CreateItem(&item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"item": itemResponse}})
}
