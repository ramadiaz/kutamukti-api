package helpers

import "regexp"

func IsValidYouTubeURL(url string) bool {
	pattern := `^(https?://)?(www\.)?(youtube\.com/watch\?v=|youtu\.be/)[\w\-]{11}(&.*)?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(url)
}
