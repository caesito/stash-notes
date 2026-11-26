package tui


import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)


type appModel struct {
	initialMenu []string
	cursor int
	selected map[int]struct{}
}

type menuModel struct {

}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down", "j":
			if m.cursor < len(m.initialMenu) - 1 {
				m.cursor++
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			_, ok := m.selected[m.cursor]

			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m appModel) View() tea.View {
	s := "Hello, from Stash.Notes\n"

	for i, itensMenu := range m.initialMenu {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := ""
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, itensMenu)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}

func initialappModel() appModel {
	return appModel{
		initialMenu: []string{"notes", "config", "quit"},
		selected: make(map[int]struct{}),
	}
}

func RunTui() {
	p := tea.NewProgram(initialappModel())

	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}



