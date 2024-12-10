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

db-migrate:
	go run scripts/go/migrate/main.go

db-exec:
	docker exec -it mini-scan-takehome-db-1 psql -U postgres -d mini_scan

mock:
	mockgen -source=internal/db/db.go -destination=internal/db/mock/mock_db.go -package=mock_db
	mockgen -source=internal/processor/processor.go -destination=internal/processor/mock/mock_processor.go -package=mock_processor
