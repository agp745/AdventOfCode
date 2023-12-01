package utils

func Pop[T comparable](arr *[]T) T {
	if len(*arr) == 0 {
		panic("Error: cannot pop empty slice")
	}

	popped := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]

	return popped
}

func Push[T comparable](arr *[]T, v T) {
	*arr = append(*arr, v)
}
