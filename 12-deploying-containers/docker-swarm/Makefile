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

CIVO_SSH:="ssh://ubuntu@212.2.244.220"

.PHONY: swarm-init
swarm-init:
	DOCKER_HOST=${CIVO_SSH} docker swarm init

.PHONY: swarm-deploy-stack
swarm-deploy-stack:
	DOCKER_HOST=${CIVO_SSH} docker stack deploy -c docker-swarm.yml example-app

.PHONY: swarm-remove-stack
swarm-remove-stack:
	DOCKER_HOST=${CIVO_SSH} docker stack rm example-app

.PHONY: create-secrets
create-secrets:
	echo -n "foobarbaz" | DOCKER_HOST=${CIVO_SSH} docker secret create postgres-passwd -
	echo -n "postgres://postgres:foobarbaz@db:5432/postgres" | DOCKER_HOST=${CIVO_SSH} docker secret create database-url -