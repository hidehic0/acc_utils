package submit_select_ui

import (
	cmdFn "hidehic0/acc_utils/internal/cmd"
	"hidehic0/acc_utils/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	items    []string
	selected int
}

func InitalModel() model {
	dirs := utils.GetTasks()
	return model{
		items:  dirs,
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmdFn.SubmitFn(m.items[m.cursor])
			return m, tea.Quit
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.cursor == len(m.items)-1 {
				m.cursor = 0
			} else {
				m.cursor++
			}
		case "k", "up":
			if m.cursor == 0 {
				m.cursor = len(m.items) - 1
			} else {
				m.cursor--
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Which one do you want to submit?"

	for i, item := range m.items {
		s += "\n"
		if i == m.cursor {
			s += "> "
		} else {
			s += "  "
		}

		s += item
	}

	return s
}
