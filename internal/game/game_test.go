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
    fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
    board := FENParser(fen + " w KQkq - 0 1")
    result := BoardToFEN(board)
    if result != fen {
        t.Fatalf("expected %s, got %s", fen, result)
    }
    fmt.Println("BoardToFEN test passed!")
}