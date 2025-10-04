# Bublsrc

A simple Go terminal user interface (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and featuring a custom logging system.

## Features

- **Terminal User Interface**: Built with Bubble Tea framework for interactive terminal applications
- **Custom Logger**: Implements a structured logging system with different log levels (DEBUG, INFO, WARN, ERROR)
- **Debug Logging**: Automatically logs to `debug.log` file for debugging purposes
- **Simple Interface**: Displays "Hello, World!" with keyboard controls

## Prerequisites

- Go 1.25.1 or later
- A terminal that supports TUI applications

## Installation

1. Clone the repository:
```bash
git clone https://github.com/abhinavrobinson/bublsrc.git
cd bublsrc
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

Run the application:
```bash
go run .
```

### Controls

- `q`, `Esc`, or `Ctrl+C`: Quit the application

## Project Structure

```
bublsrc/
├── main.go          # Entry point and main function
├── app.go           # Bubble Tea model and application logic
├── logger.go        # Custom logger implementation
├── go.mod           # Go module dependencies
├── go.sum           # Dependency checksums
└── debug.log        # Debug log file (created at runtime)
```

## Code Overview

### Main Components

- **`main.go`**: Initializes the logger, creates the app, and runs the Bubble Tea program
- **`app.go`**: Contains the Bubble Tea model with Init, Update, and View methods
- **`logger.go`**: Custom logger with different log levels and formatted output

### Logger Features

The custom logger supports:
- Multiple log levels: DEBUG, INFO, WARN, ERROR
- Formatted logging with `Debugf`, `Infof`, `Warnf`, `Errorf`
- File output to `debug.log`
- Timestamp and file location information

## Development

### Adding Features

To extend the application:

1. Modify the `Model` struct in `app.go` to add new state
2. Update the `Update` method to handle new messages
3. Modify the `View` method to display new content
4. Use the logger for debugging: `m.logger.Debug("Your debug message")`

### Logging

The application uses a custom logger that writes to both the console and `debug.log`. You can adjust the log level by changing the level parameter in `main.go`:

```go
logger := NewLogger(logFile, DEBUG) // Change to INFO, WARN, or ERROR
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - A powerful little TUI framework for Go

## License

This project is open source and available under the [MIT License](LICENSE).
