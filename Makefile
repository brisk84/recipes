OUTPUT_BUILD:=./cmd/main
GO_BUILD_FILE:=./cmd/main.go

deps:
	go mod tidy -v
	go mod download -x

.PHONY: build
build: deps
	go build -v -o ${OUTPUT_BUILD} ${GO_BUILD_FILE}

generate:
	cd api; oapi-codegen --package=api --generate types,gorilla -o recipes-srv.gen.go recipes.yml; cd ..
	go generate ./...

docker_build:
	docker-compose -f docker-compose.yml build

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down --remove-orphans

lint:
	golangci-lint run

run:
	export PG_URI="postgresql://localhost/recipes?user=postgres&password=sqlRec1pe58&sslmode=disable"
	go run ./cmd/main.go
