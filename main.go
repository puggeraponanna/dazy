package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func main() {
	p := tea.NewProgram(&model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	width = width - 3
	var border = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	var style = lipgloss.NewStyle().
		Height(height - 2).
		Border(border)
	var style1 = style.Width(width * 2 / 10)
	var style2 = style.Width(width * 8 / 10)
	s := lipgloss.JoinHorizontal(lipgloss.Left,
		style1.Render("SideTab"),
		style2.Render("Body"))
	return s
}
