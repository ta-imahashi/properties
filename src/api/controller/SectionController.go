package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ta-imahashi/properties/model"
	"github.com/ta-imahashi/properties/service"
)

type SectionController struct {
	SectionService service.SectionService
}

func NewSectionController(se service.SectionService) SectionController {
	return SectionController{SectionService: se}
}

func (ctrl *SectionController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": ctrl.SectionService.List(),
	})
}

func (ctrl *SectionController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"item": ctrl.SectionService.Find(c.Param("id")),
	})
}

func (ctrl *SectionController) GetSections(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": ctrl.SectionService.GetSections(c.Param("id")),
	})
}

func (ctrl *SectionController) Store(c *gin.Context) {
	section := model.Section{}
	c.ShouldBindJSON(&section)

	c.JSON(http.StatusOK, ctrl.SectionService.Create(section))
}

func (ctrl *SectionController) Update(c *gin.Context) {
	section := model.Section{}
	c.ShouldBindJSON(&section)
	section.Id, _ = strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, ctrl.SectionService.Update(section))
}

func (ctrl *SectionController) Destroy(c *gin.Context) {
	ctrl.SectionService.Delete(c.Param("id"))
	c.JSON(http.StatusNoContent, gin.H{})
}
