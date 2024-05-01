.PHONY: all
all: load generate

.PHONY: load
load:
	curl -O https://api.jellyfin.org/openapi/jellyfin-openapi-stable.json

.PHONY: generate
generate:
	openapi-generator generate -c config.yaml
