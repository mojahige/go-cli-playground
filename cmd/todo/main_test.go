package main

import (
	"slices"
	"testing"
)

func TestFind(t *testing.T) {
	todos := []Todo{{ID: 1}, {ID: 2}, {ID: 3}}

	tests := []struct {
		name    string
		todos   []Todo
		id      int
		want    int
		wantErr bool
	}{
		{name: "先頭", todos: todos, id: 1, want: 0},
		{name: "途中", todos: todos, id: 2, want: 1},
		{name: "末尾", todos: todos, id: 3, want: 2},
		{name: "存在しない", todos: todos, id: 99, wantErr: true},
		{name: "空", todos: nil, id: 1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := find(tt.todos, tt.id)

			if (err != nil) != tt.wantErr {
				t.Fatalf("find() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if got != tt.want {
				t.Errorf("find() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNextID(t *testing.T) {
	tests := []struct {
		name  string
		todos []Todo
		want  int
	}{
		{name: "連番", todos: []Todo{{ID: 1}, {ID: 2}}, want: 3},
		{name: "歯抜け", todos: []Todo{{ID: 1}, {ID: 3}}, want: 4},
		{name: "順不同", todos: []Todo{{ID: 3}, {ID: 1}}, want: 4},
		{name: "nil", todos: nil, want: 1},
		{name: "空スライス", todos: []Todo{}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextID(tt.todos)

			if got != tt.want {
				t.Errorf("nextID() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	todos := add(nil, "Go を練習するのだ")

	if got := len(todos); got != 1 {
		t.Errorf("add() の件数 = %d, want %d", got, 1)
	}

	if got := todos[0].ID; got != 1 {
		t.Errorf("todos[0].ID = %d, want %d", got, 1)
	}

	if got := todos[0].Title; got != "Go を練習するのだ" {
		t.Errorf("todos[0].Title = %q, want %q", got, "Go を練習するのだ")
	}

	todos = add(todos, "もう1件足すのだ")

	if got := len(todos); got != 2 {
		t.Errorf("2件目追加後の件数 = %d, want %d", got, 2)
	}

	if got := todos[1].ID; got != 2 {
		t.Errorf("todos[1].ID = %d, want %d", got, 2)
	}

	if got := todos[1].Title; got != "もう1件足すのだ" {
		t.Errorf("todos[1].Title = %q, want %q", got, "もう1件足すのだ")
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		todos   []Todo
		id      int
		want    []Todo
		wantErr bool
	}{
		{name: "先頭を削除", todos: []Todo{{ID: 1}, {ID: 2}, {ID: 3}}, id: 2, want: []Todo{{ID: 1}, {ID: 3}}},
		{name: "末尾を削除", todos: []Todo{{ID: 1}, {ID: 2}, {ID: 3}}, id: 3, want: []Todo{{ID: 1}, {ID: 2}}},
		{name: "間を削除", todos: []Todo{{ID: 1}, {ID: 2}, {ID: 3}}, id: 2, want: []Todo{{ID: 1}, {ID: 3}}},
		{name: "存在しない", todos: []Todo{{ID: 1}}, id: 99, wantErr: true},
		{name: "空", todos: nil, id: 1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := remove(tt.todos, tt.id)

			if (err != nil) != tt.wantErr {
				t.Fatalf("remove() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("remove() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestComplete(t *testing.T) {
	tests := []struct {
		name    string
		todos   []Todo
		id      int
		want    []Todo
		wantErr bool
	}{
		{name: "先頭を完了", todos: []Todo{{ID: 1}, {ID: 2}, {ID: 3}}, id: 1, want: []Todo{{ID: 1, Done: true}, {ID: 2}, {ID: 3}}},
		{name: "末尾を完了", todos: []Todo{{ID: 1}, {ID: 2}, {ID: 3}}, id: 3, want: []Todo{{ID: 1}, {ID: 2}, {ID: 3, Done: true}}},
		{name: "間を完了", todos: []Todo{{ID: 1}, {ID: 2}, {ID: 3}}, id: 2, want: []Todo{{ID: 1}, {ID: 2, Done: true}, {ID: 3}}},
		{name: "存在しない", todos: []Todo{{ID: 1}}, id: 99, wantErr: true},
		{name: "空", todos: nil, id: 1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := complete(tt.todos, tt.id)

			if (err != nil) != tt.wantErr {
				t.Fatalf("complete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("complete() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
