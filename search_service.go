package main

import (
	"strings"
)

// SearchService handles all search-related operations
type SearchService struct {
	logger *LoggerService
	// Internal search state
	query   string
	results []FishCommand
	index   int
}

// NewSearchService creates a new search service
func NewSearchService(logger *LoggerService) *SearchService {
	return &SearchService{
		logger:  logger,
		query:   "",
		results: []FishCommand{},
		index:   0,
	}
}

// UpdateQuery updates the search query and results
func (s *SearchService) UpdateQuery(commands []FishCommand, query string) {
	s.query = query
	s.results = s.searchCommands(commands, query)
	s.index = 0
	s.logger.Debugf("Search query updated to '%s', found %d results", query, len(s.results))
}

// NavigateUp moves the selection up in the results
func (s *SearchService) NavigateUp() {
	if s.index > 0 {
		s.index--
	}
	s.logger.Debugf("NavigateUp: index=%d, results=%d", s.index, len(s.results))
}

// NavigateDown moves the selection down in the results
func (s *SearchService) NavigateDown() {
	// Limit navigation to top 5 results for display
	maxIndex := len(s.results) - 1
	if maxIndex > 4 {
		maxIndex = 4 // Limit to top 5 (0-4)
	}
	if s.index < maxIndex {
		s.index++
	}
	s.logger.Debugf("NavigateDown: index=%d, maxIndex=%d, results=%d", s.index, maxIndex, len(s.results))
}

// GetSelectedCommand returns the currently selected command
func (s *SearchService) GetSelectedCommand() *FishCommand {
	if len(s.results) == 0 || s.index < 0 || s.index >= len(s.results) {
		return nil
	}
	return &s.results[s.index]
}

// HasResults returns true if there are search results
func (s *SearchService) HasResults() bool {
	return len(s.results) > 0
}

// GetResultCount returns the number of search results
func (s *SearchService) GetResultCount() int {
	return len(s.results)
}

// GetQuery returns the current search query
func (s *SearchService) GetQuery() string {
	return s.query
}

// GetResults returns the current search results
func (s *SearchService) GetResults() []FishCommand {
	return s.results
}

// GetIndex returns the current selection index
func (s *SearchService) GetIndex() int {
	return s.index
}

// Clear resets the search state
func (s *SearchService) Clear() {
	s.query = ""
	s.results = []FishCommand{}
	s.index = 0
}

// searchCommands is the internal search implementation
func (s *SearchService) searchCommands(commands []FishCommand, query string) []FishCommand {
	if query == "" {
		return []FishCommand{}
	}

	query = strings.ToLower(query)
	var results []FishCommand

	for _, cmd := range commands {
		if strings.Contains(strings.ToLower(cmd.Command), query) {
			results = append(results, cmd)
		}
	}

	return results
}
