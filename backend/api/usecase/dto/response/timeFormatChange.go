package response

import (
	"fmt"
	"time"
)

func FormatChange(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()
	return fmt.Sprintf("%d時%02d分", hour, minute)
}
