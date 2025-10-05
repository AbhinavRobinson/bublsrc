package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	logger    *LoggerService
	historyUI *FishHistoryUI
}

func (m Model) Init() tea.Cmd {
	m.logger.Info("Model initialized")
	return m.loadFishHistory
}

func (m Model) loadFishHistory() tea.Msg {
	return m.historyUI.LoadHistoryMessage()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case fishHistoryMsg:
		if msg.err != nil {
			m.logger.Errorf("Failed to load fish history: %v", msg.err)
			return m, nil
		}
		m.logger.Infof("Fish history loaded successfully")
		return m, nil
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
	return m.historyUI.RenderHistoryView()
}

func NewApp(logger *LoggerService) *Model {
	historyService := NewFishHistoryService(logger)
	historyUI := NewFishHistoryUI(historyService)
	return &Model{
		logger:    logger,
		historyUI: historyUI,
	}
}
