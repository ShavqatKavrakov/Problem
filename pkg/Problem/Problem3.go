package Problem

func Make2D(slice []int, lenMatrix int) [][]int {
	length := len(slice) / lenMatrix
	if len(slice)%lenMatrix != 0 {
		length = (len(slice) / lenMatrix) + 1
	}
	newSlice := make([]int, 0, lenMatrix*length)
	newSlice = append(newSlice, slice...)
	irreguralMatrix := make([][]int, length, length)
	for i := 0; i < length; i++ {
		irreguralMatrix[i] = append(irreguralMatrix[i], newSlice[i*lenMatrix:(i+1)*lenMatrix]...)
	}
	return irreguralMatrix
}
