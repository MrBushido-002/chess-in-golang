package game

import(

)

func FindKing(board Board, color Color) Square {
	for i := 0; i < 8; i++ {
    	for j := 0; j < 8; j++ {
			piece := board.Squares[i][j]
			if piece != nil && piece.Type == King && piece.Color == color {
				return Square{Rank: i, File: j}
			}
		}
	}
	return Square{}
}

func HypotheticalMove(board Board, move Move) Board {
    var newBoard Board
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            newBoard.Squares[i][j] = board.Squares[i][j]
        }
    }
    newBoard.Squares[move.End.Rank][move.End.File] = newBoard.Squares[move.Start.Rank][move.Start.File]
    newBoard.Squares[move.Start.Rank][move.Start.File] = nil
    return newBoard
}

func CheckValidation(board Board, move Move, color Color) bool {
	//true == puts king in check
	//false == king is safe after move 
	newBoard := HypotheticalMove(board, move)
	kingLocation := FindKing(newBoard, color)

	for i := 0; i < 8; i++ {
    	for j := 0; j < 8; j++ {
			piece := newBoard.Squares[i][j]
			if piece != nil && piece.Color != color {
				move := Move{Start: Square{Rank: i, File: j}, End: kingLocation}
				if isValidPieceMove(newBoard, move, piece.Color) {
					return true
				}
			}
		}
	}
	return false		
}