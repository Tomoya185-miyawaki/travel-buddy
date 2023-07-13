DC := docker-compose
cddk := cd docker

.PHONY: up
up:
	$(cddk) && $(DC) up

.PHONY: upd
upd:
	$(cddk) && $(DC) up -d

.PHONY: build
build:
	$(cddk) && $(DC) build

.PHONY: down
down:
	$(cddk) && $(DC) down

.PHONY: api
api:
	$(cddk) && $(DC) exec api sh

.PHONY: migrate
migrate:
	$(cddk) && $(DC) exec api sh -c 'go run db/migrate/migrate.go'

.PHONY: db
db:
	$(cddk) && $(DC) exec db bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'


