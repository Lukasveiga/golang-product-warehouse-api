.PHONY: welcome
welcome:
	@echo "**********************************************************"
	@echo "*************  Welcome to my first api in go  ************"
	@echo "**********************************************************"

.PHONY: test
test:
	go test ./... -v

test-cover:
	go test -cover ./...

run-api:
	go run cmd/api/main.go