package helpers

import (
	"fmt"
	"time"
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
