.PHONY: all
all: load generate

.PHONY: load
load:
	curl -O https://api.jellyfin.org/openapi/jellyfin-openapi-stable.json

.PHONY: generate
generate: clear
	openapi-generator-cli version-manager set latest
	openapi-generator-cli generate -c config.yaml

.PHONY: clear
clear:
	mv ./api/configuration.go .
	rm -r ./api/*
	mv ./configuration.go ./api/
