package puzzle

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	// PartAll means we want to solve both part 1 and part 2.
	PartAll = 0
	// Part1 means we want to solve part 1.
	Part1 = 1
	// Part2 means we want to solve part 2.
	Part2 = 2
)

// Solver is the main solver.
// All days and day parts will be called from here.
type Solver struct {
	Day  Day
	Part int
}

// NewSolver creates a new Solver for a specific day and part.
func NewSolver(day Day, part int) Solver {
	return Solver{
		Day:  day,
		Part: part,
	}
}

// Solve simply solves the puzzle, if it's possible.
func (s *Solver) Solve(value, file string) (string, error) {
	logrus.Debugf("Calling %T => Part%d", s.Day, s.Part)

	if value == "" && file == "" {
		return "", NoInputError{}
	}

	input, err := inputAsReader(value, file)
	if err != nil {
		logrus.Debugf("Input error: %s", err.Error())

		return "", err
	}

	defer input.Close()

	err = s.Day.SetInput(input)
	if err != nil {
		logrus.Debugf("SetInput error: %s", err.Error())

		return "", err
	}

	switch s.Part {
	case PartAll:
		part1, err := s.Day.Part1()
		if err != nil {
			return part1, err
		}

		part2, err := s.Day.Part2()

		return fmt.Sprintf(" -- Part1:\n%s\n\n -- Part2:\n%s", part1, part2), err
	case Part1:
		return s.Day.Part1()
	case Part2:
		return s.Day.Part2()
	}

	return "", UnkownPartError{Day: s.Day, Part: s.Part}
}
