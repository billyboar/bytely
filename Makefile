dev:
	go run cmd/api/main.go

proto:
	protoc -I=./pb --go_out=./pb --go-grpc_out=./pb ./pb/bytely.proto

generate-swagger:
	cd cmd/client && swag init --parseDependency

# build-api: generate-swagger
build-api: 
	GOOS=linux GOARCH=amd64 go build -o cmd/api/bytely_api cmd/api/main.go
build-client:
	GOOS=linux GOARCH=amd64 go build -o cmd/client/bytely_client cmd/client/main.go
build-bins: build-api build-client

docker-build-api: build-api
	docker build -t bytely_api:latest cmd/api/.
docker-build-client: build-client
	docker build -t bytely_client:latest cmd/client/.
docker-build: docker-build-api docker-build-client

up: docker-build
	docker-compose up

.PHONY: dev proto up build-api build-client build-bins docker-build-api
	 docker-build-client docker-build