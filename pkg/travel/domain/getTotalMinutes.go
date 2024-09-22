package traveldomain

import (
	"time"
)

func (t *Travel) GetTotalMinutes() int16 {
	totalMinutes := 0.0
	start := t.Schedule.StartDate
	end := t.Schedule.EndDate

	startDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.UTC)

	t.Schedule.TotalDays = uint8(endDate.Sub(startDate).Hours() / 24)

	// minutes of the first day
	totalMinutes += t.Schedule.EndTime.Sub(t.Schedule.StartDateTime).Minutes()

	// minutes of the middle days
	if t.Schedule.TotalDays > 1 {
		totalMinutes += t.Schedule.EndTime.Sub(t.Schedule.StartTime).Minutes() * float64(t.Schedule.TotalDays-1)
	}

	// minutes of the last day
	totalMinutes += t.Schedule.EndDateTime.Sub(t.Schedule.StartTime).Minutes()

	return int16(totalMinutes)
}
