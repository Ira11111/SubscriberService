.PHONY: $(MAKECMDGOALS)

install-api-tool:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.5.0

install-migration-tool:
	go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

install-tools: install-api-tool install-migration-tool

validate-api:
	npx swagger-cli validate api/openapi/spec.yaml

gen-chi-server:
	mkdir -p ./api/generated
	bash -c 'oapi-codegen -generate chi-server -package generated <(npx swagger-cli bundle api/openapi/spec.yaml --type yaml) > api/generated/server.gen.go'

gen-types:
	mkdir -p ./api/generated
	bash -c 'oapi-codegen -generate types -package generated <(npx swagger-cli bundle api/openapi/spec.yaml --type yaml) > api/generated/types.gen.go'

migrate-up:
	goose up
migrate-down:
	goose down-to 00001
