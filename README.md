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

## Installation

### Prerequisites

- Go 1.25+ installed on your system

### Build from Source

```bash
# Clone the repository
git clone https://github.com/liom-source/my-cli.git
cd my-cli

# Build the binary
go build -o task-cli cli_cobra_basic.go

# (Optional) Move to your PATH
mv task-cli /usr/local/bin/
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

```
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

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
