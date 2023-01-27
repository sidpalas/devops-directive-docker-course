.PHONY: run-local
run-local:
	DATABASE_URL=postgres://postgres:foobarbaz@localhost:5432/postgres \
    node ./src/index.js

.PHONY: build-naive
build-naive:
	docker build --file ./Dockerfile.naive .
