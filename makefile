.PHONY: $(MAKECMDGOALS)

validate-api:
	npx swagger-cli validate api/openapi/spec.yaml