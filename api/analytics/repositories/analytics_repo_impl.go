package repositories

import (
	"kutamukti-api/models"
	"time"

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

func lastMonthRange() (time.Time, time.Time) {
	now := time.Now()
	firstOfThisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastMonth := firstOfThisMonth.AddDate(0, -1, 0)
	firstOfLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, now.Location())
	lastOfLastMonth := firstOfThisMonth.Add(-time.Nanosecond)
	return firstOfLastMonth, lastOfLastMonth
}

func (r *CompRepositoriesImpl) CountUsersLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Users{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountComplaintsLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Complaints{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountComplaintsByStatusLastMonth(status string) (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Complaints{}).Where("status = ? AND created_at >= ? AND created_at <= ?", status, start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountSchedulesLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Schedules{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountAnnouncementsLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Announcements{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountUMKMLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.UMKM{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountUMKMProductsLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.UMKMProduct{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountGalleriesLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.ImageGalleries{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountImagesLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Images{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountVideosLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Videos{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountNewsLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.News{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}

func (r *CompRepositoriesImpl) CountFilesLastMonth() (int, error) {
	var count int64
	start, end := lastMonthRange()
	err := r.DB.Model(&models.Files{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&count).Error
	return int(count), err
}
