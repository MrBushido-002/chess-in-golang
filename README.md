

# Chess in Golang

A multiplayer chess server and terminal client written in Go. This is my capstone project for boot.devs Back-end developer course. It is meant to demonstrate my skills and what I have learned.

## Features

- User registration and authentication
- Real-time multiplayer chess
- Full move validation including:
  - Castling
  - En passant
  - Pawn promotion
  - Check and checkmate detection
- Game replay
- Terminal UI built with Bubbletea

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Terminal UI:** Bubbletea + Bubbles
- **Database queries:** sqlc
- **Deployment:** Railway + Docker

## How to Play

1. Run the client:
go run cmd/client/*.go

    Register or login
    Create or join a game
    Enter moves in algebraic notation: e2e4
    Use arrow keys to navigate menus
    Press Tab to switch between input fields
    Press Esc to go back
    Press Ctrl+C to quit
