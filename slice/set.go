package slice

// IntersectString returns with items that can be found in all provided slices.
func IntersectString(lists ...[]string) []string {
	intersect := []string{}

	if len(lists) == 0 {
		return intersect
	}

	if len(lists) == 1 {
		return lists[0]
	}

	for _, value := range lists[0] {
		exists := true

		for _, list := range lists[1:] {
			if _, found := ContainsString(list, value); !found {
				exists = false

				break
			}
		}

		if exists {
			intersect = append(intersect, value)
		}
	}

	return intersect
}
