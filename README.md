# Bublsrc

A Go terminal user interface (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) that displays your fish shell command history with a clean, organized interface.

## Features

- **Fish History Display**: Shows the last 10 commands from your fish shell history
- **Terminal User Interface**: Built with Bubble Tea framework for interactive terminal applications
- **Custom Logger Service**: Implements a structured logging system with different log levels (DEBUG, INFO, WARN, ERROR)
- **Debug Logging**: Automatically logs to `debug.log` file for debugging purposes
- **Modular Architecture**: Clean separation of concerns with service and UI layers
- **Async Loading**: Loads fish history in the background for better user experience

## Prerequisites

- Go 1.25.1 or later
- A terminal that supports TUI applications
- Fish shell (for history functionality)

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

The application will automatically load your fish shell history and display the last 10 commands with timestamps.

### Controls

- `q`, `Esc`, or `Ctrl+C`: Quit the application

## Project Structure

```
bublsrc/
├── main.go                    # Entry point and main function
├── app.go                     # Main Bubble Tea model and application logic
├── fish_history.go            # Fish history UI components
├── fish_history_service.go    # Fish history business logic and data operations
├── logger_service.go          # Custom logger service implementation
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
└── debug.log                  # Debug log file (created at runtime)
```

## Code Overview

### Main Components

- **`main.go`**: Initializes the logger service, creates the app, and runs the Bubble Tea program
- **`app.go`**: Contains the main Bubble Tea model with Init, Update, and View methods
- **`fish_history.go`**: Handles fish history UI rendering and user interactions
- **`fish_history_service.go`**: Manages fish history parsing, storage, and business logic
- **`logger_service.go`**: Custom logger service with different log levels and formatted output

### Architecture

The application follows a clean separation of concerns:

- **UI Layer** (`app.go`, `fish_history.go`): Handles user interface and interactions
- **Service Layer** (`fish_history_service.go`, `logger_service.go`): Manages business logic and data operations
- **Data Layer**: Fish history parsing and storage within the service layer

### Logger Service Features

The custom logger service supports:
- Multiple log levels: DEBUG, INFO, WARN, ERROR
- Formatted logging with `Debugf`, `Infof`, `Warnf`, `Errorf`
- File output to `debug.log`
- Timestamp and file location information
- Service-based architecture for better modularity

## Development

### Adding Features

To extend the application:

1. **UI Changes**: Modify the `Model` struct in `app.go` or add new UI components in `fish_history.go`
2. **Business Logic**: Add new methods to `fish_history_service.go` for data operations
3. **New Services**: Create new service files following the `*_service.go` pattern
4. **Logging**: Use the logger service for debugging: `m.logger.Debug("Your debug message")`

### Fish History Integration

The application automatically:
- Parses fish history from `~/.local/share/fish/fish_history`
- Displays the last 10 commands with timestamps
- Handles loading states and error conditions
- Formats commands with proper indexing and time display

### Logging

The application uses a custom logger service that writes to both the console and `debug.log`. You can adjust the log level by changing the level parameter in `main.go`:

```go
logger := NewLoggerService(logFile, DEBUG) // Change to INFO, WARN, or ERROR
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - A powerful little TUI framework for Go

## License

This project is open source and available under the [MIT License](LICENSE).
