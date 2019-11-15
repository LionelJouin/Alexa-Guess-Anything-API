
.PHONY: all
all: default

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o dist/main ./cmd

.PHONY: deploy
deploy:
	echo ""

.PHONY: default
default: build