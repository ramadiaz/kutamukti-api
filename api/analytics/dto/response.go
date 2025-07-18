package dto

type AnalyticsResponse struct {
	TotalUsers         int `json:"total_users"`
	TotalComplaints    int `json:"total_complaints"`
	ComplaintsOpen     int `json:"complaints_open"`
	ComplaintsProcess  int `json:"complaints_process"`
	ComplaintsClosed   int `json:"complaints_closed"`
	TotalSchedules     int `json:"total_schedules"`
	TotalAnnouncements int `json:"total_announcements"`
	TotalUMKM          int `json:"total_umkm"`
	TotalUMKMProducts  int `json:"total_umkm_products"`
	TotalGalleries     int `json:"total_galleries"`
	TotalImages        int `json:"total_images"`
	TotalVideos        int `json:"total_videos"`
	TotalNews          int `json:"total_news"`
	TotalFiles         int `json:"total_files"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}
