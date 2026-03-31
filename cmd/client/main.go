package main

import(
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Println("error running program", err)
		os.Exit(1)
	}
}