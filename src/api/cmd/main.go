package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ta-imahashi/properties/config"
	"github.com/ta-imahashi/properties/infrastructure"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.EnvLoad()

	r := gin.Default()
	infrastructure.DefineRoutes(r)
	r.Run(":80")
}
