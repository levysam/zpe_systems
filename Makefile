update_docs:
	docker exec zpe-systems-http swag init -g ./cmd/http/main.go --parseDependency --propertyStrategy pascalcase
