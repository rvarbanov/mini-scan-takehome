test:
	go test ./...

.PHONY: build
build: build-scanner build-consumer

build-scanner:
	DOCKER_BUILDKIT=1 docker build --platform linux/arm64 -t scanner -f cmd/scanner/Dockerfile .

build-consumer:
	DOCKER_BUILDKIT=1 docker build --platform linux/arm64 -t consumer -f cmd/consumer/Dockerfile .

run:
	docker-compose up --build

kill: 
	docker kill $(docker ps -q)

clean:
	go clean
	go fmt ./...
	go mod tidy

docker-clean:
	docker system prune -f -a --volumes
	docker rmi $(docker images -a -q)
	docker rm $(docker ps -a -f status=exited -q)
