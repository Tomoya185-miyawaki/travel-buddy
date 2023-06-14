DC := @docker-compose

.PHONY: up
up:
	$(DC) up

.PHONY: upd
upd:
	$(DC) up -d

.PHONY: build
build:
	$(DC) build

.PHONY: down
down:
	$(DC) down

.PHONY: api
api:
	$(DC) exec api sh