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

.PHONY: build-sample
build-sample:
	DOCKER_BUILDKIT=1 docker build \
		--build-arg BASE_IMAGE_TAG=19.3 \
		--secret id=secret.txt,src=local-secret.txt \
		-t sample \
		-f Dockerfile.sample \
		.

define ENTRYPOINT_CMD_DESCRIPTION
##############################

See Dockerfile.sample for image definition.

This series of docker run commands is meant
to help you understand the interaction between
CMD and ENTRYPOINT.

The image has the following:

 ENTRYPOINT [ "echo", "Hey Team 👋 (entrypoint)" ]
 CMD [ "+ (cmd)" ]

##############################
endef
export ENTRYPOINT_CMD_DESCRIPTION

.PHONY: run-sample-entrypoint-cmd
run-sample-entrypoint-cmd: build-sample
	@echo "$$ENTRYPOINT_CMD_DESCRIPTION"

	@echo "Run with no arguments:"
	docker run sample

	@echo "##############################"
	@echo "Run with argument (CMD is ignored):"
	docker run sample "+ (argument)"

	@echo "##############################"
	@echo "Overriden entrypoint with no arguments (CMD is ignored):"
	docker run --entrypoint echo sample

	@echo "##############################"
	@echo "Overriden entrypoint with arguments (CMD is ignored):"
	docker run --entrypoint echo sample "Hey Team 👋 (Overriden entrypoint + arguments)"