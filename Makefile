.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')


APP             = server
SERVER_BIN  	= ./cmd/${APP}/${APP}
RELEASE_ROOT 	= release
RELEASE_SERVER 	= release/${APP}


all: start

build:
	@go build -ldflags "-w -s" -o $(SERVER_BIN) ./cmd/${APP}/main.go

start:
	@go run ./cmd/${APP}/main.go

test:
	cd ./test && go test -v

clean:
	rm -rf $(SERVER_BIN)

docker:
	docker build -t ${APP}:v1 .

