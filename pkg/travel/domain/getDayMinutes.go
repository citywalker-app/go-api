package traveldomain

func (t *Travel) GetDayMinutes(day uint8) uint16 {
	switch day {
	case 0:
		return uint16(t.Schedule.EndTime.Sub(t.Schedule.StartDateTime).Minutes())
	case t.Schedule.TotalDays:
		return uint16(t.Schedule.EndDateTime.Sub(t.Schedule.StartTime).Minutes())
	default:
		return uint16(t.Schedule.EndTime.Sub(t.Schedule.StartTime).Minutes())
	}
}
