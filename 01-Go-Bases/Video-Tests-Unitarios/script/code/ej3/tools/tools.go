package tools

import "fmt"

var ErrIndexOutOfBoundaries = fmt.Errorf("index out of boundaries")

func ElementAtIndex(slice []int, index int) (int, error) {
	if index < 0 || index >= len(slice) {
		return 0, ErrIndexOutOfBoundaries
	}
	return slice[index], nil
}
