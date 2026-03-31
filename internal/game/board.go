package game

import(
	
	"strings"
	"fmt"
)

type Board struct{
	Squares [8][8]*Piece
	CastlingRights CastlingRights
	LastMove *Move
	EnPassant *Square
}

type CastlingRights struct {
	WhiteKingSide bool
	WhiteQueenSide bool
	BlackKingSide bool
	BlackQueenSide bool
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
	if len(parts) > 2 {
		for _, char := range parts[2] {
			if char == '-' {
				board.CastlingRights.WhiteKingSide = false
				board.CastlingRights.WhiteQueenSide = false
				board.CastlingRights.BlackKingSide = false
				board.CastlingRights.BlackQueenSide = false
				break
			}
			if char == 'K' {
				board.CastlingRights.WhiteKingSide = true
			}
			if char == 'Q' {
				board.CastlingRights.WhiteQueenSide = true
			}
			if char == 'k' {
				board.CastlingRights.BlackKingSide = true
			}
			if char == 'q' {
				board.CastlingRights.BlackQueenSide = true
			}
		}
	}
	if len(parts) > 3 && parts[3] != "-" {
		file := int(parts[3][0] - 'a')
		rank := 8 - int(parts[3][1] - '0')
		ep := Square{Rank: rank, File: file}
		board.EnPassant = &ep
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

	castling := ""
	if board.CastlingRights.WhiteKingSide {
		castling += "K"
	}
	if board.CastlingRights.WhiteQueenSide {
		castling += "Q"
	}
	if board.CastlingRights.BlackKingSide {
		castling += "k"
	}
	if board.CastlingRights.BlackQueenSide {
		castling += "q"
	}
	if castling == "" {
		castling = "-"
	}
	
	enPassant := "-"
	if board.EnPassant != nil {
		enPassant = fmt.Sprintf("%c%d", 'a'+board.EnPassant.File, 8 - board.EnPassant.Rank)
	}
	fenString += " w " + castling + " " + enPassant + " 0 1"

	return fenString
}