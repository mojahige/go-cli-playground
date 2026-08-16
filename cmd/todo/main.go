package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
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

	return 0, fmt.Errorf("id %d のタスク無し", id)
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

func reopen(todos []Todo, id int) ([]Todo, error) {
	i, err := find(todos, id)

	if err != nil {
		return todos, err
	}

	todos[i].Done = false

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

func formatTodos(todos []Todo) string {
	if len(todos) == 0 {
		return "TODO なし"
	}

	lines := make([]string, 0, len(todos))

	for _, todo := range todos {
		done := " "
		if todo.Done {
			done = "x"
		}

		lines = append(lines, fmt.Sprintf("[%s] %d %s", done, todo.ID, todo.Title))
	}

	return strings.Join(lines, "\n")
}

func run(args []string, w io.Writer, path string) error {
	if len(args) == 0 {
		return errors.New("使い方: todo <list|add|remove|complete> ...")
	}

	todos, err := load(path)

	if err != nil {
		return err
	}

	switch args[0] {
	case "list":
		fmt.Fprintln(w, formatTodos(todos))

		return nil
	case "add":
		if len(args) < 2 {
			return errors.New("使い方: todo add <title>")
		}

		todos = add(todos, args[1])
	case "remove":
		if len(args) < 2 {
			return errors.New("使い方: todo remove <id>")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		todos, err = remove(todos, id)
		if err != nil {
			return err
		}
	case "complete":
		if len(args) < 2 {
			return errors.New("使い方: todo complete <id>")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		todos, err = complete(todos, id)
		if err != nil {
			return err
		}
	case "reopen":
		if len(args) < 2 {
			return errors.New("使い方: todo reopen <id>")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		todos, err = reopen(todos, id)
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown command: " + args[0])
	}

	return save(path, todos)
}

func main() {
	if err := run(os.Args[1:], os.Stdout, "./todos.json"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
