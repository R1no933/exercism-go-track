package meetup

import "time"

// WeekSchedule is an alias for defining dates based on weeks.
type WeekSchedule int

// Aliases for WeekSchedule coerced into actual days of the month.
const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = -6
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	if wSched == Last {
		month++
	}
	t := time.Date(year, month, int(wSched), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(wDay-t.Weekday()+7)%7
}
