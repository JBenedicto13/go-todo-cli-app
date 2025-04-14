package main

import (
	"fmt"
	"go-todo-cli-app/internal/logger"
)


func main() {
	// Initialize logger
    logger.InitLogger()
    logger.Log.Info("Starting Todo CLI application")
    
    todos := Todos{}

	logger.Log.Debug("Attempting to load todos from storage")
    // Load data (todos) from storage:
    storage := NewStorage[Todos]("todos.json")
    err := storage.Load(&todos)
    if err != nil {
        logger.Log.Warnw("Could not load todos from storage", "error", err)
        fmt.Println("Warning: Could not load todos from storage. Starting fresh todos.")
    } else {
        logger.Log.Infof("Successfully loaded %d todos from storage", len(todos))
    }

    // Parse & Execute
	logger.Log.Debug("Parsing command line flags")
    cmdFlags := NewCmdFlags()
	logger.Log.Debug("Executing command")
    cmdFlags.Execute(&todos)

    // Save to storage
    logger.Log.Debug("Saving todos to storage")
    err = storage.Save(todos)
    if err != nil {
        logger.Log.Errorw("Failed to save todos to storage", "error", err)
        fmt.Printf("Error saving todos in storage: %v\n", err)
    } else {
        logger.Log.Infof("Successfully saved %d todos to storage", len(todos))
    }
    
    logger.Log.Info("Todo CLI application completed")
}