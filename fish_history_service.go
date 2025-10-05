package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// FishCommand represents a single command from fish history
type FishCommand struct {
	Command string
	When    time.Time
}

// fishHistoryMsg represents a message containing fish history data
type fishHistoryMsg struct {
	commands []FishCommand
	err      error
}

// FishHistoryService handles all fish history operations
type FishHistoryService struct {
	logger        *LoggerService
	history       []FishCommand
	historyLoaded bool
}

// NewFishHistoryService creates a new fish history service
func NewFishHistoryService(logger *LoggerService) *FishHistoryService {
	return &FishHistoryService{
		logger:        logger,
		history:       []FishCommand{},
		historyLoaded: false,
	}
}

// LoadHistory loads and parses the fish history file
func (s *FishHistoryService) LoadHistory() ([]FishCommand, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	historyPath := filepath.Join(homeDir, ".local", "share", "fish", "fish_history")
	file, err := os.Open(historyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open fish history: %w", err)
	}
	defer file.Close()

	var commands []FishCommand
	scanner := bufio.NewScanner(file)
	var currentCmd FishCommand
	var inPaths bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "- cmd: ") {
			if currentCmd.Command != "" && !currentCmd.When.IsZero() {
				commands = append(commands, currentCmd)
			}
			currentCmd = FishCommand{
				Command: strings.TrimPrefix(line, "- cmd: "),
			}
			inPaths = false
		} else if strings.HasPrefix(line, "when: ") {
			timestampStr := strings.TrimPrefix(line, "when: ")
			if timestamp, err := strconv.ParseInt(timestampStr, 10, 64); err == nil {
				currentCmd.When = time.Unix(timestamp, 0)
			}
		} else if strings.HasPrefix(line, "paths:") {
			inPaths = true
		} else if inPaths && strings.HasPrefix(line, "- ") {
			// Skip path entries for now
			continue
		}
	}

	if currentCmd.Command != "" && !currentCmd.When.IsZero() {
		commands = append(commands, currentCmd)
	}

	// Sort by timestamp (newest first)
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].When.After(commands[j].When)
	})

	s.history = commands
	s.historyLoaded = true
	s.logger.Infof("Loaded %d fish history commands", len(commands))
	return commands, nil
}

// GetLastCommands returns the last N commands from the stored history
func (s *FishHistoryService) GetLastCommands(count int) []FishCommand {
	if len(s.history) < count {
		count = len(s.history)
	}
	return s.history[:count]
}

// IsHistoryLoaded returns whether the history has been loaded
func (s *FishHistoryService) IsHistoryLoaded() bool {
	return s.historyLoaded
}

// GetHistory returns the full stored history
func (s *FishHistoryService) GetHistory() []FishCommand {
	return s.history
}

// FormatCommand formats a single command for display
func (s *FishHistoryService) FormatCommand(cmd FishCommand, index int) string {
	return fmt.Sprintf("%d. %s\n   %s", index+1, cmd.Command, cmd.When.Format("2006-01-02 15:04:05"))
}

// CreateHistoryMessage creates a fishHistoryMsg for the given commands
func (s *FishHistoryService) CreateHistoryMessage(commands []FishCommand) fishHistoryMsg {
	return fishHistoryMsg{commands: commands}
}

// CreateErrorMessage creates a fishHistoryMsg for the given error
func (s *FishHistoryService) CreateErrorMessage(err error) fishHistoryMsg {
	return fishHistoryMsg{err: err}
}
