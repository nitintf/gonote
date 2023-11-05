.PHONY: gomod run

gomod:
	go mod tidy

run:
	go run cmd/gonote.go
