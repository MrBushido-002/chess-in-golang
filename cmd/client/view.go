package main

import(
	"fmt"
)

func (m Model) View() string {
	switch m.state {
	case StateAuth:
		return m.list.View()
		

	case StateLogin:
		return fmt.Sprintf(
			"login\n\n%s\n%s\n\nPress Tab to switch fields, enter to submit",
			m.inputs[0].View(),
			m.inputs[1].View(),
		)
	case StateRegister:
		return fmt.Sprintf(
			"Register\n\n%s\n%s\n\nPress Tab to switch fields, Enter to submit",
            m.inputs[0].View(),
            m.inputs[1].View(),
		)
	case StateMainMenu:
		return m.list.View()
	
	case StateGameLobby:
		return fmt.Sprintf("Game created!\nGame ID: %s\n\nShare this ID with your opponent.\nWaiting for opponent to join...\n\nPress Esc to cancel", m.gameID)
	
	case StateGame:
		return fmt.Sprintf(
			"Game ID: %s\nTurn: %s\n\n%s\n\nEnter move (e.g. e2e4): %s\nPress Esc to go back",
			m.gameID,
			m.gameState.Turn,
			renderBoard(m.gameState.BoardState),
			m.moveInput.View(),
		)
	
	case StateJoinGame:
    	return fmt.Sprintf("Join Game\n\n%s\n\n%s\n\nPress Enter to join", m.gameIDInput.View(), m.err)

	case StateGameOver:
    	return fmt.Sprintf("Game Over!\n\nFinal Board:\n\n%s\n\nPress Esc to return to main menu", renderBoard(m.gameState.BoardState))

	
	case StateFindReplay:
    	return fmt.Sprintf("Enter Game ID to replay:\n\n%s\n\n%s\n\nPress Enter to load", m.gameIDInput.View(), m.err)

	case StateReplayGame:
		return fmt.Sprintf(
			"Replay - Move %d/%d\n\n%s\n\n← → to navigate, Esc to exit",
			m.replayIndex+1,
			len(m.replayFENs),
			renderBoard(m.replayFENs[m.replayIndex]),
		)
	}
	return ""
}