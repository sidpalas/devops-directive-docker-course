DATABASE_URL=postgres://postgres:foobarbaz@localhost:5432/postgres

.PHONY: run-local
run-local:
	DATABASE_URL=${DATABASE_URL} npm run dev

.PHONY: run-local-debug
run-local-debug:
	DATABASE_URL=${DATABASE_URL} npm run debug

.PHONY: build-naive
build-naive:
	docker build --file ./Dockerfile.naive .
