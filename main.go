package main

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m, cmd := m.KeyCommand(msg)
		if cmd != nil {
			return m, cmd
		}
	case ClearErrorMsg:
		m.ClearErrCommand()
	case ThrowErrorMsg:
		return m.ThroeErrorCommand()
	case CompleteMsg:
		m.OpenDirCommand()
	case spinner.TickMsg:
		return m.ThickCommand(msg)
	}

	var cmd tea.Cmd
	m.Filepicker, cmd = m.Filepicker.Update(msg)

	if didSelect, path := m.Filepicker.DidSelectFile(msg); didSelect {
		m.SelectedFile = path
	}

	if didSelect, path := m.Filepicker.DidSelectDisabledFile(msg); didSelect {
		m.Err = errors.New(path + " is not valid.")
		m.SelectedFile = ""
		return m, tea.Batch(cmd, ClearErrorAfter(2*time.Second))
	}
	return m, cmd
}

func (m Model) View() string {
	if m.Quitting {
		return "\n  See you later!\n\n"
	}

	if m.Complete {
		return "\n  Complete! Out path: " + m.OutPath + "\n\n"
	}

	var s strings.Builder
	s.WriteString("\n  ")

	if m.Err != nil {
		s.WriteString(m.Filepicker.Styles.DisabledFile.Render(m.Err.Error()))
	} else if m.Processing {
		s.WriteString("\n" + "  " + m.Spinner.View() + "  Processing...\n\n")
	} else if m.SelectedFile == "" {
		s.WriteString("Pick a file:")
		s.WriteString("\n\n" + m.Filepicker.View() + "\n")
	} else {
		s.WriteString("Selected file: " + m.Filepicker.Styles.Selected.Render(m.SelectedFile))
		s.WriteString("\n\n" + m.Filepicker.View() + "\n")
		s.WriteString("Press enter to convert HLS")
	}
	return s.String()
}

func main() {
	p = tea.NewProgram(InitModel())
	if _, err := p.Run(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
