### DOCKER COMPOSE

.PHONY: compose-up
compose-up:
	docker compose -f docker-compose-prod.yml up

# -d flag runs containers in the background:
.PHONY: compose-up-d
compose-up-d:
	docker compose -f docker-compose-prod.yml up -d

.PHONY: compose-down
compose-down:
	docker compose -f docker-compose-prod.yml down

### DOCKER SWARM

DOCKER_HOST:="ssh://ubuntu@212.2.244.220"

.PHONY: swarm-init
swarm-init:
	DOCKER_HOST=${DOCKER_HOST} docker swarm init

.PHONY: swarm-deploy-stack
swarm-deploy-stack:
	DOCKER_HOST=${DOCKER_HOST} docker stack deploy -c docker-swarm.yml example-app

.PHONY: swarm-remove-stack
swarm-remove-stack:
	DOCKER_HOST=${DOCKER_HOST} docker stack rm example-app

.PHONY: create-secrets
create-secrets:
	echo -n "foobarbaz" | DOCKER_HOST=${DOCKER_HOST} docker secret create postgres-passwd -
	echo -n "postgres://postgres:foobarbaz@db:5432/postgres" | DOCKER_HOST=${DOCKER_HOST} docker secret create database-url -