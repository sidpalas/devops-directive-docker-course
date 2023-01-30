### DOCKER COMPOSE COMMANDS

.PHONY: compose-build
compose-build:
	docker compose build

.PHONY: compose-up
compose-up:
	docker compose up

.PHONY: compose-up-build
compose-up-build:
	docker compose up --build

.PHONY: compose-down
compose-down:
	docker compose down

### DOCKER CLI COMMANDS

DOCKERCONTEXT_DIR:=../05-example-web-application/
DOCKERFILE_DIR:=../06-building-container-images/

.PHONY: docker-build-all
docker-build-all:
	docker build -t client-react-vite -f ${DOCKERFILE_DIR}/client-react/Dockerfile.3 ${DOCKERCONTEXT_DIR}/client-react/

	docker build -t client-react-ngnix -f ${DOCKERFILE_DIR}/client-react/Dockerfile.5 ${DOCKERCONTEXT_DIR}/client-react/

	docker build -t api-node -f ${DOCKERFILE_DIR}/api-node/Dockerfile.7 ${DOCKERCONTEXT_DIR}/api-node/

	docker build -t api-golang -f ${DOCKERFILE_DIR}/api-golang/Dockerfile.6 ${DOCKERCONTEXT_DIR}/api-golang/

.PHONY: docker-run-all
docker-run-all:
	echo "$$DOCKER_COMPOSE_NOTE"

	# Stop and remove all running containers to avoid name conflicts
	$(MAKE) docker-stop

	$(MAKE) docker-remove

	docker run -d \
		--name db \
		-e POSTGRES_PASSWORD=foobarbaz \
		-p 5432:5432 \
		--restart unless-stopped \
		postgres:15.1-alpine

	docker run -d \
		--name api-node \
		-e DATABASE_URL=postgres://postgres:foobarbaz@db:5432/postgres \
		-p 3000:3000 \
		--restart unless-stopped \
		--link=db \
		api-node

	docker run -d \
		--name api-golang \
		-e DATABASE_URL=postgres://postgres:foobarbaz@db:5432/postgres \
		-p 8080:8080 \
		--restart unless-stopped \
		--link=db \
		api-golang

	docker run -d \
		--name client-react-vite \
		-v ${PWD}/client-react/vite.config.js:/usr/src/app/vite.config.js \
		-p 5173:5173 \
		--restart unless-stopped \
		--link=api-node \
		--link=api-golang \
		client-react-vite

	docker run -d \
		--name client-react-nginx \
		-p 5174:80 \
		--restart unless-stopped \
		--link=api-node \
		--link=api-golang \
		client-react-ngnix

docker-stop:
	-docker stop db
	-docker stop api-node
	-docker stop api-golang
	-docker stop client-react-vite
	-docker stop client-react-nginx

docker-remove:
	-docker container rm db
	-docker container rm api-node
	-docker container rm api-golang
	-docker container rm client-react-vite
	-docker container rm client-react-nginx

define DOCKER_COMPOSE_NOTE

🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨

❯ NOTE:

This command runs the example app with a bunch
of individual docker run commands. This is much
easier to manage with docker-compose (see 
docker-compose.yml and compose make targets above)

🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨


endef
export DOCKER_COMPOSE_NOTE