BIN_FOLDER = bin
SERVER = rpc-runtime
GO_CMD_RUNTIME = main.go

IMAGE_NAME_SERVICE = home-services
VERSION_SERVICE = 1.0.0

HOST = 127.0.0.1
PORT = 5432
DATABASE = home_services
USER = root
PASSWORD = secret

create-db:
	docker run --name $(IMAGE_NAME_SERVICE) -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
.PHONY: create-db

init-db:
	docker exec -it $(IMAGE_NAME_SERVICE)  createdb --username=root --owner=root $(DATABASE)
.PHONY: init-db

migrate-up:
		migrate -path migrations -database "postgresql://$(USER):$(PASSWORD)@$(HOST):$(PORT)/$(DATABASE)?sslmode=disable" -verbose up
.PHONY: migration-up

migrate-down:
		migrate -path migrations -database "postgresql://$(USER):$(PASSWORD)@$(HOST):$(PORT)/$(DATABASE)?sslmode=disable" -verbose down
.PHONY: migration-down

proto:
	rm -f pb/*.go
	rm -f docs/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs --openapiv2_opt=allow_merge=true,merge_file_name=home_services \
	proto/*.proto
	statik -src=./docs -dest=./docs/ui