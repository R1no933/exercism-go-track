package strain

type Slicer interface {
	int | string | []int
}

func Keep[T Slicer](input []T, filter func(T) bool) []T {
	if input == nil {
		return input
	}
	output := make([]T, 0, len(input))
	for _, value := range input {
		if filter(value) {
			output = append(output, value)
		}
	}
	return output
}
func Discard[T Slicer](input []T, filter func(T) bool) []T {
	return Keep(input, func(val T) bool { return !filter(val) })
}
