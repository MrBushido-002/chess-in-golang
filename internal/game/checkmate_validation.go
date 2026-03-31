package game



func isInCheck(board Board, color Color) bool {
	kingLocation := FindKing(board, color)

    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            piece := board.Squares[i][j]
            if piece != nil && piece.Color != color {
                move := Move{Start: Square{Rank: i, File: j}, End: kingLocation}
                if isValidPieceMove(board, move, piece.Color) {
                    return true
                }
            }
        }
    }
    return false
}

func IsCheckMate(board Board, color Color) bool {
	if isInCheck(board, color) == false {
		return false
	}
	for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
			piece := board.Squares[i][j]
            if piece != nil && piece.Color == color {
				for k := 0; k < 8; k++ {
        			for l := 0; l < 8; l++ {
						departure := Square{Rank: i, File: j}
						destination := Square{Rank: k, File: l}
						move := Move{Start: departure, End: destination }
						if IsValidMove(board, move, color) == true {
							return false
						}
					}
				}
			}
		}
	}
	return true
}