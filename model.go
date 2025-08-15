package main

import (
	"os"
	"time"

	"github.com/beevik/guid"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Spinner      spinner.Model
	Filepicker   filepicker.Model
	SelectedFile string
	OutPath      string
	Quitting     bool
	Processing   bool
	Complete     bool
	IsGPU        bool
	Err          error
}

type ClearErrorMsg struct{}
type ThrowErrorMsg struct{}
type CompleteMsg struct{}

func (m Model) Init() tea.Cmd {
	return m.Filepicker.Init()
}

func InitModel() Model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = SpinnerStyle

	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.Getwd()

	return Model{s, fp, "", guid.NewString(), false, false, false, os.Args[len(os.Args)-1] == "GPU", nil}
}

func ClearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearErrorMsg{}
	})
}

func ThrowError() tea.Cmd {
	return func() tea.Msg {
		return ThrowErrorMsg{}
	}
}
