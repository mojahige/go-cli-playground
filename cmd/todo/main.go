package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func find(todos []Todo, id int) (int, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return i, nil
		}
	}

	return 0, fmt.Errorf("id %d のタスクが無いのだ", id)
}

func nextID(todos []Todo) int {
	max := 0

	for _, todo := range todos {
		if todo.ID > max {
			max = todo.ID
		}
	}

	return max + 1
}

func add(todos []Todo, title string) []Todo {
	return append(todos, Todo{
		ID:    nextID(todos),
		Title: title,
		Done:  false,
	})
}

func remove(todos []Todo, id int) ([]Todo, error) {
	i, err := find(todos, id)

	if err != nil {
		return todos, err
	}

	return slices.Delete(todos, i, i+1), nil
}

func complete(todos []Todo, id int) ([]Todo, error) {
	i, err := find(todos, id)

	if err != nil {
		return todos, err
	}

	todos[i].Done = true

	return todos, nil
}

func save(path string, todos []Todo) error {
	if todos == nil {
		todos = []Todo{}
	}

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func load(path string) ([]Todo, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
