.PHONY: vendor
vendor:
	GO111MODULE=on GOSUMDB=off GOPROXY=direct go mod vendor

.PHONY: internalvendor
internalvendor:
	cd toiletpaper && GO111MODULE=on GOSUMDB=off GOPROXY=direct go mod vendor