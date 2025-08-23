APP_SERVICE_NAME=api
DB_SERVICE_NAME=db
APP_CONTAINER_NAME=go_clean_api
APP_BINARY_PATH=./go-clean-api

.PHONY: init up down logs migrate seed

# Inicializa tudo
init:
	cp -n .env-example .env || true
	docker compose up -d --build

	@echo "Aguardando banco ficar pronto..."
	docker compose exec $(DB_SERVICE_NAME) sh -c 'until pg_isready -h $(DB_SERVICE_NAME) -p 5432 -U $$DB_USER -d $$DB_NAME; do sleep 1; done'

	@echo "Aguardando container $(APP_CONTAINER_NAME) estar rodando..."
	@while [ -z "$$(docker compose ps -q $(APP_SERVICE_NAME))" ] || \
		[ "$$(docker inspect -f '{{.State.Running}}' $$(docker compose ps -q $(APP_SERVICE_NAME)))" != "true" ]; do \
		echo "Aguardando container..."; sleep 2; \
	done

	@echo "Aguardando app iniciar (sleep 10s dentro do container)..."
	docker compose exec $(APP_SERVICE_NAME) sh -c 'sleep 10'

	@echo "Rodando migrate..."
	docker compose run --rm api ./go-clean-api-migrate

up:
	docker compose up -d --build

down:
	docker compose down --volumes

logs:
	docker compose logs -f $(APP_SERVICE_NAME)

migrate:
	docker compose exec $(APP_SERVICE_NAME) sh -c '$(APP_BINARY_PATH) migrate'
