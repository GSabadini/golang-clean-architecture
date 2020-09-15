up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

logs:
	docker logs -f go-challenge

test:
	go test ./...