.PHONY: run dev swag

MAIN_FILE := main.go

run: swag
	go run $(MAIN_FILE)

dev: swag
	export PATH=$PATH:$(go env GOPATH)/bin
	swag init -g $(MAIN_FILE) -d .
	go run $(MAIN_FILE)

swag:
	swag init -g $(MAIN_FILE) -d .