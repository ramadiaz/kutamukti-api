// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"kutamukti-api/internal/auth/controllers"
	"kutamukti-api/internal/auth/services"
)

// Injectors from injector.go:

func InitializeAuthController(validate *validator.Validate) controllers.CompControllers {
	compServices := services.NewComponentServices(validate)
	compControllers := controllers.NewCompController(compServices)
	return compControllers
}

// injector.go:

var authFeatureSet = wire.NewSet(services.NewComponentServices, controllers.NewCompController)
