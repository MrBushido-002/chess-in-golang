package game

import (
    "testing"
    "fmt"
)

func TestFENParser(t *testing.T) {
    fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
    board := FENParser(fen)

    // Check that e1 (rank 7, file 4) is a white king
    piece := board.Squares[7][4]
    if piece == nil {
        t.Fatal("expected a piece at e1, got nil")
    }
    if piece.Type != King || piece.Color != White {
        t.Fatalf("expected white king at e1, got %v %v", piece.Color, piece.Type)
    }
    fmt.Println("FEN parser test passed!")
}

func TestBoardToFEN(t *testing.T) {
    fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
    board := FENParser(fen)
    result := BoardToFEN(board)
    if result != fen {
        t.Fatalf("expected %s, got %s", fen, result)
    }
    fmt.Println("BoardToFEN test passed!")
}

func TestCheckValidation(t *testing.T) {
    // set up a board where white king is in check
    board := FENParser("8/8/8/8/4r3/8/8/4K3 w - - 0 1")
    
    // try a move that keeps king in check
    move := Move{
        Start: Square{Rank: 7, File: 4},
        End:   Square{Rank: 6, File: 4},
    }
    if IsValidMove(board, move, White) {
        t.Fatal("expected invalid move - king would remain in check")
    }
}

func TestCheckMateValidation(t *testing.T) {
    // set up a board where white king is in check
    board := FENParser("6rr/8/8/8/8/8/8/7K w - - 0 1")
    if IsCheckMate(board, White)  == false {
        t.Fatal("expected invalid move - should be checkmate")
    }
}

func TestCastling(t *testing.T) {
    // board with clear path for white kingside castling
    // K on e1, R on h1, empty f1 and g1
    board := FENParser("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQK2R w KQkq - 0 1")
    
    move := Move{
        Start: Square{Rank: 7, File: 4},
        End:   Square{Rank: 7, File: 6},
    }
    if !IsValidMove(board, move, White) {
        t.Fatal("expected kingside castling to be valid")
    }
}

func TestEnPassant(t *testing.T) {
    board := FENParser("rnbqkbnr/ppp1pppp/8/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 1")
    
    move := Move{
        Start: Square{Rank: 3, File: 4},  // white pawn at e5
        End:   Square{Rank: 2, File: 3},  // capture to d6
    }
    if !IsValidMove(board, move, White) {
        t.Fatal("expected en passant to be valid")
    }
}

func TestPawnPromotion(t *testing.T) {
    board := FENParser("8/4P3/8/8/8/8/8/4K2k w - - 0 1")
    move := Move{
        Start: Square{Rank: 1, File: 4},
        End: Square{Rank: 0, File: 4},
    }
    newBoard := HypotheticalMove(board, move)
    piece := newBoard.Squares[0][4]
    if piece == nil || piece.Type != Queen {
        t.Fatal("expected piece to be Queen")
    }
}