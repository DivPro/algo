package binary

func SearchGE[T any](v T, s []T) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if less(v, s[0]) {
			return 0
		}

		return 1
	}
	leftIdx := len(s)/2 - 1
	if less(v, s[leftIdx]) {
		return SearchGE(v, s[0:leftIdx])
	}

	return SearchGE(v, s[leftIdx+1:]) + leftIdx + 1
}

func less(i, j any) bool {
	switch v := i.(type) {
	case int:
		return v < j.(int)
	case float32:
		return v < j.(float32)
	case float64:
		return v < j.(float64)
	case string:
		return v < j.(string)
	default:
		return false
	}
}
