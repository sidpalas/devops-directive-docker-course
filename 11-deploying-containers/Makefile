.PHONY: compose-up
compose-up:
	docker compose -f docker-compose-prod.yml up

# -d flag runs containers in the background:
.PHONY: compose-up-d
compose-up-d:
	docker compose -f docker-compose-prod.yml up -d

###

.PHONY: swarm-init
swarm-init:
	docker swarm init

.PHONY: swarm-deploy-stack
swarm-deploy-stack:
	docker stack deploy -c stack.yaml example-app

.PHONY: swarm-remove-stack
swarm-remove-stack:
	docker stack rm example-app

.PHONY: create-secrets
create-secrets:
	echo -n "foobarbaz" | docker secret create postgres-passwd -
	echo -n "postgres://postgres:foobarbaz@db:5432/postgres" | docker secret create database-url -