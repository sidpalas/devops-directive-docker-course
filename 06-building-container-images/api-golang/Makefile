API_GOLANG_PATH:=../../05-example-web-application/api-golang/

N?=0
.PHONY: build-N
build-N:
	docker build --file ./Dockerfile.$(N) -t api-golang:$(N) ${API_GOLANG_PATH}

.PHONY: build-all
build-all:
	for number in 0 1 2 3 4 5 6 ; do \
			N=$$number $(MAKE) build-N; \
	done