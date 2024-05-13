package main

import (
	"strings"
	"time"
)

func formatDateTime(date time.Time) string {

	timeStr := strings.ReplaceAll(date.Format("2006-01-02 15:04:05"), " ", "T")
	return timeStr
}
