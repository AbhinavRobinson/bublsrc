package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	logger        *LoggerService
	historyUI     *FishHistoryUI
	searchService *SearchService
	// Search state
	searchMode bool
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

		if m.searchMode {
			// Handle search mode
			switch key {
			case "esc":
				m.logger.Info("Exiting search mode")
				m.searchMode = false
				m.searchService.Clear()
				return m, nil
			case "q", "ctrl+c":
				m.logger.Info("Quit command received")
				return m, tea.Quit
			case "up", "k":
				m.searchService.NavigateUp()
				return m, nil
			case "down", "j":
				m.searchService.NavigateDown()
				return m, nil
			case "backspace":
				if len(m.searchService.GetQuery()) > 0 {
					newQuery := m.searchService.GetQuery()[:len(m.searchService.GetQuery())-1]
					m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), newQuery)
				}
				return m, nil
			case "enter":
				// Could implement command execution here
				return m, nil
			default:
				// Add character to search query
				if len(key) == 1 {
					newQuery := m.searchService.GetQuery() + key
					m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), newQuery)
				}
				return m, nil
			}
		} else {
			// Handle normal mode
			switch key {
			case "q", "esc", "ctrl+c":
				m.logger.Info("Quit command received")
				return m, tea.Quit
			case "/":
				m.logger.Info("Entering search mode")
				m.searchMode = true
				m.searchService.Clear()
				return m, nil
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.searchMode {
		return m.historyUI.RenderSearchView(m.searchService.GetQuery(), m.searchService.GetResults(), m.searchService.GetIndex())
	}
	return m.historyUI.RenderHistoryView()
}

func NewApp(logger *LoggerService) *Model {
	historyService := NewFishHistoryService(logger)
	historyUI := NewFishHistoryUI(historyService)
	searchService := NewSearchService(logger)
	return &Model{
		logger:        logger,
		historyUI:     historyUI,
		searchService: searchService,
		searchMode:    false,
	}
}
