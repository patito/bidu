NAME    = bidu
BIN     = $(GOPATH)/bin/$(NAME)

DREDD_BIN:=dredd
DREDD_API:=api.md
DREDD_LOCAL:=http://localhost:3000

all: build

test: deps
		go test -cover -race -v ./...

build: deps test
		go build -o $(GOPATH)/$(BIN)

run:
		$(GOPATH)/$(BIN)

deps:
		go get -t ./...

fmt:
		go fmt ./...

lint:
		golint ./...

clean:
		rm $(GOPATH)/$(BIN)

migrate:
		go run migrations/main.go

docker:
		sudo docker-compose up -d

dredd:
		dredd $(DREDD_API) $(DREDD_LOCAL)