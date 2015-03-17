package timeutil

import (
	"time"
)

//  for example:
//  2015-03-17 12:37:06 +0800 CST

//03-17
func getDate(timeNow time.Time) string {
	times := timeNow.String()
	return times[5:11]
}

//2015
func getYear(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:4]
}

//2015-03
func getYearWithMon(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:7]
}

//17
func getDay(timeNow time.Time) string {
	times := timeNow.String()
	return times[8:10]
}

//2015-03-17
func getFullDate(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:11]
}

//2015-03-17 12
func getDateWithHour(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:13]
}

//2015-03-17 12:37
func getDateWithTime(timeNow time.Time) string {
	times := timeNow.String()
	return times[0:16]
}

//12:37
func getTime(timeNow time.Time) string {
	times := timeNow.String()
	return times[11:16]
}

//12:37:06
func getTimeWithSec(timeNow time.Time) string {
	times := timeNow.String()
	return times[11:19]
}
