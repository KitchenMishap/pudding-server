.PHONY: all windows linux-amd64 linux-arm64

all: windows linux-amd64 linux-arm64

windows:
	GOOS=windows GOARCH=amd64 go build -o pudding-server.exe .

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o pudding-server-linux-amd64 .

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o pudding-server-linux-arm64 .
