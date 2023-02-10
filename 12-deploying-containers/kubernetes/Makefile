.PHONY: deploy-postgres
deploy-postgres:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm install db bitnami/postgresql --set auth.postgresPassword=foobarbaz

.PHONY: deploy-app
deploy-app:
	kubectl apply -f .
