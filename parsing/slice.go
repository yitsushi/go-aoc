package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

// Int64Slice parses a string as a slice of int64 values.
func Int64Slice(text, separator string) ([]int64, error) {
	list := []int64{}

	for _, value := range strings.Split(text, separator) {
		intValue, err := strconv.ParseInt(value, base10, bit64)
		if err != nil {
			return list, fmt.Errorf("parse error: %w", err)
		}

		list = append(list, intValue)
	}

	return list, nil
}
