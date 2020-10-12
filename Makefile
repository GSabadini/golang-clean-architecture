init:
	cp .env.example .env

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

logs:
	docker logs -f go-challenge

test:
	go test ./...

fmt:
	go fmt ./...

ci:
	docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.24.0 \
	golangci-lint run
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...