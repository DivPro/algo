package insertion

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSortInts(t *testing.T) {
	for i, v := range []struct {
		actual   []int
		expected []int
	}{
		{
			actual:   []int{2, 1},
			expected: []int{1, 2},
		},
		{
			actual:   []int{},
			expected: []int{},
		},
		{
			actual:   []int{1},
			expected: []int{1},
		},
		{
			actual:   []int{1, 2},
			expected: []int{1, 2},
		},
		{
			actual:   []int{2, 2},
			expected: []int{2, 2},
		},
		{
			actual:   []int{3, 8, 4},
			expected: []int{3, 4, 8},
		},
		{
			actual:   []int{3, 8, 4, 0, 2, 5, 7, 7, 1, 0, 34},
			expected: []int{0, 0, 1, 2, 3, 4, 5, 7, 7, 8, 34},
		},
	} {
		assert.EqualValues(t, v.expected, Sort(v.actual))
		t.Logf("#%d passed", i)
	}
}

func TestSortStrings(t *testing.T) {
	for i, v := range []struct {
		actual   []string
		expected []string
	}{
		{
			actual:   []string{"4", "0", "3", "01", "02", "30", "31"},
			expected: []string{"0", "01", "02", "3", "30", "31", "4"},
		},
	} {
		assert.EqualValues(t, v.expected, Sort(v.actual))
		t.Logf("#%d passed", i)
	}
}

func BenchmarkSort(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(2)
	b.StopTimer()
	data := make([]struct {
		actual []int
		result []int
	}, b.N)
	for i := 0; i < b.N; i++ {
		count := rand.Intn(3) + 3
		data[i].actual = make([]int, 0, count)
		for j := 0; j < cap(data[i].actual); j++ {
			data[i].actual = append(data[i].actual, rand.Intn(100))
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[i].result = Sort(data[i].actual)
	}
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		result, actual := data[n].result, data[n].actual
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[j] > result[j] {
					b.Error(actual)
				}
			}
		}
	}
}
