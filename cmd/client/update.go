package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func fetchGameCmd(token, gameID string) tea.Cmd {
	return func() tea.Msg {
		gameState, err := GetGame(token, gameID)
		if err != nil {
			return err
		}
		return gameState
	}
}

func pollGameCmd(token, gameID string) tea.Cmd {
    return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		gameState, err := GetGame(token, gameID)
        if err != nil {
            return err
        }
        return gameState
    })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case StateAuth:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "tab":
				m.inputs[m.focused].Blur()
				m.focused = (m.focused + 1) % len(m.inputs)
				m.inputs[m.focused].Focus()
				return m, nil
			case "enter":
				selected := m.list.SelectedItem().(item)
				if selected == "Login" {
					m.state = StateLogin
				} else {
					m.state = StateRegister
				}
				return m, nil
			}
		case StateLogin, StateRegister:
			switch msg.String() {
			case "tab":
				m.inputs[m.focused].Blur()
				m.focused = (m.focused + 1) % len(m.inputs)
				m.inputs[m.focused].Focus()
				return m, nil
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = StateAuth
				m.list = authMenuList()
				m.focused = 0
				m.inputs = initialInputs()
				return m, nil
			case "enter":
				if m.focused < len(m.inputs)-1 {
					m.inputs[m.focused].Blur()
					m.focused++
					m.inputs[m.focused].Focus()
					return m, nil
				}
				username := m.inputs[0].Value()
				password := m.inputs[1].Value()
				if m.state == StateLogin {
					token, err := LoginUser(username, password)
					if err != nil {
						return m, nil
					}
					m.token = token
					m.state = StateMainMenu
					m.list = mainMenuList()
					m.username = username
				}
				if m.state == StateRegister {
					err := RegisterUser(username, password)
					if err != nil {
						return m, nil
					}
					m.state = StateAuth
				}
			}
		case StateMainMenu:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				selected := m.list.SelectedItem().(item)
				switch selected {
				case "Create Game":
					gameID, err := CreateGame(m.token)
					if err != nil {
						m.err = err.Error()
						return m, nil
					}
					m.gameID = gameID
					m.state = StateGameLobby
					return m, pollGameCmd(m.token, gameID)
				case "Join Game":
					m.gameIDInput.Focus()
					m.state = StateJoinGame
				case "View Game":
					m.state = StateGame
				}
				return m, nil
			}
		case StateGameLobby:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = StateMainMenu
				return m, nil
			}
		case StateJoinGame:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = StateMainMenu
				return m, nil
			case "enter":
				gameID, err := JoinGame(m.token, m.gameIDInput.Value())
				if err != nil {
					m.err = err.Error()
					return m, nil
				}
				m.gameID = gameID
				m.state = StateGame
				m.moveInput.Focus()
				return m, fetchGameCmd(m.token, gameID)
			}
		case StateGame:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = StateMainMenu
				return m, nil
			case "enter":
				moveStr := m.moveInput.Value()
				if len(moveStr) < 4 {
					m.err = "Invalid move format - use e2e4"
					return m, nil
				}
				err := MakeMove(m.token, m.gameID, moveStr)
				if err != nil {
					m.err = err.Error()
					return m, nil
				}
				m.moveInput.SetValue("")
				return m, fetchGameCmd(m.token, m.gameID)
			}
		case StateGameOver:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = StateMainMenu
				return m, nil
			}
			
		}
	case GameState:
		m.gameState = msg
		if m.state == StateGameLobby && msg.Status == "active" {
			m.state = StateGame
			m.moveInput.Focus()
		}
		if m.state == StateGame && msg.Status == "complete" {
        	m.state = StateGameOver
		}
		if m.state == StateGameLobby {
			return m, pollGameCmd(m.token, m.gameID)
		}
		if m.state == StateGame {
			return m, pollGameCmd(m.token, m.gameID)
		}
		return m, nil
	}

	var cmd tea.Cmd
	switch m.state {
	case StateAuth, StateMainMenu:
		m.list, cmd = m.list.Update(msg)
	case StateLogin, StateRegister:
		m.inputs[m.focused], cmd = m.inputs[m.focused].Update(msg)
	case StateJoinGame:
		m.gameIDInput, cmd = m.gameIDInput.Update(msg)
	case StateGame:
		m.moveInput, cmd = m.moveInput.Update(msg)
	}
	return m, cmd
}