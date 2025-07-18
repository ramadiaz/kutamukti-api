package repositories

type CompRepositories interface {
	CountUsers() (int, error)
	CountComplaints() (int, error)
	CountComplaintsByStatus(status string) (int, error)
	CountSchedules() (int, error)
	CountAnnouncements() (int, error)
	CountUMKM() (int, error)
	CountUMKMProducts() (int, error)
	CountGalleries() (int, error)
	CountImages() (int, error)
	CountVideos() (int, error)
	CountNews() (int, error)
	CountFiles() (int, error)
}
