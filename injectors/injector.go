// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	userControllers "kutamukti-api/api/users/controllers"
	userRepositories "kutamukti-api/api/users/repositories"
	userServices "kutamukti-api/api/users/services"
	
	complaintControllers "kutamukti-api/api/complaint/controllers"
	complaintRepositories "kutamukti-api/api/complaint/repositories"
	complaintServices "kutamukti-api/api/complaint/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	userServices.NewComponentServices,
	userControllers.NewCompController,
)

var complaintFeatureSet = wire.NewSet(
	complaintRepositories.NewComponentRepository,
	complaintServices.NewComponentServices,
	complaintControllers.NewCompController,
)

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeComplaintController(db *gorm.DB, validate *validator.Validate) complaintControllers.CompControllers {
	wire.Build(complaintFeatureSet)
	return nil
}
