build:
	GOOS=linux GOARCH=amd64 go build -o ./tmp/server ./cmd/server/main.go
	scp -r ./public mvpfoundry:/home/ubuntu/public
	scp ./tmp/server mvpfoundry:/home/ubuntu/

dev:
	npx vite public
