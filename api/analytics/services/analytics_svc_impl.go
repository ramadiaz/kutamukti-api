package services

import (
	"kutamukti-api/api/analytics/dto"
	"kutamukti-api/api/analytics/repositories"
)

type CompServicesImpl struct {
	repo repositories.CompRepositories
}

func NewComponentServices(compRepositories repositories.CompRepositories) CompServices {
	return &CompServicesImpl{
		repo: compRepositories,
	}
}

func (s *CompServicesImpl) GetAnalyticsData() (*dto.AnalyticsResponse, error) {
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

	return &dto.AnalyticsResponse{
		TotalUsers:         users,
		TotalComplaints:    complaints,
		ComplaintsOpen:     complaintsOpen,
		ComplaintsProcess:  complaintsProcess,
		ComplaintsClosed:   complaintsClosed,
		TotalSchedules:     schedules,
		TotalAnnouncements: announcements,
		TotalUMKM:          umkm,
		TotalUMKMProducts:  umkmProducts,
		TotalGalleries:     galleries,
		TotalImages:        images,
		TotalVideos:        videos,
		TotalNews:          news,
		TotalFiles:         files,
	}, nil
}
