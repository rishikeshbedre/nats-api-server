all:
	env CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o nats-api-server .
