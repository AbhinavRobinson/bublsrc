package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	logFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := NewLogger(logFile, DEBUG)

	logger.Info("Program started")

	app := NewApp(logger)

	if _, err := tea.NewProgram(app).Run(); err != nil {
		logger.Errorf("Error running program: %v", err)
		os.Exit(1)
	}

	logger.Info("Program ended")
}
