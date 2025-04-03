### Build + Push Images
.PHONY: build-push
build-push:
	cd ../06-building-container-images/client-react && N=5 $(MAKE) build-N && N=5 $(MAKE) push-N
	cd ../06-building-container-images/api-node && N=9 $(MAKE) build-N && N=9 $(MAKE) push-N
	cd ../06-building-container-images/api-golang && N=8 $(MAKE) build-N && N=8 $(MAKE) push-N