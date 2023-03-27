.PHONY: deploy-traefik
deploy-traefik:
	-helm repo add traefik https://traefik.github.io/charts
	helm upgrade --install -n traefik --create-namespace traefik traefik/traefik --version 20.8.0

.PHONY: deploy-postgres
deploy-postgres:
	-helm repo add bitnami https://charts.bitnami.com/bitnami
	helm upgrade --install db bitnami/postgresql --set auth.postgresPassword=foobarbaz --version 12.1.15

.PHONY: deploy-app
deploy-app:
	kubectl apply -f .
