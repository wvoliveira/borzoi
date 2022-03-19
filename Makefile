export CGO_ENABLED = 1
export NEXT_TELEMETRY_DISABLED = 1

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
	rm -f spitz
	rm -rf ./cmd/borzoi/web
	rm -rf ./web/dist
	rm -rf ./web/.next
