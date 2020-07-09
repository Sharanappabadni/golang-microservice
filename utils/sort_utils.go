package utils

func BubbleSort(el []int) []int {
	keepRunning := true

	for keepRunning {
		keepRunning = false

		for i := 0; i < len(el)-1; i++ {
			if el[i] > el[i+1] {
				el[i], el[i+1] = el[i+1], el[i]
				keepRunning = true
			}
		}
	}
	return el
}
