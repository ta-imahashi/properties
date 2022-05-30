//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ta-imahashi/properties/controller"
	"github.com/ta-imahashi/properties/db"
	"github.com/ta-imahashi/properties/repository"
	"github.com/ta-imahashi/properties/service"
)

func Initialize() controller.PropertyController {
	wire.Build(controller.NewPropertyController, service.NewPropertyService, repository.NewPropertyRepository, db.NewDynamodb, db.NewConfig)
	return controller.PropertyController{}
}
