package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle  = FocusedStyle
	NoStyle      = lipgloss.NewStyle()
	HelpStyle    = BlurredStyle

	FocusedButton = FocusedStyle.Render("[ Confirm ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Confirm"))
)
