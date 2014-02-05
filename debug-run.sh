go build -o ./oascent src/main.go
gdb -tui ./oascent -ex "break main.go:1" -ex run
