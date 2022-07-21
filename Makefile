
## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run

SERVICE=sut-gateway-go

# Swagger API docs
SWAGGER_PORT=51234

proto-gen:
	protoc --proto_path=proto/ --go_out=paths=source_relative,plugins=grpc:./pb proto/*/*.proto

init:
	$(GOCMD) mod init $(SERVICE)

tidy:
	$(GOCMD) mod tidy

run:
	echo "for local development, please run: make run ENV=local"
	$(GORUN) cmd/main.go

check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger-gen:check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

swagger-serve:
	swagger serve -F=swagger ./swagger.yaml -p=$(SWAGGER_PORT) --no-open