API_NODE_PATH:=../../05-example-web-application/client-react/

N?=0
.PHONY: build-N
build-N:
	docker build --file ./Dockerfile.$(N) -t client-react:$(N) ${API_NODE_PATH}

.PHONY: build-all
build-all:
	for number in 0 1 2 3 4 5; do \
			N=$$number $(MAKE) build-N; \
	done