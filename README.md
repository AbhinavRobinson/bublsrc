# Bublsrc

A Go terminal user interface (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) that displays your fish shell command history with a clean, organized interface and powerful search capabilities.

## Features

- **Fish History Display**: Shows the last 5 commands from your fish shell history
- **Smart Search**: Automatic search mode when typing - no need to press `/` first
- **Clipboard Integration**: Copy selected commands to OS clipboard with visual feedback
- **Real-time Search**: Instant search results as you type
- **Beautiful UI**: Modern, colorful interface with elegant styling
- **Status Messages**: Visual feedback for copy operations with auto-hide
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

The application will automatically load your fish shell history and display the last 5 commands with timestamps.

### Controls

- **Navigation**: `↑/↓` or `Ctrl+J/K` to navigate through commands
- **Search**: Start typing to automatically enter search mode
- **Copy**: `Enter` to copy the selected command to clipboard
- **Search Mode**: `Esc` to exit search mode
- **Quit**: `Ctrl+C` to quit the application

### Features in Action

1. **Start the app** - See your recent fish history commands
2. **Start typing** - Automatically enters search mode with real-time filtering
3. **Navigate results** - Use arrow keys to select commands
4. **Copy commands** - Press Enter to copy with visual feedback
5. **Exit search** - Press Esc to return to history view

## Project Structure

```
bublsrc/
├── main.go                    # Entry point and main function
├── app.go                     # Main Bubble Tea model and application logic
├── fish_history.go            # Fish history UI components
├── fish_history_service.go    # Fish history business logic and data operations
├── search_service.go          # Search functionality and filtering
├── logger_service.go          # Custom logger service implementation
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
├── run.sh                     # Convenience script to run the application
└── debug.log                  # Debug log file (created at runtime)
```

## Code Overview

### Main Components

- **`main.go`**: Initializes the logger service, creates the app, and runs the Bubble Tea program
- **`app.go`**: Contains the main Bubble Tea model with Init, Update, and View methods, handles key events and clipboard operations
- **`fish_history.go`**: Handles fish history UI rendering and user interactions with beautiful styling
- **`fish_history_service.go`**: Manages fish history parsing, storage, and business logic
- **`search_service.go`**: Handles search functionality, filtering, and result management
- **`logger_service.go`**: Custom logger service with different log levels and formatted output

### Architecture

The application follows a clean separation of concerns:

- **UI Layer** (`app.go`, `fish_history.go`): Handles user interface and interactions
- **Service Layer** (`fish_history_service.go`, `search_service.go`, `logger_service.go`): Manages business logic and data operations
- **Data Layer**: Fish history parsing and storage within the service layer
- **Search Layer** (`search_service.go`): Handles search functionality, filtering, and result management

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
- Displays the last 5 commands with timestamps by default
- Shows recent commands when entering search mode
- Handles loading states and error conditions
- Formats commands with proper indexing and time display
- Provides real-time search filtering
- Enables clipboard integration for easy command copying

### Logging

The application uses a custom logger service that writes to both the console and `debug.log`. You can adjust the log level by changing the level parameter in `main.go`:

```go
logger := NewLoggerService(logFile, DEBUG) // Change to INFO, WARN, or ERROR
```

## Key Features

### Smart Search
- **Automatic Search Mode**: Just start typing to enter search mode
- **Real-time Filtering**: Instant results as you type
- **Default History**: Shows recent commands when no search query
- **Case-insensitive**: Searches work regardless of case

### Clipboard Integration
- **One-click Copy**: Press Enter to copy selected commands
- **Visual Feedback**: Status messages show copy success/failure
- **Auto-hide Messages**: Status messages disappear after 3 seconds
- **Cross-platform**: Works on macOS, Linux, and Windows

### Beautiful UI
- **Modern Design**: Clean, colorful interface with elegant styling
- **Responsive**: Adapts to terminal size
- **Status Indicators**: Clear visual feedback for all operations
- **Consistent Styling**: Professional look and feel

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - A powerful little TUI framework for Go
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for beautiful terminal applications
- [Atotto Clipboard](https://github.com/atotto/clipboard) - Cross-platform clipboard operations

## License

This project is open source and available under the [MIT License](LICENSE).
