package dto

import (
	// "encoding/json"
	// "fmt"
	"strings"
	"time"
)

type Schedule struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`

	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
}

type CustomTime time.Time

const customTimeLayout = "2006-01-02T15:04:05.00000-07:00"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(customTimeLayout, s)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	return []byte(`"` + t.Format(customTimeLayout) + `"`), nil
}

func (ct CustomTime) ToTime() time.Time {
	return time.Time(ct)
}

// func (h *Schedule) UnmarshalJSON(data []byte) error {
// 	type Alias Schedule
// 	aux := &struct {
// 		StartTime string `json:"start_time" binding:"required"`
// 		EndTime   string `json:"end_time" binding:"required"`
// 		*Alias
// 	}{
// 		Alias: (*Alias)(h),
// 	}

// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}

// 	layout := "2006-01-02"
// 	if aux.StartTime != "" {
// 		parsedStartTime, err := time.Parse(layout, aux.StartTime)
// 		if err != nil {
// 			return fmt.Errorf("invalid format for check_in_date: %v", err)
// 		}
// 		h.StartTime = parsedStartTime
// 	}

// 	if aux.EndTime != "" {
// 		parsedEndTime, err := time.Parse(layout, aux.EndTime)
// 		if err != nil {
// 			return fmt.Errorf("invalid format for check_out_date: %v", err)
// 		}
// 		h.EndTime = parsedEndTime
// 	}

// 	return nil
// }
