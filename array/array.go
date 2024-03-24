package array

// Reverses an array inline
func ReverseArray[T any](toReverse []T) {
    for i, j := 0, len(toReverse) - 1; i <= j; i, j = i+1, j-1 {
        toReverse[i], toReverse[j] = toReverse[j], toReverse[i]
    }
}

func GetCopyOfArray[T any](initial []T) []T {
    arrayCopy :=  make([]T, len(initial))
    copy(arrayCopy, initial[:])
    return arrayCopy
}
