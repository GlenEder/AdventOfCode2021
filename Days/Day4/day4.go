package Day4

import (
	"fmt"
	"src/Utils"
	"strings"
)

type boardMarker struct {
	number int
	called bool
}

type bingo struct {
	 board [][]boardMarker
	 won bool
}

func (b bingo) String() string  {
	toReturn := "===Board===\n"
	for _, board := range b.board {
		for _, marker := range board {
			if marker.called { toReturn += Utils.GetColorCode(Utils.Green) }
			toReturn += Utils.IntToString(marker.number)
			if marker.called { toReturn += Utils.GetColorCode(Utils.ColorReset) }
			toReturn += "\t"
		}
		toReturn += "\n"
	}
	return toReturn
}

//Checks board for token, if has checks for win
//@param token -- token called by bingo roller
//@return true if board now winning, otherwise false
func (b *bingo)CheckWin(token int) bool {
	//Find token in board
	for x, row := range b.board {
		for y, marker := range row {
			if marker.number == token {
				b.board[x][y].called = true

				//check for vert win
				win := true
				for i := 0; i < len(b.board[x]); i++ {
					if !b.board[x][i].called {
						win = false
						break
					}
				}
				//check for win
				if win {
					b.won = true
					return win
				}
				//reset win flag
				win = true
				for i := 0; i < len(b.board); i++ {
					if !b.board[i][y].called {
						win = false
						break
					}
				}
				if win { b.won = true }
				//return win state
				return win
			}
		}
	}

	// return false if number not on board
	return false
}

//Calcs sum of unmarked numbers on the board
//@return sum of unmarked board numbers
func (b bingo) GetUnmarkedSum() int {
	sum := 0
	for _, row := range b.board {
		for _, marker := range row {
			if !marker.called {
				sum += marker.number
			}
		}
	}

	return sum
}

func Run() {
	Utils.PrintWithColor(Utils.Cyan, "===Day 4===\n")

	//get input
	input := Utils.ReadInputToSlice("Days/Day4/input.txt")
	//Get bingo calls
	bingoCalls := input[0]
	//Create boards
	var bingoBoards []bingo
	var board bingo
	row := 0
	for i := 2; i < len(input); i++ {
		line := input[i]
		if line == "" {
			//Create new board
			bingoBoards = append(bingoBoards, board)
			row = 0
			board = bingo{}
		} else {
			//board line
			var boardLine []boardMarker
			parts := strings.Fields(line)
			for _, part := range parts {
				marker := boardMarker{
					called: false,
					number: Utils.StringToInt(part),
				}
				boardLine = append(boardLine, marker)
			}
			board.won = false
			board.board = append(board.board, boardLine)
			row++
		}
	}
	//Add last board
	bingoBoards = append(bingoBoards, board)
	
	// Play bingo !!
	numWinners := 0
	for _, sNum := range strings.Split(bingoCalls, ",") {
		numCalled := Utils.StringToInt(sNum)
		for i := range bingoBoards {
			if bingoBoards[i].won { continue }
			if bingoBoards[i].CheckWin(numCalled) {
				//fmt.Printf("board won on number: %d\n%s\n", numCalled, bingoBoards[i])
				numWinners++
				if numWinners == 1 {
					//Calc winning data
					fmt.Printf("Part 1: %d\n", bingoBoards[i].GetUnmarkedSum() * numCalled)
				} else if numWinners == len(bingoBoards) {
					fmt.Printf("Part 2: %d\n", bingoBoards[i].GetUnmarkedSum() * numCalled)
				}
			}
		}
	}


}