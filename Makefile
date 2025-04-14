run:
	go run .

deps:
	go mod tidy

add:
	go run . -add "$(text)"

list:
	go run . -list

delete:
	go run . -delete "$(index)"

toggle:
	go run . -toggle "$(index)"

help:
	@echo "Usage:"
	@echo "  make run           - Run the application"
	@echo "  make deps          - Update dependencies"
	@echo ""
	@echo "Todo Operations:"
	@echo "  make add text=\"Buy milk\"  - Add a new todo"
	@echo "  make list                   - List all todos"
	@echo "  make delete index=0         - Delete a todo by index"
	@echo "  make toggle index=1         - Toggle completion status of a todo"
	@echo ""
	@echo "Examples:"
	@echo "  make add text=\"Complete the project\""
	@echo "  make list"
	@echo "  make toggle index=0"