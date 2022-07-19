package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (occupied int) {
	for _, r := range cb[rank] {
		if r {
			occupied++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (occupied int) {
	if file < 1 || file > 8 {
		return 0
	}
	for _, r := range cb {
		if r[file-1] {
			occupied++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (square int) {
	for _, r := range cb {
		for range r {
			square++
		}
	}
	return square
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (occupied int) {
	for _, r := range cb {
		for _, v := range r {
			if v {
				occupied++
			}
		}
	}
	return
}
