# Task Tracker CLI

A command-line interface (CLI) application to track and manage your tasks and to-do list. Built with Go and [Cobra](https://github.com/spf13/cobra).

**Project URL:** https://roadmap.sh/projects/task-tracker

## Features

- ✅ Add, Update, and Delete tasks
- ✅ Mark a task as in progress or done
- ✅ List all tasks
- ✅ List tasks by status (todo, in-progress, done)
- ✅ Data persistence using JSON file storage

## Task Properties

Each task contains the following properties:

| Property      | Description                                      |
|---------------|--------------------------------------------------|
| `id`          | A unique identifier for the task                 |
| `description` | A short description of the task                  |
| `status`      | The status of the task (`todo`, `in-progress`, `done`) |
| `createdAt`   | The date and time when the task was created      |
| `updatedAt`   | The date and time when the task was last updated |

## How It Works

Task Tracker CLI is built with:

- **[Go](https://go.dev/)** - Programming language
- **[Cobra](https://github.com/spf13/cobra)** - CLI framework for Go
- **JSON file storage** - Tasks are persisted in a local `data.json` file

### Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   CLI Commands  │────▶│  Helper Layer   │────▶│   data.json     │
│   (cmd/*.go)    │     │  (helper/*.go)  │     │   (storage)     │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

1. **Commands** (`cmd/`) - Define CLI commands using Cobra (add, list, update, delete, etc.)
2. **Helper** (`helper/`) - Handles data structures and JSON file read/write operations
3. **Storage** - Tasks are stored in `data.json` in the current directory

---

## Installation

### Prerequisites

- [Go 1.25+](https://go.dev/dl/) installed on your system

### Option 1: Go Install (Recommended)

```bash
go install github.com/plh97/task-cli@latest
```

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/plh97/task-cli.git
cd task-cli

# Build the binary
go build -o task-cli .

# (Optional) Move to your PATH for global access
sudo mv task-cli /usr/local/bin/
```

### Option 3: Download Binary

Download pre-built binaries from the [Releases](https://github.com/plh97/task-cli/releases) page.

---

## Running the CLI

### Quick Start

```bash
# If installed via go install or moved to PATH
task-cli add "My first task"
task-cli list

# If running from source directory
go run . add "My first task"
go run . list
```

## Usage

### Adding a New Task

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Updating a Task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### Deleting a Task

```bash
task-cli delete 1
```

### Marking Task Status

```bash
# Mark a task as in progress
task-cli mark-in-progress 1

# Mark a task as done
task-cli mark-done 1
```

### Listing Tasks

```bash
# List all tasks
task-cli list

# List tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Data Storage

Tasks are stored in a `data.json` file in the current directory. The file is automatically created if it does not exist.

Example `data.json` structure:

```json
{
  "data": [
    {
      "id": "abc123",
      "description": "Buy groceries",
      "status": "todo",
      "created_at": "2026-01-18T10:00:00Z",
      "updated_at": "2026-01-18T10:00:00Z"
    }
  ]
}
```

## Project Structure

```bash
task-cli/
├── cli_cobra_basic.go   # Main entry point
├── cmd/
│   ├── root.go          # Root command configuration
│   └── add.go           # Add task command
├── helper/
│   └── helper.go        # Data structures and file operations
├── data.json            # Task storage (auto-generated)
├── go.mod
└── README.md
```

---

## Development

### Setup Development Environment

```bash
# Clone the repository
git clone https://github.com/plh97/task-cli.git
cd task-cli

# Install dependencies
go mod download

# Run in development mode
go run . <command>
```

### Adding a New Command

1. Create a new file in `cmd/` directory (e.g., `cmd/mycommand.go`)
2. Define your command using Cobra:

```go
package cmd

import "github.com/spf13/cobra"

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description of my command",
    Run: func(cmd *cobra.Command, args []string) {
        // Your logic here
    },
}

func init() {
    rootCmd.AddCommand(myCmd)
}
```

### Building for Production

```bash
# Build for current platform
go build -o task-cli .

# Build for multiple platforms
GOOS=darwin GOARCH=arm64 go build -o task-cli-darwin-arm64 .
GOOS=linux GOARCH=amd64 go build -o task-cli-linux-amd64 .
GOOS=windows GOARCH=amd64 go build -o task-cli-windows-amd64.exe .
```

### Running Tests

```bash
go test ./...
```

---

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
