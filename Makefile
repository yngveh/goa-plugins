

.PHONY: test
test: fmt
	go test -v ./logrus

.PHONY: fmt
fmt:
	go fmt ./...
