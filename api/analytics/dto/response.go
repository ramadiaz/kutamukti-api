package dto

type AnalyticsResponse struct {
	TotalUsers                      int     `json:"total_users"`
	TotalUsersPercentChange         float64 `json:"total_users_percent_change"`
	TotalComplaints                 int     `json:"total_complaints"`
	TotalComplaintsPercentChange    float64 `json:"total_complaints_percent_change"`
	ComplaintsOpen                  int     `json:"complaints_open"`
	ComplaintsOpenPercentChange     float64 `json:"complaints_open_percent_change"`
	ComplaintsProcess               int     `json:"complaints_process"`
	ComplaintsProcessPercentChange  float64 `json:"complaints_process_percent_change"`
	ComplaintsClosed                int     `json:"complaints_closed"`
	ComplaintsClosedPercentChange   float64 `json:"complaints_closed_percent_change"`
	TotalSchedules                  int     `json:"total_schedules"`
	TotalSchedulesPercentChange     float64 `json:"total_schedules_percent_change"`
	TotalAnnouncements              int     `json:"total_announcements"`
	TotalAnnouncementsPercentChange float64 `json:"total_announcements_percent_change"`
	TotalUMKM                       int     `json:"total_umkm"`
	TotalUMKMPercentChange          float64 `json:"total_umkm_percent_change"`
	TotalUMKMProducts               int     `json:"total_umkm_products"`
	TotalUMKMProductsPercentChange  float64 `json:"total_umkm_products_percent_change"`
	TotalGalleries                  int     `json:"total_galleries"`
	TotalGalleriesPercentChange     float64 `json:"total_galleries_percent_change"`
	TotalImages                     int     `json:"total_images"`
	TotalImagesPercentChange        float64 `json:"total_images_percent_change"`
	TotalVideos                     int     `json:"total_videos"`
	TotalVideosPercentChange        float64 `json:"total_videos_percent_change"`
	TotalNews                       int     `json:"total_news"`
	TotalNewsPercentChange          float64 `json:"total_news_percent_change"`
	TotalFiles                      int     `json:"total_files"`
	TotalFilesPercentChange         float64 `json:"total_files_percent_change"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}
