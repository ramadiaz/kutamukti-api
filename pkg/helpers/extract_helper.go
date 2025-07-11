package helpers

import (
	"fmt"
	"kutamukti-api/pkg/exceptions"
	"net/url"
	"regexp"
	"strings"
)

func ExtractYouTubeID(youtubeURL string) (string, *exceptions.Exception) {
	parsedURL, err := url.Parse(youtubeURL)
	if err != nil {
		return "", exceptions.NewValidationException(fmt.Errorf("invalid URL: %v", err))
	}

	switch parsedURL.Host {
	case "www.youtube.com", "youtube.com":

		if parsedURL.Path == "/watch" {
			query := parsedURL.Query()
			return query.Get("v"), nil
		}
	case "youtu.be":

		return strings.TrimPrefix(parsedURL.Path, "/"), nil
	}

	re := regexp.MustCompile(`(?:v=|\/)([0-9A-Za-z_-]{11})`)
	match := re.FindStringSubmatch(youtubeURL)
	if len(match) > 1 {
		return match[1], nil
	}

	return "", exceptions.NewValidationException(fmt.Errorf("no valid video ID found in URL"))
}
