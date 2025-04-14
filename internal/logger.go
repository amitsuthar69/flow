package internal

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Log(_type string, message string) {

	switch _type {
	case "error":
		styles := log.DefaultStyles()
		styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
			SetString("ERROR").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#c90016")).
			Foreground(lipgloss.Color("0"))
		styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
		styles.Values["err"] = lipgloss.NewStyle().Bold(true)
		logger := log.New(os.Stderr)
		logger.SetStyles(styles)
		logger.Error(message)
	case "fatal":
		styles := log.DefaultStyles()
		styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
			SetString("FATAL!!").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#c90016")).
			Foreground(lipgloss.Color("0"))
		styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
		styles.Values["err"] = lipgloss.NewStyle().Bold(true)
		logger := log.New(os.Stderr)
		logger.SetStyles(styles)
		logger.Error(message)
		os.Exit(1)
	case "info":
		styles := log.DefaultStyles()
		styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
			SetString("INFO").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#08e8de")).
			Foreground(lipgloss.Color("0"))
		styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
		styles.Values["err"] = lipgloss.NewStyle().Bold(true)
		logger := log.New(os.Stderr)
		logger.SetStyles(styles)
		logger.Info(message)
	}
}
