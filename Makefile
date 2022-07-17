export CGO_ENABLED = 1
export NEXT_TELEMETRY_DISABLED = 1

.PHONY: run
run:
	go run ./cmd/borzoi/

.PHONY: run-web
run-web:
	cd web && npm run dev

.PHONY: build
build: build-web
	go build ./cmd/borzoi

.PHONY: build-web
build-web:
	cd web && \
	npm install --frozen-lockfile && \
	npm run export && \
    mv dist ../cmd/borzoi/web

.PHONY: clean
clean:
	rm -f borzoi
	rm -rf ./cmd/borzoi/web
	rm -rf ./web/dist
	rm -rf ./web/.next

.PHONY: docker
docker:
	docker build -t borzoi:local . && \
	docker run --rm -p 80:8080 borzoi:local
