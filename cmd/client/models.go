package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"

    
)    

type State int

const (
	StateAuth State = iota
	StateLogin
	StateRegister
	StateMainMenu
	StateCreateGame
	StateJoinGame
	StateGameLobby
	StateGame
    StateGameOver

)

type Model struct {
	state State
	token string
	gameID string
	username string
	list list.Model
	inputs []textinput.Model
	focused int
	err string
	gameState GameState
	gameIDInput textinput.Model
	moveInput textinput.Model
}

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

func initialInputs() []textinput.Model {
    username := textinput.New()
    username.Placeholder = "Username"
    username.Focus()

    password := textinput.New()
    password.Placeholder = "Password"
    password.EchoMode = textinput.EchoPassword

    return []textinput.Model{username, password}
}

func authMenuList() list.Model {
    items := []list.Item{
        item("Login"),
        item("Register"),
    }
    l := list.New(items, list.NewDefaultDelegate(), 30, 30)
    l.Title = "Chess in Golang"
    l.SetShowStatusBar(false)
    l.SetFilteringEnabled(false)
    l.SetShowHelp(false)
    return l
}

func mainMenuList() list.Model {
    items := []list.Item{
        item("Create Game"),
        item("Join Game"),
        item("View Games"),
    }
    l := list.New(items, list.NewDefaultDelegate(), 30, 30)
    l.Title = "Main Menu"
    l.SetShowStatusBar(false)
    l.SetFilteringEnabled(false)
    l.SetShowHelp(false)
    return l
}

func InitialModel() Model {
    gameIDInput := textinput.New()
	gameIDInput.Placeholder = "Enter Game ID"
	moveInput := textinput.New()
	moveInput.Placeholder = "Enter move (e.g. e2 e4)"

	return Model{
        state:  StateAuth,
        list:   authMenuList(),
        inputs: initialInputs(),
		gameIDInput: gameIDInput,
		moveInput: moveInput,
    }
}



