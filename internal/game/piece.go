package game

import(
	
)

type Color string
type PieceType string

const (
    White Color = "white"
    Black Color = "black"
)

const (
    King   PieceType = "K"
    Queen  PieceType = "Q"
    Rook   PieceType = "R"
    Bishop PieceType = "B"
    Knight PieceType = "N"
    Pawn   PieceType = "P"
)

type Piece struct {
	Type PieceType
	Color Color
}

func charToPiece(char rune) *Piece {
    switch char {
    case 'K':
        return &Piece{Type: King, Color: White}
    case 'k':
        return &Piece{Type: King, Color: Black}
    case 'Q':
        return &Piece{Type: Queen, Color: White}
    case 'q':
        return &Piece{Type: Queen, Color: Black}
    case 'R':
        return &Piece{Type: Rook, Color: White}
    case 'r':
        return &Piece{Type: Rook, Color: Black}
    case 'B':
        return &Piece{Type: Bishop, Color: White}
    case 'b':
        return &Piece{Type: Bishop, Color: Black}
    case 'N':
        return &Piece{Type: Knight, Color: White}
    case 'n':
        return &Piece{Type: Knight, Color: Black}
    case 'P':
        return &Piece{Type: Pawn, Color: White}
    case 'p':
        return &Piece{Type: Pawn, Color: Black}
	}
    return nil
}

func pieceToChar(piece *Piece) rune {
    switch piece.Type {
    case King:
        if piece.Color == White {
            return 'K'
        }
        return 'k'
    case Queen:
        if piece.Color == White {
            return 'Q'
        }
        return 'q'
	case Rook:
        if piece.Color == White {
            return 'R'
        }
        return 'r'
	case Bishop:
        if piece.Color == White {
            return 'B'
        }
        return 'b'
	case Knight:
        if piece.Color == White {
            return 'N'
        }
        return 'n'
	case Pawn:
        if piece.Color == White {
            return 'P'
        }
        return 'p'
    }
    return ' '
}