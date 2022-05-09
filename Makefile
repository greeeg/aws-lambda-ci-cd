build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap cmd/main.go