package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/examples"
	"github.com/craiggwilson/teakwood/frame"
	"github.com/craiggwilson/teakwood/label"
	"github.com/craiggwilson/teakwood/theme"
)

func main() {
	mm := newMainModel()

	examples.Run(mm, tea.WithAltScreen())
}

func newMainModel() *mainModel {
	return &mainModel{
		root: frame.New(label.New("Hello, World!")),
		styler: theme.New().
			Set("frame", lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder(), true),
			).
			Build(),
	}
}

type mainModel struct {
	root teakwood.Visual

	bounds teakwood.Rectangle
	styler teakwood.Styler
}

func (m *mainModel) Init() tea.Cmd {
	return m.root.Init(m.styler)
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch tmsg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case tmsg.String() == "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.bounds = teakwood.NewRectangle(0, 0, tmsg.Width, tmsg.Height)
		log.Println(m.bounds.Width, m.bounds.Height)
	}

	m.root, cmd = m.root.Update(msg, m.bounds)

	return m, cmd
}

func (m *mainModel) View() string {
	return m.root.View()
}