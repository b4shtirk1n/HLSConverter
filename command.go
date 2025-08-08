package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) KeyCommand(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		m.Quitting = true
		if cmd != nil {
			cmd.Process.Kill()
		}
		return m, tea.Quit
	case "enter", " ":
		if !m.Processing && m.SelectedFile != "" {
			m.Processing = true
			go m.Convert()
			return m, m.Spinner.Tick
		}
	}
	return m, nil
}

func (m *Model) ClearErrCommand() {
	m.Err = nil
}
