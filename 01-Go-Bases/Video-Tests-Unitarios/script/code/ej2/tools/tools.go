package tools

import "fmt"

func ElementAtIndex(slice []int, index int) (int, error) {
	if index < 0 || index >= len(slice) {
		return 0, fmt.Errorf("index out of boundaries")
	}
	return slice[index], nil
}
