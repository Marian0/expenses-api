.PHONY: setup expenses-api test

all: expenses-api
# all: expenses-api go-swagger

# setup:
# 	yarn
# 	make expenses-api

devinit: expenses-api.go
	GO111MODULE=on go mod init gitlab.com/marian0/expenses-api
	go mod tidy

expenses-api: expenses-api.go
	GO111MODULE=on go build expenses-api.go

# swaggerlocal: go-swagger redoc

# go-swagger:
# 	go build -o ./swagger -i github.com/go-swagger/go-swagger/cmd/swagger
# 	./swagger generate spec -o ./docs/swagger.json


test:
	go test -cover -count=1
