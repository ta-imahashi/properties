package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ta-imahashi/properties/model"
	"github.com/ta-imahashi/properties/service"
)

type PropertyController struct {
	PropertyService service.PropertyService
}

func NewPropertyController(se service.PropertyService) PropertyController {
	return PropertyController{PropertyService: se}
}

func (ctrl *PropertyController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": ctrl.PropertyService.List(),
	})
}

func (ctrl *PropertyController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"item": ctrl.PropertyService.Find(c.Param("id")),
	})
}

func (ctrl *PropertyController) Store(c *gin.Context) {
	property := model.Property{}
	c.ShouldBindJSON(&property)

	c.JSON(http.StatusOK, ctrl.PropertyService.Create(property))
}

func (ctrl *PropertyController) Update(c *gin.Context) {
	property := model.Property{}
	c.ShouldBindJSON(&property)
	property.Id, _ = strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, ctrl.PropertyService.Update(property))
}

func (ctrl *PropertyController) Destroy(c *gin.Context) {
	ctrl.PropertyService.Delete(c.Param("id"))
	c.JSON(http.StatusNoContent, gin.H{})
}
