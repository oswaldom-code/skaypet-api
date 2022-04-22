package dto

import (
	"fmt"
	"time"
)

func DateFormat(t time.Time) string {
	return t.Format("2006-01-02")
}

func AgeInYearsMonthsDays(t time.Time) string {
	now := time.Now()
	years := now.Year() - t.Year()
	months := int(now.Month()) - int(t.Month())
	days := now.Day() - t.Day()
	if months < 0 {
		years--
		months = 12 + months
	}
	if days < 0 {
		months--
		days = 30 + days
	}
	return fmt.Sprintf("%d años, %d meses, %d dias", years, months, days)
}
