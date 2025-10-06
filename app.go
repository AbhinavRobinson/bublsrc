package main

import (
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
)

// statusMsg represents a status message update
type statusMsg struct {
	message string
}

// showStatus creates a status message that will be displayed
func (m Model) showStatus(message string) tea.Cmd {
	return func() tea.Msg {
		return statusMsg{message: message}
	}
}

type Model struct {
	logger        *LoggerService
	historyUI     *FishHistoryUI
	searchService *SearchService
	// Search state
	searchMode bool
	// History selection state
	historySelectedIndex int
	// Status message for UI feedback
	statusMessage string
	statusTimer   int
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
	case statusMsg:
		m.statusMessage = msg.message
		m.statusTimer = 3 // Show for 3 seconds
		return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
			return time.Time{}
		})
	case time.Time:
		if m.statusTimer > 0 {
			m.statusTimer--
			if m.statusTimer == 0 {
				// If this was a copy success message, quit the app
				if m.statusMessage == "✅ Copied to clipboard" {
					return m, tea.Quit
				}
				m.statusMessage = ""
			}
			if m.statusTimer > 0 {
				return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
					return time.Time{}
				})
			}
		}
		return m, nil
	case fishHistoryMsg:
		if msg.err != nil {
			m.logger.Errorf("Failed to load fish history: %v", msg.err)
			return m, nil
		}
		m.logger.Infof("Fish history loaded successfully")
		return m, nil
	case tea.WindowSizeMsg:
		// Handle window resizing
		m.historyUI.SetSize(msg.Width, msg.Height)
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
			case "ctrl+c":
				m.logger.Info("Quit command received")
				return m, tea.Quit
			case "up", "ctrl+k":
				m.searchService.NavigateUp()
				return m, nil
			case "down", "ctrl+j":
				m.searchService.NavigateDown()
				return m, nil
			case "backspace":
				if len(m.searchService.GetQuery()) > 0 {
					newQuery := m.searchService.GetQuery()[:len(m.searchService.GetQuery())-1]
					m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), newQuery)
				}
				return m, nil
			case "enter":
				// Copy selected command to clipboard
				selectedCmd := m.searchService.GetSelectedCommand()
				if selectedCmd != nil {
					err := clipboard.WriteAll(selectedCmd.Command)
					if err != nil {
						m.logger.Errorf("Failed to copy to clipboard: %v", err)
						return m, m.showStatus("❌ Copy failed")
					} else {
						m.logger.Infof("Copied command to clipboard: %s", selectedCmd.Command)
						m.statusMessage = "✅ Copied to clipboard"
						m.statusTimer = 1 // Show for 1 second then quit
						return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
							return time.Time{}
						})
					}
				}
				return m, nil
			default:
				// Add character to search query (including j, k, q)
				if len(key) == 1 {
					newQuery := m.searchService.GetQuery() + key
					m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), newQuery)
				}
				return m, nil
			}
		} else {
			// Handle normal mode
			switch key {
			case "esc", "ctrl+c":
				m.logger.Info("Quit command received")
				return m, tea.Quit
			case "/":
				m.logger.Info("Entering search mode")
				m.searchMode = true
				// Initialize with empty query to show last 5 commands
				m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), "")
				return m, nil
			case "up", "ctrl+k":
				if m.historySelectedIndex > 0 {
					m.historySelectedIndex--
					m.logger.Debugf("History navigation up: index=%d", m.historySelectedIndex)
				}
				return m, nil
			case "down", "ctrl+j":
				// Limit to top 5 commands (0-4)
				maxIndex := 4
				if m.historySelectedIndex < maxIndex {
					m.historySelectedIndex++
					m.logger.Debugf("History navigation down: index=%d", m.historySelectedIndex)
				}
				return m, nil
			case "enter":
				// Copy selected command to clipboard
				history := m.historyUI.service.GetHistory()
				if m.historySelectedIndex >= 0 && m.historySelectedIndex < len(history) {
					selectedCmd := history[m.historySelectedIndex]
					err := clipboard.WriteAll(selectedCmd.Command)
					if err != nil {
						m.logger.Errorf("Failed to copy to clipboard: %v", err)
						return m, m.showStatus("❌ Copy failed")
					} else {
						m.logger.Infof("Copied command to clipboard: %s", selectedCmd.Command)
						m.statusMessage = "✅ Copied to clipboard"
						m.statusTimer = 1 // Show for 1 second then quit
						return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
							return time.Time{}
						})
					}
				}
				return m, nil
			case "1", "2", "3", "4", "5":
				// Direct copy by number (1-5) - follow the existing numbering
				lastCommands := m.historyUI.service.GetLastCommands(5)
				commandNum, _ := strconv.Atoi(key)

				if commandNum >= 1 && commandNum <= len(lastCommands) {
					// The numbering is: 5=oldest(top), 4, 3, 2, 1=newest(bottom)
					// So: displayIndex = commandNum - 1, and we want the command at that displayIndex
					displayIndex := commandNum - 1
					selectedCmd := lastCommands[displayIndex]

					// Move selection to match the display position (i in the loop)
					// displayIndex = len(lastCommands) - 1 - i, so i = len(lastCommands) - 1 - displayIndex
					displayPosition := len(lastCommands) - 1 - displayIndex
					m.historySelectedIndex = displayPosition
					m.logger.Debugf("Pressed %d, displayIndex=%d, displayPosition=%d, copying command: %s", commandNum, displayIndex, displayPosition, selectedCmd.Command)

					err := clipboard.WriteAll(selectedCmd.Command)
					if err != nil {
						m.logger.Errorf("Failed to copy to clipboard: %v", err)
						return m, m.showStatus("❌ Copy failed")
					} else {
						m.logger.Infof("Copied command to clipboard: %s", selectedCmd.Command)
						m.statusMessage = "✅ Copied to clipboard"
						m.statusTimer = 2 // Show for 2 seconds so user can see the selection move
						return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
							return time.Time{}
						})
					}
				}
				return m, nil
			default:
				// Auto-enter search mode when typing
				if len(key) == 1 {
					m.logger.Info("Auto-entering search mode")
					m.searchMode = true
					// Initialize with empty query to show last 5 commands
					m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), "")
					// Add the typed character to the search query
					m.searchService.UpdateQuery(m.historyUI.service.GetHistory(), key)
					return m, nil
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var content string
	if m.searchMode {
		content = m.historyUI.RenderSearchView(m.searchService.GetQuery(), m.searchService.GetResults(), m.searchService.GetIndex())
	} else {
		content = m.historyUI.RenderHistoryView(m.historySelectedIndex)
	}

	// Add status message if present - positioned at the bottom
	if m.statusMessage != "" {
		content += "\n\n" + m.historyUI.RenderStatusMessage(m.statusMessage)
	}

	return content
}

func NewApp(logger *LoggerService) *Model {
	historyService := NewFishHistoryService(logger)
	historyUI := NewFishHistoryUI(historyService, logger)
	searchService := NewSearchService(logger)
	return &Model{
		logger:               logger,
		historyUI:            historyUI,
		searchService:        searchService,
		searchMode:           false,
		historySelectedIndex: 4, // Start with most recent command selected (last in list)
	}
}
