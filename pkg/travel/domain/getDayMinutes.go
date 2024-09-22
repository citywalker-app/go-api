package traveldomain

func (t *Travel) GetDayMinutes(day uint8) int16 {
	switch day {
	case 0:
		return int16(t.Schedule.EndTime.Sub(t.Schedule.StartDateTime).Minutes())
	case t.Schedule.TotalDays:
		return int16(t.Schedule.EndDateTime.Sub(t.Schedule.StartTime).Minutes())
	default:
		return int16(t.Schedule.EndTime.Sub(t.Schedule.StartTime).Minutes())
	}
}
