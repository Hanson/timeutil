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
func GetDateWithHour(timeNow string) string {
	return timeNow[0:13]
}

//2015-03-17 12:37
func GetDateWithTime(timeNow string) string {
	return timeNow[0:16]
}

//12:37
func GetTime(timeNow string) string {
	return timeNow[11:16]
}

//12:37:06
func GetTimeWithSec(timeNow string) string {
	return timeNow[11:19]
}

//03-17
func GetSDate(timeNow string) string {
	return timeNow[5:11]
}

//2015
func GetSYear(timeNow string) string {
	return timeNow[0:4]
}

//2015-03
func GetSYearWithMon(timeNow string) string {
	return timeNow[0:7]
}

//17
func GetSDay(timeNow string) string {
	return timeNow[8:10]
}

//2015-03-17
func GetSFullDate(timeNow string) string {
	return timeNow[0:11]
}

//2015-03-17 12
func GetSDateWithHour(timeNow string) string {
	return timeNow[0:13]
}

//2015-03-17 12:37
func GetSDateWithTime(timeNow string) string {
	return timeNow[0:16]
}

//12:37
func GetSTime(timeNow string) string {
	return timeNow[11:16]
}

//12:37:06
func GetSTimeWithSec(timeNow string) string {
	return timeNow[11:19]
}
