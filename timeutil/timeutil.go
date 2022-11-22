package timeutil

import (
	"strconv"
	"time"
)

func GetNowDateTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func GetNowDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func GetNowTime() string {
	now := time.Now()
	return now.Format("15:04:05")
}

func GetNowDateTimeReportName() string {
	now := time.Now()
	return now.Format("20060102-150405")
}

func GetNowByFormat(format string) string {
	now := time.Now()
	return now.Format(format)
}

func GetNowUnixSecond() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func GetNowUnixMillisecond() string {
	return strconv.Itoa(int(time.Now().UnixNano() / 1e6))
}

func GetNowUnixNanosecond() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}
