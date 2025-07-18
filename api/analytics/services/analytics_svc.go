package services

import "kutamukti-api/api/analytics/dto"

type CompServices interface {
	GetAnalyticsData() (*dto.AnalyticsResponse, error)
}
