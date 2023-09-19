serve:
	go run main.go

docker-dev:
	docker-compose -f docker-compose-dev.yaml up -d

docker-dev-down:
	docker-compose -f docker-compose-dev.yaml down

docker-clean:
	docker-compose rm -fv postgres

deps:
	@go mod download
	@go mod verify