build:
	docker compose up -d --build --remove-orphans

swag:
	go fmt ./...
	swag fmt -g cmd/main.go -d ./
	swag init --pd --parseInternal -g cmd/main.go

# This one spesifically for local development on windows
dev:
	compiledaemon -command="./bin/server.exe" -build="go build -o ./bin/server.exe ./cmd/main.go"

run:
	go build -o ./bin/server ./cmd/main.go
	./bin/server

lint:
	golangci-lint run --fix

prod:
	go build -v -ldflags='-s -w' -o server cmd/main.go

prod-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.2
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w .
	golangci-lint run

pre-all:
	pre-commit run --all-files

configure-pre-commit:
	pip install pre-commit
	pre-commit install
	pre-commit autoupdate