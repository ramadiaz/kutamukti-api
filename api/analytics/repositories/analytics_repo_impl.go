package repositories

import (
	"kutamukti-api/models"

	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
	DB *gorm.DB
}

func NewComponentRepository(db *gorm.DB) CompRepositories {
	return &CompRepositoriesImpl{DB: db}
}

func (r *CompRepositoriesImpl) CountUsers() (int, error) {
	var count int64
	err := r.DB.Model(&models.Users{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountComplaints() (int, error) {
	var count int64
	err := r.DB.Model(&models.Complaints{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountComplaintsByStatus(status string) (int, error) {
	var count int64
	err := r.DB.Model(&models.Complaints{}).Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountSchedules() (int, error) {
	var count int64
	err := r.DB.Model(&models.Schedules{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountAnnouncements() (int, error) {
	var count int64
	err := r.DB.Model(&models.Announcements{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountUMKM() (int, error) {
	var count int64
	err := r.DB.Model(&models.UMKM{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountUMKMProducts() (int, error) {
	var count int64
	err := r.DB.Model(&models.UMKMProduct{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountGalleries() (int, error) {
	var count int64
	err := r.DB.Model(&models.ImageGalleries{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountImages() (int, error) {
	var count int64
	err := r.DB.Model(&models.Images{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountVideos() (int, error) {
	var count int64
	err := r.DB.Model(&models.Videos{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountNews() (int, error) {
	var count int64
	err := r.DB.Model(&models.News{}).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountFiles() (int, error) {
	var count int64
	err := r.DB.Model(&models.Files{}).Count(&count).Error
	return int(count), err
}
