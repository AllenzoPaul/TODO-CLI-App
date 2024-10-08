package main

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time //it is a pointed reference because it can be null
}

//declaring the todos struct of type Todo
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	//append todo to todos struct
	*todos = append(*todos, todo)
}
