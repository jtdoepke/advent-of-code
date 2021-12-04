package main

// BingoBoard represents the state of a Bingo board.
type BingoBoard struct {
	numbers [][]int
	called  [][]bool
}

// NewBingoBoard returns a new BingoBoard from a grid of ints.
func NewBingoBoard(numbers [][]int) *BingoBoard {
	b := &BingoBoard{
		numbers: numbers,
		called:  make([][]bool, len(numbers)),
	}
	for i := range numbers {
		b.called[i] = make([]bool, len(numbers))
	}
	return b
}

// Call marks num as called on this board, if present.
func (b *BingoBoard) Call(num int) {
	for i, row := range b.numbers {
		for j, cel := range row {
			if cel == num {
				b.called[i][j] = true
				return
			}
		}
	}
}

// HasWon returns true if the board is a winner.
func (b *BingoBoard) HasWon() bool {
	columnCounts := make([]int, len(b.numbers))
	for _, row := range b.called {
		allTrue := true
		for i, called := range row {
			if !called {
				allTrue = false
			} else {
				columnCounts[i]++
			}
		}
		if allTrue {
			return true
		}
	}
	for _, c := range columnCounts {
		if c == len(b.called) {
			return true
		}
	}
	return false
}

// Score returns the score of this board given the last call before the board won.
func (b *BingoBoard) Score(winningCall int) (score int) {
	for i, row := range b.numbers {
		for j, cel := range row {
			if !b.called[i][j] {
				score += cel
			}
		}
	}
	score *= winningCall
	return
}

// Clear unmarks all board numbers as called.
func (b *BingoBoard) Clear() {
	for _, row := range b.called {
		for j := range row {
			row[j] = false
		}
	}
}
