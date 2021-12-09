package math

// SumInt64 returns the sum of all the provided values.
func SumInt64(values ...int64) int64 {
	var result int64

	for _, value := range values {
		result += value
	}

	return result
}
