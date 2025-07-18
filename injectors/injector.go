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

	scheduleControllers "kutamukti-api/api/schedule/controllers"
	scheduleRepositories "kutamukti-api/api/schedule/repositories"
	scheduleServices "kutamukti-api/api/schedule/services"

	announcementControllers "kutamukti-api/api/announcement/controllers"
	announcementRepositories "kutamukti-api/api/announcement/repositories"
	announcementServices "kutamukti-api/api/announcement/services"

	umkmControllers "kutamukti-api/api/umkm/controllers"
	umkmRepositories "kutamukti-api/api/umkm/repositories"
	umkmServices "kutamukti-api/api/umkm/services"

	galleryControllers "kutamukti-api/api/gallery/controllers"
	galleryRepositories "kutamukti-api/api/gallery/repositories"
	galleryServices "kutamukti-api/api/gallery/services"

	storageControllers "kutamukti-api/api/storages/controllers"
	storageRepositories "kutamukti-api/api/storages/repositories"
	storageServices "kutamukti-api/api/storages/services"

	newsControllers "kutamukti-api/api/news/controllers"
	newsRepositories "kutamukti-api/api/news/repositories"
	newsServices "kutamukti-api/api/news/services"

	analyticsControllers "kutamukti-api/api/analytics/controllers"
	analyticsRepositories "kutamukti-api/api/analytics/repositories"
	analyticsServices "kutamukti-api/api/analytics/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"google.golang.org/api/drive/v3"
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

var scheduleFeatureSet = wire.NewSet(
	scheduleRepositories.NewComponentRepository,
	scheduleServices.NewComponentServices,
	scheduleControllers.NewCompController,
)

var announcementFeatureSet = wire.NewSet(
	announcementRepositories.NewComponentRepository,
	announcementServices.NewComponentServices,
	announcementControllers.NewCompController,
)

var umkmFeatureSet = wire.NewSet(
	umkmRepositories.NewComponentRepository,
	umkmServices.NewComponentServices,
	umkmControllers.NewCompController,
)

var galleryFeatureSet = wire.NewSet(
	galleryRepositories.NewComponentRepository,
	galleryServices.NewComponentServices,
	galleryControllers.NewCompController,
)

var storageFeatureSet = wire.NewSet(
	storageRepositories.NewComponentRepository,
	storageServices.NewComponentServices,
	storageControllers.NewCompController,
)

var newsFeatureSet = wire.NewSet(
	newsRepositories.NewComponentRepository,
	newsServices.NewComponentServices,
	newsControllers.NewCompController,
)

var analyticsFeatureSet = wire.NewSet(
	analyticsRepositories.NewComponentRepository,
	analyticsServices.NewComponentServices,
	analyticsControllers.NewCompController,
)

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeComplaintController(db *gorm.DB, validate *validator.Validate) complaintControllers.CompControllers {
	wire.Build(complaintFeatureSet)
	return nil
}

func InitializeScheduleController(db *gorm.DB, validate *validator.Validate) scheduleControllers.CompControllers {
	wire.Build(scheduleFeatureSet)
	return nil
}

func InitializeAnnouncementController(db *gorm.DB, validate *validator.Validate) announcementControllers.CompControllers {
	wire.Build(announcementFeatureSet)
	return nil
}

func InitializeUMKMController(db *gorm.DB, validate *validator.Validate) umkmControllers.CompControllers {
	wire.Build(umkmFeatureSet)
	return nil
}

func InitializeGalleryController(db *gorm.DB, validate *validator.Validate) galleryControllers.CompControllers {
	wire.Build(galleryFeatureSet)
	return nil
}

func InitializeStorageController(db *gorm.DB, validate *validator.Validate, drive *drive.Service) storageControllers.CompControllers {
	wire.Build(storageFeatureSet)
	return nil
}

func InitializeNewsController(db *gorm.DB, validate *validator.Validate) newsControllers.CompControllers {
	wire.Build(newsFeatureSet)
	return nil
}

func InitializeAnalyticsController(db *gorm.DB) analyticsControllers.CompControllers {
	wire.Build(analyticsFeatureSet)
	return nil
}
