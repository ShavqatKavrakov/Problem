package Problem

//RemoveRepeatedElementsSlice удаляет повторяюший элемент в слайс
func RemoveRepeatedElementsSlice(slice []int) []int {
	newSlice := make([]int, 0, len(slice))
	for i := 0; i <= len(slice); i++ {
		i = 0
		newSlice = append(newSlice, slice[i])
		slice = RemoveElemetSlice(slice, slice[i])
	}
	return newSlice
}

////RemoveRepeatedElementsSlice удаляет определённый элемент в слайс
func RemoveElemetSlice(slice []int, elem int) []int {
	newSlice := make([]int, 0, len(slice))
	for _, value := range slice {
		if value == elem {
			continue
		}
		newSlice = append(newSlice, value)
	}
	return newSlice
}
