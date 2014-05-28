build: *.go
	go build -o build/docker-testing

crosscompile: *.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/linux/docker-testing
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o build/mac/docker-testing
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/windows/docker-testing

all: clean build crosscompile

.PHONY: clean

clean:
	rm -rf build
