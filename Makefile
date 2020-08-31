.PHONY: build clean deploy

build:
	for CMD in `ls cmd`; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$CMD ./cmd/$$CMD/...; \
	done

clean:
	rm -rf ./bin