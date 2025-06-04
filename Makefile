API_DIR=api

USER_API=$(API_DIR)/user/openapi.yaml
USER_OUT=$(API_DIR)/user/user.gen.go

PRODUCT_API=$(API_DIR)/product/openapi.yaml
PRODUCT_OUT=$(API_DIR)/product/product.gen.go

.PHONY: all openapi-user openapi-product sqlc clean

all: sqlc openapi-user openapi-product

sqlc:
	sqlc generate

openapi-user:
	oapi-codegen -generate types,fiber -package api -o $(USER_OUT) $(USER_API)

openapi-product:
	oapi-codegen -generate types,fiber -package api -o $(PRODUCT_OUT) $(PRODUCT_API)

clean:
	rm -f $(API_DIR)/user/*.gen.go $(API_DIR)/product/*.gen.go
