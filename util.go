package nut

type MetaType interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 |
		~int | ~uint | ~int64 | ~uint64 | ~float32 | ~float64 | ~string
}

func ArrayContains[T MetaType](array []T, target T) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return true
		}
	}

	return false
}
