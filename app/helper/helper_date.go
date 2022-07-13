package helper

import (
	"time"
)

const (
	YYYYMMDD       = "2006-01-02"
	DDMMYYYY       = "02/01/2006"
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	YYYY           = "2006"
)

func GetCurrentDateYYYYMMDD() string {
	now := time.Now().UTC()
	return now.Format(YYYYMMDD)
}

func FormatDateYYYYMMDD(date time.Time) string {
	now := date.UTC()
	return now.Format(YYYYMMDD)
}

func GetCurrentDateDDMMYYYY() string {
	now := time.Now().UTC()
	return now.Format(DDMMYYYY)
}
func FormatDateDDMMYYYY(date time.Time) string {
	now := date.UTC()
	return now.Format(DDMMYYYY)
}

func GetCurrentDateDDMMYYYYhhmmss() string {
	now := time.Now().UTC()
	return now.Format(DDMMYYYYhhmmss)
}
func FormatDateDDMMYYYYhhmmss(date time.Time) string {
	now := date.UTC()
	return now.Format(DDMMYYYYhhmmss)
}

func GetCurrentYear() string {
	now := time.Now().UTC()
	return now.Format(YYYY)
}
