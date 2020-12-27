.PHONY: build clean run

clean:
	rm -rf ./bin

test:
	go test ./internal/...

integration:
	go test --tags=integration ./internal/...

build: clean integration
	for CMD in `ls cmd`; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$CMD ./cmd/$$CMD/...; \
	done
	cp endpointCredentials.json bin/

run: build
	GOOGLE_APPLICATION_CREDENTIALS=./bin/endpointCredentials.json ./bin/endpoint

mock: build
	GOOGLE_APPLICATION_CREDENTIALS=./bin/endpointCredentials.json ENDPOINT_LOCAL=true ./bin/endpoint

deploy:
	serverless deploy --verbose
