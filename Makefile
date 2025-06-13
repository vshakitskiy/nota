compose.up:
	docker compose --env-file .env \
	-f deployments/docker-compose.yaml \
	up -d

compose.down:
	docker compose --env-file .env \
	-f deployments/docker-compose.yaml \
	down

tidy:
	cd nota.auth && go mod tidy && cd .. && \
	cd nota.gateway && go mod tidy && cd .. && \
	cd nota.shared && go mod tidy && cd .. && \
	cd nota.snippet && go mod tidy && cd ..

fmt:
	cd nota.auth && go fmt ./... && cd .. && \
	cd nota.gateway && go fmt ./... && cd .. && \
	cd nota.shared && go fmt ./... && cd .. && \
	cd nota.snippet && go fmt ./... && cd ..

auth.start:
	cd nota.auth && go run cmd/main.go

gateway.start:
	cd nota.gateway && go run cmd/main.go

snippet.start:
	cd nota.snippet && go run cmd/main.go