package main

import (
	"errors"
	"time"

	"github.com/beevik/guid"
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
			m.OutPath = guid.NewString()
			go m.Convert()
			return m, m.Spinner.Tick
		}
	}
	return m, nil
}

func (m *Model) ClearErrCommand() {
	m.Err = nil
}

func (m Model) ThroeErrorCommand() (tea.Model, tea.Cmd) {
	m.Err = errors.ErrUnsupported
	m.Processing = false
	m.SelectedFile = ""
	return m, ClearErrorAfter(2 * time.Second)
}

func (m *Model) OpenDirCommand() {
	m.Complete = true
	OpenDir(m.OutPath)
}

func (m Model) ThickCommand(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Spinner, cmd = m.Spinner.Update(msg)
	return m, cmd
}
