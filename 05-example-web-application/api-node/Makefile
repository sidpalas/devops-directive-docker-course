.PHONY: run-local
run-local:
	PGUSER=postgres \
    PGHOST=localhost \
    PGPASSWORD=foobarbaz \
    PGDATABASE=postgres \
    PGPORT=5432 \
    node ./src/index.js

.PHONY: build-naive
build-naive:
	docker build --file ./Dockerfile.naive .
