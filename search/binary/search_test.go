package binary

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSearchPos(t *testing.T) {
	source := []int{4, 8, 7, 45, 52, 135}
	for i, v := range []struct {
		searched int
		expected int
	}{
		{
			searched: 56,
			expected: 5,
		},
		{
			searched: 0,
			expected: 0,
		},
		{
			searched: 201,
			expected: 6,
		},
		{
			searched: 7,
			expected: 3,
		},
	} {
		assert.EqualValues(t, v.expected, SearchPos(v.searched, source), v.searched)
		t.Logf("#%d passed", i)
	}
}

func BenchmarkSearchPos(b *testing.B) {
	source := []int{4, 8, 7, 45, 52, 135}
	for _, v := range []struct {
		name  string
		value int
	}{
		{"first", 0},
		{"middle", 8},
		{"last", 201},
	} {
		b.Run(v.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(2)
			for i := 0; i < b.N; i++ {
				SearchPos(v.value, source)
			}
		})
	}
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(2)
		b.StopTimer()
		actual := make([]int, 1, 1e6)
		var value int
		for i := 0; i < cap(actual); i++ {
			if i == 0 {
				actual[i] = rand.Intn(100)
				continue
			}
			actual = append(actual, actual[i-1]+rand.Intn(100))
			if i == cap(actual)/2 {
				value = actual[i]
			}
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			SearchPos(value, source)
		}
	})
}
