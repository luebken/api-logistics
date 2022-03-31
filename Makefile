init:
	go mod init github.com/luebken/api-logistics
	go mod tidy

build:
	go build -o bin/api ./cmd/api

run:
	go run cmd/api/main.go

docker-build:
	docker build . -f Dockerfile -t luebken/api-logistics:latest

docker-run:
	docker run luebken/api-logistics:latest

docker-push:
	docker push luebken/api-logistics:latest

test:
	curl localhost:8090/vessels?id=1001 | jq .
	curl localhost:8090/vessels | jq .
	curl localhost:8090/containers?id=2001 | jq .
	curl localhost:8090/containers | jq .
	curl -X POST -s localhost:8090/vessels -d \ '{"id": 1004, "description":"New Boat", "lat": 155.703461, "lng": 55.703461}'
	curl localhost:8090/vessels | jq .
