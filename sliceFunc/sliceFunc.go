package sliceFunc

func MapSlice[Input, Output any](originalSlice []Input, f func(value Input, index int) Output) []Output {
	result := make([]Output, len(originalSlice))
	for index, value := range originalSlice {
		result[index] = f(value, index)
	}
	return result
}
