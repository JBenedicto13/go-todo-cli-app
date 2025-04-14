package main

import (
    "flag"
    "fmt"
)

type CmdFlags struct {
    add    string
    list   bool
    del    int
    toggle int
}

func NewCmdFlags() *CmdFlags {
    cmdFlags := &CmdFlags{}
    
    flag.StringVar(&cmdFlags.add, "add", "", "Add a new todo")
    flag.BoolVar(&cmdFlags.list, "list", false, "List all todos")
    flag.IntVar(&cmdFlags.del, "del", -1, "Delete a todo by index")
    flag.IntVar(&cmdFlags.toggle, "toggle", -1, "Toggle completion status of a todo by index")
    
    flag.Parse()
    
    return cmdFlags
}

func (c *CmdFlags) Execute(todos *Todos) {
    if c.add != "" {
        todos.add(c.add)
        fmt.Println("Todo added successfully!")
    } else if c.list {
        if len(*todos) == 0 {
            fmt.Println("No todos found!")
            return
        }
        
        todos.print()
    } else if c.del >= 0 {
        err := todos.delete(c.del)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Println("Todo deleted successfully!")
    } else if c.toggle >= 0 {
        err := todos.toggle(c.toggle)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Println("Todo toggled successfully!")
    } else {
        flag.PrintDefaults()
    }
}