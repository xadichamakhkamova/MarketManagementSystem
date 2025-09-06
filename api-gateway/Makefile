PROTO_DIR=protos

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/product-service/product.proto \
		$(PROTO_DIR)/debt-service/debt.proto \
		$(PROTO_DIR)/dashboard-service/dashboard.proto 
		
		
SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/https/api.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)
