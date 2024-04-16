.PHONY: all
all: load generate

.PHONY: load
load:
	curl -O https://api.jellyfin.org/openapi/jellyfin-openapi-stable.json

.PHONY: generate
generate:
	openapi-generator generate \
		-i jellyfin-openapi-stable.json \
		-g go \
		-o api \
		--git-repo-id jellyfin-go \
		--git-user-id sj14 \
		--additional-properties=enumClassPrefix=true,generateInterfaces=true,packageName=api,withGoMod=false,isGoSubmodule=true
