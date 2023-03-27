LDFLAGS := -s -w

all: fmt build linux darwin windows

build: linux

fmt:
	go fmt ./...

linux:
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/little_robots-linux .
	cp config/config.yaml bin/config.yaml

darwin:
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/little_robot-darwin .
	cp config/config.yaml bin/config.yaml

windows:
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/little_robot-windows.exe .
	cp config/config.yaml bin/config.yaml

clean:
	rm -f ./bin/little_robot