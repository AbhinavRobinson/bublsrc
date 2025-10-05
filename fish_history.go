package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define beautiful styles using LipGloss
var (
	// Colors
	primaryColor   = lipgloss.Color("#00D4AA") // Teal
	secondaryColor = lipgloss.Color("#7C3AED") // Purple
	accentColor    = lipgloss.Color("#F59E0B") // Amber
	textColor      = lipgloss.Color("#F8FAFC") // Light gray
	mutedColor     = lipgloss.Color("#64748B") // Slate
	errorColor     = lipgloss.Color("#EF4444") // Red
	successColor   = lipgloss.Color("#10B981") // Green

	// Header styles
	headerStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Margin(1, 0).
			Align(lipgloss.Center)

	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Margin(0, 1)

	// Command styles
	commandStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Margin(0, 2)

	commandNumberStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true)

	commandTextStyle = lipgloss.NewStyle().
				Foreground(textColor)

	timestampStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)

	// Search styles
	searchBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(0, 1).
			Margin(1, 0)

	searchPromptStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true)

	// List styles
	listStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondaryColor).
			Padding(1, 2).
			Margin(1, 0)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				Background(lipgloss.Color("#1E293B")).
				Padding(0, 1)

	// Status styles
	statusStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Margin(1, 0)

	// Status message styles - more subtle and elegant
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true).
				Background(lipgloss.Color("#0F172A")).
				Padding(0, 1).
				Margin(0, 2).
				Align(lipgloss.Center)

	statusErrorStyle = lipgloss.NewStyle().
				Foreground(errorColor).
				Bold(true).
				Background(lipgloss.Color("#0F172A")).
				Padding(0, 1).
				Margin(0, 2).
				Align(lipgloss.Center)

	loadingStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			Margin(1, 0)

	// Help styles
	helpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Margin(1, 0)

	keyStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	// Container styles
	containerStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Margin(1, 0)
)

// FishHistoryUI handles fish history specific UI operations
type FishHistoryUI struct {
	service     *FishHistoryService
	list        list.Model
	searchInput textinput.Model
	viewport    viewport.Model
	width       int
	height      int
	logger      *LoggerService
}

// NewFishHistoryUI creates a new fish history UI component
func NewFishHistoryUI(service *FishHistoryService, logger *LoggerService) *FishHistoryUI {
	// Initialize search input
	searchInput := textinput.New()
	searchInput.Placeholder = "Search commands..."
	searchInput.Focus()
	searchInput.CharLimit = 100
	searchInput.Width = 50

	// Initialize viewport
	vp := viewport.New(80, 20)

	// Initialize list
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "🐟 Fish History"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return &FishHistoryUI{
		service:     service,
		list:        l,
		searchInput: searchInput,
		viewport:    vp,
		width:       80,
		height:      20,
		logger:      logger,
	}
}

// RenderHistoryView renders the fish history view with beautiful styling
func (ui *FishHistoryUI) RenderHistoryView(selectedIndex int) string {
	if !ui.service.IsHistoryLoaded() {
		loading := loadingStyle.Render("🔄 Loading fish history...")
		help := helpStyle.Render("Press " + keyStyle.Render("q") + " to quit")
		return containerStyle.Render(loading + "\n\n" + help)
	}

	// Create beautiful header
	header := headerStyle.Render("🐟 Fish History")
	subtitle := statusStyle.Render("Last 5 Commands")

	// Get last 5 commands using service
	lastCommands := ui.service.GetLastCommands(5)

	// Create command list with beautiful styling
	var commands []string
	for i, cmd := range lastCommands {
		var prefix string
		if i == selectedIndex {
			prefix = selectedItemStyle.Render("▶")
		} else {
			prefix = "  "
		}

		number := commandNumberStyle.Render(fmt.Sprintf("%d.", i+1))
		command := commandTextStyle.Render(cmd.Command)
		timestamp := timestampStyle.Render(cmd.When.Format("2006-01-02 15:04:05"))

		commandLine := fmt.Sprintf("%s %s %s\n   %s", prefix, number, command, timestamp)
		commands = append(commands, commandLine)
	}

	// Join commands with spacing
	commandList := strings.Join(commands, "\n\n")

	// Create help text
	help := helpStyle.Render("Press " + keyStyle.Render("q") + " to quit, " + keyStyle.Render("↑/↓") + " or " + keyStyle.Render("Ctrl+J/K") + " to navigate, " + keyStyle.Render("Enter") + " to copy, " + keyStyle.Render("type") + " to search")

	// Combine everything
	content := header + "\n" + subtitle + "\n\n" + commandList + "\n\n" + help

	return containerStyle.Render(content)
}

