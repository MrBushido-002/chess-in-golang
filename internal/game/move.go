package game


type Move struct {
	Start Square
	End Square
}

func IsValidMove(board Board, move Move, color Color) bool {
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

}

func isValidQueenMove(board Board, move Move, color Color) bool {
	
}

func isValidRookMove(board Board, move Move, color Color) bool {
	
}

func isValidBishopMove(board Board, move Move, color Color) bool {
	
}

func isValidKnightMove(board Board, move Move, color Color) bool {
	
}

func isValidPawnMove(board Board, move Move, color Color) bool {
	if color == "white" {
		if 
	}

	if color == "black" {}
}
