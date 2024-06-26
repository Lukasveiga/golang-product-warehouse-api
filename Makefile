ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: welcome
welcome:
	@echo "**********************************************************"
	@echo "*************  Welcome to my first api in go  ************"
	@echo "**********************************************************"

GOOSE=goose
DB_URL=postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

.PHONY: test
test:
	go test ./... -v

test-cover:
	go test -cover ./...

run-api:
	go run cmd/api/main.go

migrate-status:
	$(GOOSE) -dir ./migrations postgres "$(DB_URL)" status

migrate-reset:
	$(GOOSE) -dir ./migrations postgres "$(DB_URL)" reset

migrate-up:
	$(GOOSE) -dir ./migrations postgres "$(DB_URL)" up

migrate-down:
	$(GOOSE) -dir ./migrations postgres "$(DB_URL)" down

migrate-create:
	@read -p "Enter migration name: " name; \
	$(GOOSE) -dir ./migrations create $$name sql