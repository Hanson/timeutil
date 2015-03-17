package timeutil

import (
	"time"
)

//  for example:
//  2015-03-17 12:37:06 +0800 CST

//03-17
func GetDate(timeNow time.Time) string {
	times := timeNow.String()
	return times[5:11]
}

//2015
func GetYear(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:4]
}

//2015-03
func GetYearWithMon(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:7]
}

//17
func GetDay(timeNow time.Time) string {
	times := timeNow.String()
	return times[8:10]
}

//2015-03-17
func GetFullDate(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:11]
}

//2015-03-17 12
func GetDateWithHour(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:13]
}

//2015-03-17 12:37
func GetDateWithTime(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:16]
}

//12:37
func GetTime(timeNow time.Time) string {
	times := timeNow.String()
	return times[11:16]
}

//12:37:06
func GetTimeWithSec(timeNow time.Time) string {
	times := timeNow.String()
	return times[11:19]
}
