compose.up:
	docker compose --env-file .env \
	-f deployments/postgres/docker-compose.yaml \
	-p nota_postgres up -d

compose.down:
	docker compose --env-file .env \
	-f deployments/postgres/docker-compose.yaml \
	-p nota_postgres down

fmt:
	cd nota.auth && go fmt ./... \
	# cd ../nota.shared && go fmt ./...