package util

import "time"

func FormatTime(str string) time.Time {
	var rs time.Time
	if rs, err := time.Parse("2006-01-02T15:04:05-0700", str+"00"); err == nil {
		return rs
	}
	if rs, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", str); err == nil {
		return rs
	}
	if rs, err := time.Parse("2006-01-02T15:04:05-0700", str); err == nil {
		return rs
	}
	rs, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Now()
	}
	return rs
}

func FormatTimeError(str string) (*time.Time, error) {
	var rs time.Time
	rs, err := time.Parse("2006-01-02T15:04:05-0700", str+"00");
	if err == nil {
		return &rs, nil
	}
	rs, err = time.Parse("2006-01-02 15:04:05 +0000 UTC", str)
	if err == nil {
		return &rs, nil
	}
	rs, err = time.Parse("2006-01-02T15:04:05-0700", str)
	if err == nil {
		return &rs, nil
	}
	rs, err = time.Parse(time.RFC3339, str)
	if err == nil {
		return &rs, nil
	}
	return nil, err
}

func GetMonth(date time.Time) (time.Time, time.Time) {
	firstOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth, lastOfMonth
}

func GetTimeDay(date time.Time) (time.Time, time.Time) {
	startTime := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endTime := startTime.AddDate(0, 0, 1).Add(-time.Nanosecond)
	return startTime, endTime
}
