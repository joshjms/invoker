.PHONY: build-worker-image
build-worker-image:
	docker build -t invoker-worker -f ./worker/Dockerfile .
