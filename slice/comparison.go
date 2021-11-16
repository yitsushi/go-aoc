package slice

// EqualInt64 test if two slices are equal or not.
// They are equal if the yhave the same set of element
// in the same order.
func EqualInt64(slice1, slice2 []int64) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for idx := 0; idx < len(slice1); idx++ {
		if slice1[idx] != slice2[idx] {
			return false
		}
	}

	return true
}
