.PHONY: build
build:
	docker buildx build \
		--platform linux/amd64 \
		--load \
		--tag axellofberg/my-scratch-image:latest \
		.

.PHONY: push-dockerhub
push-dockerhub:
	docker tag my-scratch-image axellofberg/my-scratch-image # defaults to latest
	docker push axellofberg/my-scratch-image

	docker tag my-scratch-image axellofberg/my-scratch-image:abc-123
	docker push axellofberg/my-scratch-image:abc-123

# Have to authenticate to ghcr.io first
# https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-to-the-container-registry
.PHONY: push-github-packages
push-github-packages:
	docker tag my-scratch-image ghcr.io/axellofberg/my-scratch-image # defaults to latest
	docker push ghcr.io/axellofberg/my-scratch-image

	docker tag my-scratch-image ghcr.io/axellofberg/my-scratch-image:abc-123 # defaults to latest
	docker push ghcr.io/axellofberg/my-scratch-image:abc-123