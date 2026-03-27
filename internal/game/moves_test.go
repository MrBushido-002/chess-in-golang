package game

import "testing"

func TestPawnMove(t *testing.T) {
    board := FENParser("8/8/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
    
    // valid forward one square
    move := Move{
        Start: Square{Rank: 6, File: 4},
        End:   Square{Rank: 5, File: 4},
    }
    if !IsValidMove(board, move, White) {
        t.Fatal("expected valid pawn move forward one square")
    }
}