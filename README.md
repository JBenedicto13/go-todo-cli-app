# Go Todo CLI Application

A simple command-line interface application for managing your todo list, built with Go.

## Description

A lightweight, terminal-based todo management tool that allows you to create, view, toggle, and delete tasks. The application stores your todos between sessions in a local JSON file, providing convenient task tracking directly from your command line.

## Getting Started

### Dependencies

- Go 1.18 or higher
- Make (optional, for using the Makefile commands)
- Windows, Linux, or macOS

### Installing
- Clone the repository:
  ```
  git clone https://github.com/yourusername/go-todo-cli-app.git
  cd go-todo-cli-app
  ```
- Install dependencies:
  ```
  go mod tidy
  ```
- Or using Make:
  ```
  make deps
  ```
### Executing program
- Build the application:
  ```
  go build todo-cli-app
  ```
- Using direct CLI commands:
  ```
  # Add a new todo
  ./todo-cli -add "Buy groceries"
  
  # List all todos
  ./todo-cli -list
  
  # Delete a todo (where 0 is the index)
  ./todo-cli -del 0
  
  # Toggle completion status (where 1 is the index)
  ./todo-cli -toggle 1
  ```
- Using Make shortcuts:
  ```
  # Add a new todo
  make add text="Buy groceries"
  
  # List all todos
  make list
  
  # Delete a todo
  make delete index=0
  
  # Toggle a todo's completion status
  make toggle index=1
  ```
## Help

If you encounter any issues with the application:

```
# Display available commands and their usage
./todo-cli
# or
make help
```

## Acknowledgments

- Built with Go's standard library
- Uses github.com/aquasecurity/table for formatted table output
- Uses go.uber.org/zap for logging