// RenderSearchView renders the search results view with beautiful styling
func (ui *FishHistoryUI) RenderSearchView(query string, results []FishCommand, selectedIndex int) string {
	ui.logger.Debugf("RenderSearchView: query='%s', results=%d, selectedIndex=%d", query, len(results), selectedIndex)

	if !ui.service.IsHistoryLoaded() {
		loading := loadingStyle.Render("🔄 Loading fish history...")
		help := helpStyle.Render("Press " + keyStyle.Render("q") + " to quit")
		return containerStyle.Render(loading + "\n\n" + help)
	}

	// Create beautiful header
	var header string
	if query == "" {
		header = headerStyle.Render("🐟 Fish History")
	} else {
		header = headerStyle.Render("🔍 Search Results")
	}

	// Create search query display
	var queryDisplay string
	if query == "" {
		queryDisplay = statusStyle.Render("Recent Commands")
	} else {
		queryDisplay = searchPromptStyle.Render("Query: ") + commandTextStyle.Render(query)
	}

	var content string

	if len(results) == 0 {
		noResults := statusStyle.Render("No commands found matching your search.")
		content = header + "\n\n" + queryDisplay + "\n\n" + noResults
	} else {
		// Limit results to top 5
		displayResults := results
		if len(results) > 5 {
			displayResults = results[:5]
		}

		// Create results count
		totalCount := len(results)
		displayCount := len(displayResults)
		var countText string
		if query == "" {
			countText = fmt.Sprintf("Showing %d recent commands:", displayCount)
		} else if totalCount > 5 {
			countText = fmt.Sprintf("Found %d matching commands (showing top %d):", totalCount, displayCount)
		} else {
			countText = fmt.Sprintf("Found %d matching commands:", displayCount)
		}
		count := statusStyle.Render(countText)

		// Create results list with beautiful styling
		var resultItems []string
		for i, cmd := range displayResults {
			var prefix string
			// Only adjust selectedIndex if it's out of bounds
			displaySelectedIndex := selectedIndex
			if selectedIndex >= len(displayResults) {
				displaySelectedIndex = len(displayResults) - 1
			} else if selectedIndex < 0 {
				displaySelectedIndex = 0
			}

			if i == displaySelectedIndex {
				prefix = selectedItemStyle.Render("▶")
			} else {
				prefix = "  "
			}

			number := commandNumberStyle.Render(fmt.Sprintf("%d.", i+1))
			command := commandTextStyle.Render(cmd.Command)
			timestamp := timestampStyle.Render(cmd.When.Format("2006-01-02 15:04:05"))

			resultLine := fmt.Sprintf("%s %s %s\n   %s", prefix, number, command, timestamp)
			resultItems = append(resultItems, resultLine)
		}

		resultsList := strings.Join(resultItems, "\n\n")
		content = header + "\n\n" + queryDisplay + "\n\n" + count + "\n\n" + resultsList
	}

	// Create help text
	var help string
	if query == "" {
		help = helpStyle.Render("Press " + keyStyle.Render("Ctrl+C") + " to quit, " + keyStyle.Render("esc") + " to exit, " + keyStyle.Render("↑/↓") + " or " + keyStyle.Render("Ctrl+J/K") + " to navigate, " + keyStyle.Render("Enter") + " to copy, " + keyStyle.Render("type") + " to search")
	} else {
		help = helpStyle.Render("Press " + keyStyle.Render("Ctrl+C") + " to quit, " + keyStyle.Render("esc") + " to exit search, " + keyStyle.Render("↑/↓") + " or " + keyStyle.Render("Ctrl+J/K") + " to navigate, " + keyStyle.Render("Enter") + " to copy")
	}

	// Combine everything
	fullContent := content + "\n\n" + help

	return containerStyle.Render(fullContent)
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

// SetSize updates the UI dimensions
func (ui *FishHistoryUI) SetSize(width, height int) {
	ui.width = width
	ui.height = height
	ui.viewport.Width = width
	ui.viewport.Height = height - 4 // Leave space for header and help
	ui.list.SetWidth(width)
	ui.list.SetHeight(height - 4)
}

// GetSearchInput returns the search input model
func (ui *FishHistoryUI) GetSearchInput() textinput.Model {
	return ui.searchInput
}

// UpdateSearchInput updates the search input
func (ui *FishHistoryUI) UpdateSearchInput(input textinput.Model) {
	ui.searchInput = input
}

// RenderSearchInput renders just the search input with styling
func (ui *FishHistoryUI) RenderSearchInput() string {
	searchBox := searchBoxStyle.Render(ui.searchInput.View())
	prompt := searchPromptStyle.Render("Search: ")
	return prompt + searchBox
}

// RenderStatusMessage renders a status message with appropriate styling
func (ui *FishHistoryUI) RenderStatusMessage(message string) string {
	if strings.Contains(message, "❌") {
		return statusErrorStyle.Render(message)
	}
	return statusMessageStyle.Render(message)
}
