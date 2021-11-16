package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

// Int64Range parses a string as an int64 range.
func Int64Range(text, separator string) (int64, int64, error) {
	parts := strings.SplitN(text, "-", rangeParts)

	from, err := strconv.ParseInt(parts[0], base10, bit64)
	if err != nil {
		return 0, 0, fmt.Errorf("parse error: %w", err)
	}

	to, err := strconv.ParseInt(parts[1], base10, bit64)
	if err != nil {
		return from, 0, fmt.Errorf("parse error: %w", err)
	}

	return from, to, nil
}
