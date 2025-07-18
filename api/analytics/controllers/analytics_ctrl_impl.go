package controllers

import (
	"kutamukti-api/api/analytics/dto"
	"kutamukti-api/api/analytics/services"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var analyticsCache struct {
	sync.Mutex
	data *dto.AnalyticsResponse
	time time.Time
}

const analyticsCacheDuration = 10 * time.Minute

func startAnalyticsCacheRefresher(s services.CompServices) {
	go func() {
		for {
			data, err := s.GetAnalyticsData()
			if err == nil {
				analyticsCache.Lock()
				analyticsCache.data = data
				analyticsCache.time = time.Now()
				analyticsCache.Unlock()
			} else {
				log.Printf("[analytics cache] failed to refresh: %v", err)
			}
			time.Sleep(analyticsCacheDuration)
		}
	}()
}

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	startAnalyticsCacheRefresher(compServices)
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) GetAnalyticsData(ctx *gin.Context) {
	analyticsCache.Lock()
	if analyticsCache.data != nil && time.Since(analyticsCache.time) < analyticsCacheDuration {
		data := analyticsCache.data
		analyticsCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	analyticsCache.Unlock()

	data, err := h.services.GetAnalyticsData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get analytics data",
			Body:    nil,
		})
		return
	}

	analyticsCache.Lock()
	analyticsCache.data = data
	analyticsCache.time = time.Now()
	analyticsCache.Unlock()

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}
