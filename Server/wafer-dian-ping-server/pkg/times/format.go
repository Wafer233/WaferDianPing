package times

import "time"

func FormatTime(t *time.Time, layout string) string {
	if t == nil {
		return ""
	}
	return t.Format(layout)
}

func ToTime(timeStr string) *time.Time {

	if timeStr == "" {
		now := time.Now()
		return &now
	}

	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return nil
	}
	return &t
}
