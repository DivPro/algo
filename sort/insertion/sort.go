package insertion

import (
	"github.com/divpro/algo/search/binary"
)

func Sort[T any](values []T) []T {
	if len(values) < 2 {
		return values
	}

	result := make([]T, 0, len(values))
	for i := 0; i < len(values); i++ {
		v := values[i]
		pos := binary.SearchPos(v, result)
		result = append(result[:pos+1], result[pos:]...)
		result[pos] = v
	}

	return result
}
