package puzzle

import "fmt"

// UnkownDayError occurs when given day is not defined.
type UnkownDayError struct {
	Day int
}

func (e UnkownDayError) Error() string {
	return fmt.Sprintf("Unknown day: %d", e.Day)
}

// NewMonth creates a new Month.
func NewMonth() Month {
	return Month{
		days: map[int]Day{},
	}
}

// Month manages daily solvers.
type Month struct {
	days map[int]Day
}

// Register solver for a day.
func (m *Month) Register(day int, solver Day) {
	m.days[day] = solver
}

// Get solver for a day.
func (m *Month) Get(day int) (Day, error) {
	selectedDay, ok := m.days[day]

	if !ok {
		return nil, UnkownDayError{Day: day}
	}

	if selectedDay == nil {
		return nil, NotImplementedError{}
	}

	return selectedDay, nil
}
