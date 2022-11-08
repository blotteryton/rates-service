.PHONY: build

build:
	go build cmd/main.go

.PHONY: run
run:
	go run cmd/main.go

.PHONY: docker-build
docker-build:
	DOCKER_SCAN_SUGGEST=false docker build -t cmc_fetcher .