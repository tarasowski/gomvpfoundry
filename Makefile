build:
	GOOS=linux GOARCH=amd64 go build -o ./tmp/server ./cmd/server/main.go
	ssh mvpfoundry 'rm -rf /home/ubuntu/public'
	scp -r ./public mvpfoundry:/home/ubuntu/public
	ssh mvpfoundry 'sudo systemctl stop mvpfoundry-go-server'
	ssh mvpfoundry 'sudo systemctl stop nginx'
	scp ./tmp/server mvpfoundry:/home/ubuntu/
	ssh mvpfoundry 'sudo systemctl daemon-reload'
	ssh mvpfoundry 'sudo systemctl restart mvpfoundry-go-server'
	ssh mvpfoundry 'sudo systemctl restart nginx'

dev:
	npx vite public
