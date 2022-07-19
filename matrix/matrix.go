package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.

type Matrix struct {
	matrix [][]int
}

func New(s string) (*Matrix, error) {
	var m [][]int
	for e, s := range strings.Split(s, "\n") {
		m = append(m, []int{})
		strArray := strings.Split(strings.TrimSpace(s), " ")
		for i := range strArray {
			n, err := strconv.Atoi(strArray[i])
			if err != nil {
				return nil, err
			}
			m[e] = append(m[e], n)
		}
		if e > 0 && len(m[0]) != len(m[e]) {
			return nil, errors.New("uneven rows")
		}
	}
	return &Matrix{m}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() (cols [][]int) {
	for i := range m.matrix[0] {
		var intSlice []int
		for _, r := range m.matrix {
			intSlice = append(intSlice, r[i])
		}
		cols = append(cols, intSlice)
	}
	return
}

func (m *Matrix) Rows() (rows [][]int) {
	for _, r := range m.matrix {
		var intSlice []int
		for _, c := range r {
			intSlice = append(intSlice, c)
		}
		rows = append(rows, intSlice)
	}
	return
}

func (m *Matrix) Set(row, col, val int) (canSet bool) {
	switch {
	case row < 0 || row > len(m.matrix)-1 || col < 0 || col > len(m.matrix[0])-1:
    default:
    	m.matrix[row][col] = val
        canSet = true
	}
	return
}
