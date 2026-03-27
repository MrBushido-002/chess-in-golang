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