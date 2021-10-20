GO_CMD = go

NAME = "eval"
VERSION = "0.1"

build-linux:export GOOS=linux
build-linux:export GOARCH=amd64
build-linux:export CGO_ENABLED=0
build-linux:
	$(GO_CMD) build -o $(NAME) main.go

# darwin
build-linux:export GOOS=darwin
build-linux:export GOARCH=amd64
build-linux:export CGO_ENABLED=0
build-darwin:
	$(GO_CMD) build -o $(NAME) main.go

# windows nt
build-linux:export GOOS=windows
build-linux:export GOARCH=amd64
build-linux:export CGO_ENABLED=0
build-windows:
	$(GO_CMD) build -o $(NAME) main.go

# test
.PHONY:test
test:
	$(GO_CMD) test -test.v ./test/alg_test.go
	$(GO_CMD) test -test.v ./test/task_test.go

benchmark:
	$(GO_CMD) test -bench=. ./test/alg_test.go
	$(GO_CMD) test -bench=. ./test/task_test.go