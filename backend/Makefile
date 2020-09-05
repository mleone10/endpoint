.PHONY: build clean run

clean:
	rm -rf ./bin

build: clean
	for CMD in `ls cmd`; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$CMD ./cmd/$$CMD/...; \
	done
	cp endpointCredentials.json bin/

run: clean build
	./bin/endpoint

deploy:
	serverless deploy --verbose
