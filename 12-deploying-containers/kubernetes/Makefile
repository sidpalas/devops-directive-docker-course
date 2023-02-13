.PHONY: deploy-traefik
deploy-traefik:
	-helm repo add traefik https://traefik.github.io/charts
	helm upgrade --install -n traefik --create-namespace -f traefik-values traefik traefik/traefik

.PHONY: deploy-postgres
deploy-postgres:
	-helm repo add bitnami https://charts.bitnami.com/bitnami
	helm upgrade --install db bitnami/postgresql --set auth.postgresPassword=foobarbaz

.PHONY: deploy-app
deploy-app:
	kubectl apply -f .
