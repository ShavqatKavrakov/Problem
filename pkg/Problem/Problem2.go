package Problem

func Flatten(irreguralMatrix [][]int) []int {
	value := irreguralMatrix[0]
	slice := make([]int, 0, len(irreguralMatrix)*len(value))
	for _, value := range irreguralMatrix {
		slice = append(slice, value...)
	}
	return slice
}
