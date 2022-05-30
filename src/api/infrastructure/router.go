package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/ta-imahashi/properties/controller"
	"github.com/ta-imahashi/properties/db"
	"github.com/ta-imahashi/properties/repository"
	"github.com/ta-imahashi/properties/service"
)

func DefineRoutes(r gin.IRouter) {
	cnf := db.NewConfig()
	dynamo := db.NewDynamodb(cnf)
	counterRe := repository.NewCounterRepository(dynamo)

	propertyRe := repository.NewPropertyRepository(dynamo)
	propertySe := service.NewPropertyService(propertyRe, counterRe)
	propertyCo := controller.NewPropertyController(propertySe)

	sectionRe := repository.NewSectionRepository(dynamo)
	sectionSe := service.NewSectionService(sectionRe, counterRe)
	sectionCo := controller.NewSectionController(sectionSe)

	// co := main.Initialize()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthcheck.",
		})
	})

	v1 := r.Group("/v1")
	{
		properties := v1.Group("/properties")
		{
			properties.GET("", propertyCo.Index)
			properties.POST("", propertyCo.Store)
			properties.GET("/:id", propertyCo.Show)
			properties.PUT("/:id", propertyCo.Update)
			properties.DELETE("/:id", propertyCo.Destroy)
			properties.GET("/:id/sections", sectionCo.GetSections)
		}

		sections := v1.Group("/sections")
		{
			sections.GET("", sectionCo.Index)
			sections.POST("", sectionCo.Store)
			sections.GET("/:id", sectionCo.Show)
			sections.PUT("/:id", sectionCo.Update)
			sections.DELETE("/:id", sectionCo.Destroy)
		}
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v2/pong",
			})
		})
	}
}
