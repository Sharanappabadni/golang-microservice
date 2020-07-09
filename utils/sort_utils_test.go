package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//initialization
	el := []int{9, 8, 7, 6, 4}

	//execution
	el = BubbleSort(el)

	//validation
	assert.NotNil(t, el)

	assert.EqualValues(t, []int{4, 6, 7, 8, 9}, el)
	assert.EqualValues(t, 5, len(el))
}

func TestBubbleSortBestCase(t *testing.T) {
	el := []int{2, 3, 4, 5, 6}

	el = BubbleSort(el)

	//validation
	assert.NotNil(t, el)

	assert.EqualValues(t, []int{2, 3, 4, 5, 6}, el)
	assert.EqualValues(t, 5, len(el))
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort100(b *testing.B) {
	els := getElements(100)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkBubbleSort100(b *testing.B) {
	els := getElements(100)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func Sort(els []int) {
	if len(els) > 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints((els))
}
