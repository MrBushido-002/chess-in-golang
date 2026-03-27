package game

import(
	"math"
)

type Move struct {
	Start Square
	End Square
}

func isValidPieceMove(board Board, move Move, color Color) bool {
	piece := board.Squares[move.Start.Rank][move.Start.File]
	if piece == nil {
		return false
	}

	if piece.Color != color {
		return false
	}

	switch piece.Type {
	case King:
		return isValidKingMove(board, move, color)
	case Queen:
		return isValidQueenMove(board, move, color)
	case Rook:
		return isValidRookMove(board, move, color)
	case Bishop:
		return isValidBishopMove(board, move, color)
	case Knight:
		return isValidKnightMove(board, move, color)
	case Pawn:
		return isValidPawnMove(board, move, color)
	}
	return false
}


func isValidKingMove(board Board, move Move, color Color) bool {
	if move.Start.Rank == move.End.Rank && move.Start.File == move.End.File {
		return false
	}
	if board.Squares[move.End.Rank][move.End.File] != nil && board.Squares[move.End.Rank][move.End.File].Color == color {
		return false
	}
	return math.Abs(float64(move.Start.Rank - move.End.Rank)) <= 1 && math.Abs(float64(move.Start.File - move.End.File)) <= 1
}

func isValidQueenMove(board Board, move Move, color Color) bool {
	if move.Start.Rank == move.End.Rank || move.Start.File == move.End.File {
        return isValidRookMove(board, move, color)
    }
    return isValidBishopMove(board, move, color)
}

func isValidRookMove(board Board, move Move, color Color) bool {
	if move.Start.Rank != move.End.Rank && move.Start.File != move.End.File {
		return false
	}
	//right
	if move.Start.Rank == move.End.Rank && move.Start.File < move.End.File {
		for i := move.Start.File + 1; i < move.End.File; i++ {
			if board.Squares[move.Start.Rank][i] != nil {
				return false
			}
		}
	}
	//left
	if move.Start.Rank == move.End.Rank && move.Start.File > move.End.File {
		for i := move.Start.File - 1; i > move.End.File; i-- {
			if board.Squares[move.Start.Rank][i] != nil {
				return false
			}
		}
	}
	
	//up
	if move.Start.File == move.End.File && move.Start.Rank < move.End.Rank {
		for i := move.Start.Rank + 1; i < move.End.Rank; i++ {
			if board.Squares[i][move.Start.File] != nil {
				return false
			}
		}
	}
	//down
	if move.Start.File == move.End.File && move.Start.Rank > move.End.Rank{
		for i := move.Start.Rank - 1; i > move.End.Rank; i-- {
			if board.Squares[i][move.Start.File] != nil {
				return false
			}
		}
	}
	if board.Squares[move.End.Rank][move.End.File] != nil && board.Squares[move.End.Rank][move.End.File].Color == color {
		return false
	}
	return true
}


func isValidBishopMove(board Board, move Move, color Color) bool {
	if board.Squares[move.End.Rank][move.End.File] != nil && board.Squares[move.End.Rank][move.End.File].Color == color {
		return false
	}
	if math.Abs(float64(move.Start.Rank - move.End.Rank)) != math.Abs(float64(move.Start.File - move.End.File)) {
		return false
	}
	
	diff := int(math.Abs(float64(move.Start.Rank - move.End.Rank)))
	//Down-Right
	if move.Start.Rank < move.End.Rank && move.Start.File < move.End.File {
		for i := 1; i < diff; i++ {
			if board.Squares[move.Start.Rank + i][move.Start.File + i] != nil {
				return false
			}
		}
		return true
	}

	//Down-Left
	if move.Start.Rank < move.End.Rank && move.Start.File > move.End.File {
		for i := 1; i < diff; i++ {
			if board.Squares[move.Start.Rank + i][move.Start.File - i] != nil {
				return false
			}
		}
		return true
	}

	//Up-Left
		//Down-Left
	if move.Start.Rank > move.End.Rank && move.Start.File > move.End.File {
		for i := 1; i < diff; i++ {
			if board.Squares[move.Start.Rank - i][move.Start.File - i] != nil {
				return false
			}
		}
		return true
	}

	//Up-Right

		//Down-Left
	if move.Start.Rank > move.End.Rank && move.Start.File < move.End.File {
		for i := 1; i < diff; i++ {
			if board.Squares[move.Start.Rank - i][move.Start.File + i] != nil {
				return false
			}
		}
		return true
	}
	return false
}

func isValidKnightMove(board Board, move Move, color Color) bool {
	if board.Squares[move.End.Rank][move.End.File] != nil && board.Squares[move.End.Rank][move.End.File].Color == color {
		return false
	}
	rank_diff := int(math.Abs(float64(move.Start.Rank - move.End.Rank)))
	file_diff := int(math.Abs(float64(move.Start.File - move.End.File)))

	return (rank_diff == 1 && file_diff == 2) || (rank_diff == 2 && file_diff == 1) 
}

func isValidPawnMove(board Board, move Move, color Color) bool {
	if board.Squares[move.End.Rank][move.End.File] != nil && board.Squares[move.End.Rank][move.End.File].Color == color {
		return false
	}
	if color == White {
		forwardOne := move.Start.Rank - 1
		forwardTwo := move.Start.Rank - 2

		if move.End.Rank == forwardOne && move.End.File == move.Start.File {
			return board.Squares[forwardOne][move.Start.File] == nil
		}

		if move.Start.Rank == 6 && move.End.Rank == forwardTwo && move.End.File == move.Start.File {
			return board.Squares[forwardOne][move.Start.File] == nil &&
				board.Squares[forwardTwo][move.Start.File] == nil
		}

		if board.Squares[move.End.Rank][move.End.File] != nil {
			return move.End.Rank == forwardOne && (move.End.File == move.Start.File + 1 || move.End.File == move.Start.File - 1)
		}
	}
	return false

	if color == Black {
		forwardOne := move.Start.Rank + 1
		forwardTwo := move.Start.Rank + 2

		if move.End.Rank == forwardOne && move.End.File == move.Start.File {
			return board.Squares[forwardOne][move.Start.File] == nil
		}

		if move.Start.Rank == 1 && move.End.Rank == forwardTwo && move.End.File == move.Start.File {
			return board.Squares[forwardOne][move.Start.File] == nil &&
				board.Squares[forwardTwo][move.Start.File] == nil
		}

		if board.Squares[move.End.Rank][move.End.File] != nil {
			return move.End.Rank == forwardOne && (move.End.File == move.Start.File + 1 || move.End.File == move.Start.File - 1)
		}

	}
	return false
}

func IsValidMove(board Board, move Move, color Color) bool {
    if !isValidPieceMove(board, move, color) {
        return false
    }
    if CheckValidation(board, move, color) {
        return false
    }
    return true
}
