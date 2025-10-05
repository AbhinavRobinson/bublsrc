package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// FishHistoryUI handles fish history specific UI operations
type FishHistoryUI struct {
	service *FishHistoryService
}

// NewFishHistoryUI creates a new fish history UI component
func NewFishHistoryUI(service *FishHistoryService) *FishHistoryUI {
	return &FishHistoryUI{
		service: service,
	}
}

// RenderHistoryView renders the fish history view
func (ui *FishHistoryUI) RenderHistoryView() string {
	if !ui.service.IsHistoryLoaded() {
		return "Loading fish history...\n\nPress 'q' to quit"
	}

	var sb strings.Builder
	sb.WriteString("🐟 Fish History - Last 10 Commands\n")
	sb.WriteString("=====================================\n\n")

	// Get last 10 commands using service
	lastCommands := ui.service.GetLastCommands(10)

	for i, cmd := range lastCommands {
		sb.WriteString(ui.service.FormatCommand(cmd, i))
		sb.WriteString("\n\n")
	}

	sb.WriteString("Press 'q' to quit")
	return sb.String()
}

// LoadHistoryMessage creates a message for loading fish history
func (ui *FishHistoryUI) LoadHistoryMessage() tea.Msg {
	commands, err := ui.service.LoadHistory()
	if err != nil {
		return ui.service.CreateErrorMessage(err)
	}
	return ui.service.CreateHistoryMessage(commands)
}

// IsHistoryLoaded checks if history is loaded
func (ui *FishHistoryUI) IsHistoryLoaded() bool {
	return ui.service.IsHistoryLoaded()
}
