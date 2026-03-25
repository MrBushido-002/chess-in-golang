package game

import(
	
	"strings"
	"fmt"
)

type Board struct{
	Squares [8][8]*Piece
}

func FENParser(FENString string) Board {
	var board Board
	parts := strings.Split(FENString, " ")
	ranks := strings.Split(parts[0], "/")
	for i, rank := range ranks {
    file := 0
		for _, char := range rank {
			if char >= '1' && char <= '8' {
				file += int(char - '0')
			} else {
				board.Squares[i][file] = charToPiece(char)
				file++ 
			}
		}
	}
	return board
}

func BoardToFEN(board Board) string {
	fenString := ""
	for i := 0; i < 8; i++ {
		emptyCount := 0
    	for j := 0; j < 8; j++ {
			if board.Squares[i][j] == nil {
				emptyCount++
			} else {
				if emptyCount > 0 {
					fenString += fmt.Sprintf("%d", emptyCount)
					emptyCount = 0
				}
				fenString += string(pieceToChar(board.Squares[i][j]))
			}
		}
		if emptyCount > 0 {
			fenString += fmt.Sprintf("%d", emptyCount)
		}
		if i < 7 {
			fenString += "/"
		}
	}
	return fenString
}