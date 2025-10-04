package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	logger *Logger
}

func (m Model) Init() tea.Cmd {
	m.logger.Info("Model initialized")
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		m.logger.Debugf("Key pressed: %s", key)
		switch key {
		case "q", "esc", "ctrl+c":
			m.logger.Info("Quit command received")
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "Hello, World!"
}

func NewApp(logger *Logger) *Model {
	return &Model{
		logger: logger,
	}
}
