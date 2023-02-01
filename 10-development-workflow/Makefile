DEV_COMPOSE_FILE=docker-compose-dev.yml
DEBUG_COMPOSE_FILE=docker-compose-debug.yml

### DOCKER COMPOSE COMMANDS

.PHONY: compose-build
compose-build:
	docker compose -f $(DEV_COMPOSE_FILE) build

.PHONY: compose-up
compose-up:
	docker compose -f $(DEV_COMPOSE_FILE) up

.PHONY: compose-up-build
compose-up-build:
	docker compose -f $(DEV_COMPOSE_FILE) up --build

.PHONY: compose-up-debug-build
compose-up-debug-build:
	docker compose -f $(DEV_COMPOSE_FILE) -f $(DEBUG_COMPOSE_FILE) up --build

.PHONY: compose-down
compose-down:
	docker compose -f $(DEV_COMPOSE_FILE) down

###

DOCKERCONTEXT_DIR:=../05-example-web-application/
DOCKERFILE_DIR:=../10-development-workflow/

.PHONY: docker-build-all
docker-build-all:
	docker build -t api-node -f ${DOCKERFILE_DIR}/api-node/Dockerfile.dev ${DOCKERCONTEXT_DIR}/api-node/
	docker build -t api-golang -f ${DOCKERFILE_DIR}/api-golang/Dockerfile.dev ${DOCKERCONTEXT_DIR}/api-golang/

.PHONY: run-tests
run-tests:
	docker run -t api-golang go test -v ./...
	docker run -it api-node npm run test