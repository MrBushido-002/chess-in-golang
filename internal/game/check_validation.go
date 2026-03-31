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
	
	newBoard.CastlingRights = board.CastlingRights
    
    piece := newBoard.Squares[move.Start.Rank][move.Start.File]
    
    if piece != nil {
        if piece.Type == King {
            if piece.Color == White {
                newBoard.CastlingRights.WhiteKingSide = false
                newBoard.CastlingRights.WhiteQueenSide = false
            } else {
                newBoard.CastlingRights.BlackKingSide = false
                newBoard.CastlingRights.BlackQueenSide = false
            }
        }
        if piece.Type == Rook {
            if move.Start == (Square{Rank: 7, File: 7}) {
                newBoard.CastlingRights.WhiteKingSide = false
            }
            if move.Start == (Square{Rank: 7, File: 0}) {
                newBoard.CastlingRights.WhiteQueenSide = false
            }
            if move.Start == (Square{Rank: 0, File: 7}) {
                newBoard.CastlingRights.BlackKingSide = false
            }
            if move.Start == (Square{Rank: 0, File: 0}) {
                newBoard.CastlingRights.BlackQueenSide = false
            }
        }
    }

    newBoard.Squares[move.End.Rank][move.End.File] = newBoard.Squares[move.Start.Rank][move.Start.File]
    newBoard.Squares[move.Start.Rank][move.Start.File] = nil

	if piece != nil && piece.Type == King {
        if move.Start == (Square{Rank: 7, File: 4}) && move.End == (Square{Rank: 7, File: 6}) {
            newBoard.Squares[7][5] = newBoard.Squares[7][7]
            newBoard.Squares[7][7] = nil
        }
        if move.Start == (Square{Rank: 7, File: 4}) && move.End == (Square{Rank: 7, File: 2}) {
            newBoard.Squares[7][3] = newBoard.Squares[7][0]
            newBoard.Squares[7][0] = nil
        }
        if move.Start == (Square{Rank: 0, File: 4}) && move.End == (Square{Rank: 0, File: 6}) {
            newBoard.Squares[0][5] = newBoard.Squares[0][7]
            newBoard.Squares[0][7] = nil
        }
        if move.Start == (Square{Rank: 0, File: 4}) && move.End == (Square{Rank: 0, File: 2}) {
            newBoard.Squares[0][3] = newBoard.Squares[0][0]
            newBoard.Squares[0][0] = nil
        }
    }

    if piece != nil && piece.Type == Pawn {
        if piece.Color == White && move.Start.Rank == 6 && move.End.Rank == 4 {
            ep := Square{Rank: 5, File: move.Start.File}
            newBoard.EnPassant = &ep
        } else if piece.Color == Black && move.Start.Rank == 1 && move.End.Rank == 3 {
            ep := Square{Rank: 2, File: move.Start.File}
            newBoard.EnPassant = &ep 
        } else {
            newBoard.EnPassant = nil
        }
    }
    if piece != nil && piece.Type == Pawn && board.EnPassant != nil && move.End == *board.EnPassant {
        if piece.Color == White {
            newBoard.Squares[move.End.Rank+1][move.End.File] = nil
        } else {
            newBoard.Squares[move.End.Rank-1][move.End.File] = nil
        }
    }
    if piece != nil && piece.Type == Pawn {
        if piece.Color == White && move.End.Rank == 0 {
            newBoard.Squares[move.End.Rank][move.End.File] = &Piece{Type: Queen, Color: White}
        }
        if piece.Color == Black && move.End.Rank == 7 {
            newBoard.Squares[move.End.Rank][move.End.File] = &Piece{Type: Queen, Color: Black}
        }
    }
	
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