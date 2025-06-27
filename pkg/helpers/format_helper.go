package helpers

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var indonesianMonths = [...]string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

func FormatIndonesianTime(t time.Time) string {
	day := t.Day()
	month := indonesianMonths[t.Month()-1]
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	return fmt.Sprintf("%02d %s %d %02d:%02d", day, month, year, hour, minute)
}

func FormatSlug(text string) string {
	slug := strings.ToLower(text)
	slug = removeDiacritics(slug)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	uniqueSuffix := fmt.Sprintf("%d%s", time.Now().Unix(), randomString(2))
	return fmt.Sprintf("%s-%s", slug, uniqueSuffix)
}

func removeDiacritics(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case 'à', 'á', 'â', 'ã', 'ä', 'å':
			b.WriteRune('a')
		case 'è', 'é', 'ê', 'ë':
			b.WriteRune('e')
		case 'ì', 'í', 'î', 'ï':
			b.WriteRune('i')
		case 'ò', 'ó', 'ô', 'õ', 'ö':
			b.WriteRune('o')
		case 'ù', 'ú', 'û', 'ü':
			b.WriteRune('u')
		case 'ý', 'ÿ':
			b.WriteRune('y')
		case 'ñ':
			b.WriteRune('n')
		case 'ç':
			b.WriteRune('c')
		default:
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				b.WriteRune(r)
			} else {
				b.WriteRune(' ')
			}
		}
	}
	return b.String()
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}
