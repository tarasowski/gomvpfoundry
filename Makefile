build:
	GOOS=linux GOARCH=amd64 go build -o ./tmp/server ./cmd/server/main.go
	scp ./public/index.html mvpfoundry:/home/ubuntu/public/index.html
	scp ./tmp/server mvpfoundry:/home/ubuntu/
