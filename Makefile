compose.up:
	docker compose --env-file .env \
	-f deployments/docker-compose.yaml \
	up -d

compose.down:
	docker compose --env-file .env \
	-f deployments/docker-compose.yaml \
	down

fmt:
	cd nota.auth && go fmt ./... \
	# cd ../nota.shared && go fmt ./...

auth.start:
	cd nota.auth && go run cmd/main.go

gateway.start:
	cd nota.gateway && go run cmd/main.go