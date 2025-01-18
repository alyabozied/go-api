include .env
MIGRATIONS_PATH= ./model/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))


.PHONY: migrate-up
migration-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(database) up
	

.PHONY: migrate-down
migration-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(database) down $(filter-out $@,$(MAKECMDGOALS))
	