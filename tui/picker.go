package tui

import (
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const TRUNCATE_MAX = 70

func Pick(items []string, selected *string) error {
	return huh.NewSelect[string]().
		Title("Pick a clip").
		OptionsFunc(func() []huh.Option[string] {
			opts := make([]huh.Option[string], len(items))
			for i, item := range items {
				opts[i] = huh.NewOption(truncateString(item, TRUNCATE_MAX), item)
			}
			return opts
		}, nil).
		Value(selected).
		Height(15).
		WithTheme(theme()).
		Run()
}

func theme() *huh.Theme {
	t := huh.ThemeBase()
	selected := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	t.Focused.Base = lipgloss.NewStyle().PaddingLeft(1)
	t.Focused.Title = lipgloss.NewStyle().Bold(true).PaddingTop(1)
	t.Focused.SelectSelector = selected.SetString("> ")
	t.Focused.SelectedOption = selected.Italic(true)
	return t
}

func truncateString(s string, max int) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > max {
		return s[:max] + "..."
	}
	return s
}
