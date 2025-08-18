package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	PROGRESS_BAR_WIDTH  = 71
	PROGRESS_BAR_CHAR   = "█"
	PROGRESS_EMPTY_CHAR = "░"
)

var (
	SpinnerStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7e22c9ff"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	progressEmpty = subtleStyle.Render(PROGRESS_EMPTY_CHAR)
	ramp          = lipgloss.NewStyle().Foreground(lipgloss.Color("#7e22c9ff"))
)

func ProgressBar(percent float64) string {
	w := float64(PROGRESS_BAR_WIDTH)

	fullSize := int(math.Round(w * percent))
	fullCells := ramp.Render(PROGRESS_BAR_CHAR)

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}
