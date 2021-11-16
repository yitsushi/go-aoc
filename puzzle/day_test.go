package puzzle_test

import (
	"bufio"
	"io"
	"strings"

	"github.com/yitsushi/go-aoc/puzzle"
)

type testDay struct {
	input []string
}

func (d *testDay) Part1() (string, error) {
	if len(d.input) > 1 {
		return "yey", nil
	}

	return "", puzzle.NoSolutionError{}
}

func (d *testDay) Part2() (string, error) {
	return "", puzzle.NotImplementedError{}
}

func (d *testDay) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, "nope") {
			return puzzle.NoInputError{}
		}

		d.input = append(d.input, text)
	}

	return scanner.Err()
}
