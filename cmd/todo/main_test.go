package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
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

func TestSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")

	if err := save(path, []Todo{{ID: 1, Title: "アライさん", Done: true}}); err != nil {
		t.Fatalf("save() error = %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := `[
  {
    "id": 1,
    "title": "アライさん",
    "done": true
  }
]`

	if string(got) != want {
		t.Errorf("save() の中身 =\n%s\nwant\n%s", got, want)
	}
}

func TestSaveOverwrite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")

	if err := save(path, []Todo{
		{ID: 1, Title: "todo 1", Done: true},
		{ID: 2, Title: "todo 2", Done: true},
		{ID: 3, Title: "todo 3", Done: true},
	}); err != nil {
		t.Fatalf("save() error = %v", err)
	}

	if err := save(path, []Todo{
		{ID: 4, Title: "todo 4", Done: true},
	}); err != nil {
		t.Fatalf("save() error = %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := `[
  {
    "id": 4,
    "title": "todo 4",
    "done": true
  }
]`

	if string(got) != want {
		t.Errorf("save() の中身 =\n%s\nwant\n%s", got, want)
	}
}

func TestSaveNil(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")

	if err := save(path, nil); err != nil {
		t.Fatalf("save() error = %v, want nil", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := `[]`

	if string(got) != want {
		t.Errorf("save() の中身 =\n%s\nwant\n%s", got, want)
	}
}

func TestLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")
	want := []Todo{{ID: 1, Title: "todo 1"}, {ID: 2, Title: "todo 2"}, {ID: 3, Title: "todo 3"}}

	if err := save(path, want); err != nil {
		t.Fatalf("save() error = %v, want nil", err)
	}

	got, err := load(path)
	if err != nil {
		t.Fatalf("load() error = %v", err)
	}

	if !slices.Equal(got, want) {
		t.Errorf("load() = %+v, want %+v", got, want)
	}
}

func TestLoadNotExist(t *testing.T) {
	path := filepath.Join(t.TempDir(), "not-exist.json")

	todos, err := load(path)

	if err != nil {
		t.Fatalf("load() error = %v, want nil", err)
	}

	if todos != nil {
		t.Errorf("load() = %+v, want nil", todos)
	}
}

func TestLoadBroken(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.json")

	if err := os.WriteFile(path, []byte("{💣"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := load(path); err == nil {
		t.Error("load() = nil, want error")
	}
}

func TestFormatTodos(t *testing.T) {
	tests := []struct {
		name  string
		todos []Todo
		want  string
	}{
		{
			name:  "完了",
			todos: []Todo{{ID: 1, Title: "あ", Done: true}},
			want:  "[x] 1 あ",
		},
		{
			name:  "未完了",
			todos: []Todo{{ID: 1, Title: "あ"}},
			want:  "[ ] 1 あ",
		},
		{
			name:  "完了済み混在",
			todos: []Todo{{ID: 1, Title: "あ"}, {ID: 2, Title: "い", Done: true}},
			want:  "[ ] 1 あ\n[x] 2 い",
		},
		{name: "nil", todos: nil, want: "TODO なし"},
		{name: "空スライス", todos: []Todo{}, want: "TODO なし"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatTodos(tt.todos)

			if got != tt.want {
				t.Errorf("formatTodos() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")

	var buf bytes.Buffer
	if err := run([]string{"add", "テスト"}, &buf, path); err != nil {
		t.Fatalf("run(add) error = %v", err)
	}

	buf.Reset()
	if err := run([]string{"list"}, &buf, path); err != nil {
		t.Fatalf("run(list) error = %v", err)
	}

	want := "[ ] 1 テスト\n"
	if got := buf.String(); got != want {
		t.Errorf("run(list) = %q, want %q", got, want)
	}

	if err := run([]string{"complete", "1"}, &buf, path); err != nil {
		t.Fatalf("run(complete) error = %v", err)
	}

	buf.Reset()
	if err := run([]string{"list"}, &buf, path); err != nil {
		t.Fatalf("run(list) error = %v", err)
	}

	want = "[x] 1 テスト\n"
	if got := buf.String(); got != want {
		t.Errorf("run(list) = %q, want %q", got, want)
	}

	if err := run([]string{"remove", "1"}, &buf, path); err != nil {
		t.Fatalf("run(remove) error = %v", err)
	}

	buf.Reset()
	if err := run([]string{"list"}, &buf, path); err != nil {
		t.Fatalf("run(list) error = %v", err)
	}

	want = "TODO なし\n"
	if got := buf.String(); got != want {
		t.Errorf("run(list) = %q, want %q", got, want)
	}
}

func TestRunErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "引数なし", args: []string{}},
		{name: "不明なコマンド", args: []string{"foo"}},
		{name: "add にタイトルなし", args: []string{"add"}},
		{name: "id が数字じゃない", args: []string{"remove", "abc"}},
		{name: "存在しない id", args: []string{"complete", "99"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "todos.json")
			var buf bytes.Buffer

			if err := run(tt.args, &buf, path); err == nil {
				t.Error("run() = nil, want error")
			}
		})
	}
}

func TestRunListDoesNotSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "todos.json")
	var buf bytes.Buffer

	if err := run([]string{"list"}, &buf, path); err != nil {
		t.Fatalf("run(list) error = %v", err)
	}

	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Error("list がファイルを作ってしまっている")
	}
}
