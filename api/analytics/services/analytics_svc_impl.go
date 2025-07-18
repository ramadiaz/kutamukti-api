package services

import (
	"kutamukti-api/api/analytics/dto"
	"kutamukti-api/api/analytics/repositories"
	"math"
)

type CompServicesImpl struct {
	repo repositories.CompRepositories
}

func NewComponentServices(compRepositories repositories.CompRepositories) CompServices {
	return &CompServicesImpl{
		repo: compRepositories,
	}
}

func percentChange(current, last int) float64 {
	if last == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return math.Round(((float64(current)-float64(last))/math.Max(float64(last), 1))*100*100) / 100
}

func (s *CompServicesImpl) GetAnalyticsData() (*dto.AnalyticsResponse, error) {
	// Current counts
	users, err := s.repo.CountUsers()
	if err != nil {
		return nil, err
	}
	complaints, err := s.repo.CountComplaints()
	if err != nil {
		return nil, err
	}
	complaintsOpen, err := s.repo.CountComplaintsByStatus("open")
	if err != nil {
		return nil, err
	}
	complaintsProcess, err := s.repo.CountComplaintsByStatus("process")
	if err != nil {
		return nil, err
	}
	complaintsClosed, err := s.repo.CountComplaintsByStatus("closed")
	if err != nil {
		return nil, err
	}
	schedules, err := s.repo.CountSchedules()
	if err != nil {
		return nil, err
	}
	announcements, err := s.repo.CountAnnouncements()
	if err != nil {
		return nil, err
	}
	umkm, err := s.repo.CountUMKM()
	if err != nil {
		return nil, err
	}
	umkmProducts, err := s.repo.CountUMKMProducts()
	if err != nil {
		return nil, err
	}
	galleries, err := s.repo.CountGalleries()
	if err != nil {
		return nil, err
	}
	images, err := s.repo.CountImages()
	if err != nil {
		return nil, err
	}
	videos, err := s.repo.CountVideos()
	if err != nil {
		return nil, err
	}
	news, err := s.repo.CountNews()
	if err != nil {
		return nil, err
	}
	files, err := s.repo.CountFiles()
	if err != nil {
		return nil, err
	}

	// Last month counts
	usersLast, err := s.repo.CountUsersLastMonth()
	if err != nil {
		return nil, err
	}
	complaintsLast, err := s.repo.CountComplaintsLastMonth()
	if err != nil {
		return nil, err
	}
	complaintsOpenLast, err := s.repo.CountComplaintsByStatusLastMonth("open")
	if err != nil {
		return nil, err
	}
	complaintsProcessLast, err := s.repo.CountComplaintsByStatusLastMonth("process")
	if err != nil {
		return nil, err
	}
	complaintsClosedLast, err := s.repo.CountComplaintsByStatusLastMonth("closed")
	if err != nil {
		return nil, err
	}
	schedulesLast, err := s.repo.CountSchedulesLastMonth()
	if err != nil {
		return nil, err
	}
	announcementsLast, err := s.repo.CountAnnouncementsLastMonth()
	if err != nil {
		return nil, err
	}
	umkmLast, err := s.repo.CountUMKMLastMonth()
	if err != nil {
		return nil, err
	}
	umkmProductsLast, err := s.repo.CountUMKMProductsLastMonth()
	if err != nil {
		return nil, err
	}
	galleriesLast, err := s.repo.CountGalleriesLastMonth()
	if err != nil {
		return nil, err
	}
	imagesLast, err := s.repo.CountImagesLastMonth()
	if err != nil {
		return nil, err
	}
	videosLast, err := s.repo.CountVideosLastMonth()
	if err != nil {
		return nil, err
	}
	newsLast, err := s.repo.CountNewsLastMonth()
	if err != nil {
		return nil, err
	}
	filesLast, err := s.repo.CountFilesLastMonth()
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticsResponse{
		TotalUsers:                      users,
		TotalUsersPercentChange:         percentChange(users, usersLast),
		TotalComplaints:                 complaints,
		TotalComplaintsPercentChange:    percentChange(complaints, complaintsLast),
		ComplaintsOpen:                  complaintsOpen,
		ComplaintsOpenPercentChange:     percentChange(complaintsOpen, complaintsOpenLast),
		ComplaintsProcess:               complaintsProcess,
		ComplaintsProcessPercentChange:  percentChange(complaintsProcess, complaintsProcessLast),
		ComplaintsClosed:                complaintsClosed,
		ComplaintsClosedPercentChange:   percentChange(complaintsClosed, complaintsClosedLast),
		TotalSchedules:                  schedules,
		TotalSchedulesPercentChange:     percentChange(schedules, schedulesLast),
		TotalAnnouncements:              announcements,
		TotalAnnouncementsPercentChange: percentChange(announcements, announcementsLast),
		TotalUMKM:                       umkm,
		TotalUMKMPercentChange:          percentChange(umkm, umkmLast),
		TotalUMKMProducts:               umkmProducts,
		TotalUMKMProductsPercentChange:  percentChange(umkmProducts, umkmProductsLast),
		TotalGalleries:                  galleries,
		TotalGalleriesPercentChange:     percentChange(galleries, galleriesLast),
		TotalImages:                     images,
		TotalImagesPercentChange:        percentChange(images, imagesLast),
		TotalVideos:                     videos,
		TotalVideosPercentChange:        percentChange(videos, videosLast),
		TotalNews:                       news,
		TotalNewsPercentChange:          percentChange(news, newsLast),
		TotalFiles:                      files,
		TotalFilesPercentChange:         percentChange(files, filesLast),
	}, nil
}
