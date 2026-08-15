# go-cli-playground

## 構成

```
cmd/<tool-name>/main.go   各 CLI のエントリポイント
internal/                 複数 CLI で共有するコード
```

## 使い方

```bash
go run ./cmd/hello
go build ./cmd/hello
```
