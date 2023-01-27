.PHONY: run-postgres
run-postgres:
	@echo Starting postgres container
	-docker run \
		-e POSTGRES_PASSWORD=foobarbaz \
		-p 5432:5432 \
		postgres:15.1-alpine

.PHONY: run-api-node
run-api-node:
	@echo Starting node api
	cd api-node && \
		PGUSER=postgres \
    PGHOST=localhost \
    PGPASSWORD=foobarbaz \
    PGDATABASE=postgres \
    PGPORT=5432 \
		npm run dev

.PHONY: run-api-golang
run-api-golang:
	@echo Starting golang api
	cd api-golang && \
		DATABASE_URL=postgres://postgres:foobarbaz@localhost:5432/postgres \
		go run main.go

.PHONY: run-client-react
run-client-react:
	@echo Starting react client
	cd client-react && \
		npm run dev
