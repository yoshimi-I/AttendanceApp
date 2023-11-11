package response

import (
	"fmt"
	"time"
)

func FormatChange(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()
	return fmt.Sprintf("%02d：%02d", hour, minute)
	
}
