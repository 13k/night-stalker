package time

import (
	"time"
)

type Travel struct {
	time.Time
}

func TravelFrom(from time.Time) Travel {
	return Travel{Time: from}
}

// T returns the underlying time.
func (t Travel) T() time.Time {
	return t.Time
}

// Ago travels time back by given number of years, months and days.
func (t Travel) Ago(years int, months int, days int) Travel {
	if years > 0 {
		years = -years
	}

	if months > 0 {
		months = -months
	}

	if days > 0 {
		days = -days
	}

	to := t.Time.AddDate(years, months, days)
	return TravelFrom(to)
}

// YearsAgo travels time back by given number of years.
func (t Travel) YearsAgo(years int) Travel {
	return t.Ago(years, 0, 0)
}

// MonthsAgo travels time back by given number of months.
func (t Travel) MonthsAgo(months int) Travel {
	return t.Ago(0, months, 0)
}

// DaysAgo travels time back by given number of days.
func (t Travel) DaysAgo(days int) Travel {
	return t.Ago(0, 0, days)
}

// BeginningOfDay travels time back to beginning of time's day (00:00:00).
func (t Travel) BeginningOfDay() Travel {
	to := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return TravelFrom(to)
}
