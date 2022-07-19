package cards

import (
	"math"
)

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) int {
	if index > len(slice)-1 || math.Signbit(float64(index)) {
		return -1
	}
	return slice[index]
}

// PrependItems accepts a slice of integers and prepends it with any number of passed integers
func PrependItems(slice []int, value ...int) []int {
	return append(value, slice...)
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index > len(slice)-1 || math.Signbit(float64(index)) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if !math.Signbit(float64(index)) && index <= len(slice)-1 {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}
