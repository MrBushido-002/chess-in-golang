package main

import (
	"fmt"
	"strings"
)

func renderBoard(boardState string) string {
	var pieceSymbols = map[rune]string{
		'K': "♔", 'Q': "♕", 'R': "♖", 'B': "♗", 'N': "♘", 'P': "♙",
		'k': "♚", 'q': "♛", 'r': "♜", 'b': "♝", 'n': "♞", 'p': "♟",
	}
	parts := strings.Split(boardState, " ")
    ranks := strings.Split(parts[0], "/")
    
    result := "  a b c d e f g h\n"
    
    for i, rank := range ranks {
        result += fmt.Sprintf("%d ", 8-i)
        file := 0
        for _, char := range rank {
            if char >= '1' && char <= '8' {
                for j := 0; j < int(char-'0'); j++ {
                    result += ". "
                }
                file += int(char - '0')
            } else {
                // add piece character
                if symbol, ok := pieceSymbols[char]; ok {
					result += symbol + " "
				} else {
					result += string(char) + " "
				}

                file++
            }
        }
        result += fmt.Sprintf("%d\n", 8-i)
    }
    result += "  a b c d e f g h\n"
    return result
}
