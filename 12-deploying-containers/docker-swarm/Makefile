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
VM_IP?=212.2.244.152
DOCKER_HOST:="ssh://ubuntu@${VM_IP}"

.PHONY: swarm-init
swarm-init:
	DOCKER_HOST=${DOCKER_HOST} docker swarm init

.PHONY: swarm-deploy-stack
swarm-deploy-stack:
	DOCKER_HOST=${DOCKER_HOST} docker stack deploy -c docker-swarm.yml example-app

.PHONY: swarm-ls
swarm-ls:
	DOCKER_HOST=${DOCKER_HOST} docker service ls

.PHONY: swarm-remove-stack
swarm-remove-stack:
	DOCKER_HOST=${DOCKER_HOST} docker stack rm example-app

.PHONY: create-secrets
create-secrets:
	printf "foobarbaz" | DOCKER_HOST=${DOCKER_HOST} docker secret create postgres-passwd -
	printf "postgres://postgres:foobarbaz@db:5432/postgres" | DOCKER_HOST=${DOCKER_HOST} docker secret create database-url -

.PHONY: delete-secrets
delete-secrets:
	DOCKER_HOST=${DOCKER_HOST} docker secret rm postgres-passwd database-url

.PHONY: redeploy-all
redeploy-all:
	-$(MAKE) swarm-remove-stack
	-$(MAKE) delete-secrets
	@sleep 3
	-$(MAKE) create-secrets
	-$(MAKE) swarm-deploy-stack
