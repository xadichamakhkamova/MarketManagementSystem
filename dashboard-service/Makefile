# .env fayldan DB_URL o'qish
ifneq (,$(wildcard .env))
	include .env
	export
endif

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		proto/dashboard-service.proto

mig-create:
	@migrate create -ext sql -dir migrations -seq create_dashboarddb_table

migrate-up:
	@migrate -path ./migrations -database $(DB_URL) up

migrate-down:
	@migrate -path ./migrations -database $(DB_URL) down

migrate-force:
	@migrate -path ./migrations -database $(DB_URL) verbose -force 1
