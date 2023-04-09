# !/bin/zsh
go run $(ls -1 cmd/web/*.go | grep -v _test.go) -addr="$1"