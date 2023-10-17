package tools

import "errors"

var (
	// ErrIndexOutOfRange is returned when the index is out of range
	ErrIndexOutOfRange = errors.New("index out of range")
)

// ElementAtIndex returns the element of the slice at the given index
func ElementAtIndex(slice []int, index int) (result int, err error) {
	if index < 0 || index >= len(slice) {
		err = ErrIndexOutOfRange
		return
	}
	result = slice[index]
	return
}