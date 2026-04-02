

# Chess in Golang

A multiplayer chess server and terminal client written in Go.

## Motivation
This is my capstone project for boot.devs Back-end developer course. I wanted to create something fun and interesting while still displaying the wide range of skills I have developed throughout the course including programming in golang, using PostgresSQL, and creating a server and client application that work together.

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

## Quick Start

1. Run the client:
```bash
go run cmd/client/*.go
```
2. Register
3. Login
4. Create Game
5. Send another user the game_id
6. Have them join your game
7. Start playing chess in your terminal!

## Usage
Register or login
    Create or join a game
    Enter moves in algebraic notation: e2e4
    Use arrow keys to navigate menus
    Press Tab to switch between input fields
    Press Esc to go back
    Press Ctrl+C to quit

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/MrBushido-002/chess-in-golang
cd zipzod
```

### Build the compiled binary

```bash
go build
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
