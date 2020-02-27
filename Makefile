all:
	env CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o nats-api-server .

test:
	cp -r configuration apis/
	env GO111MODULE=on go test -race -coverprofile=coverage.txt -covermode=atomic github.com/rishikeshbedre/nats-api-server/...
