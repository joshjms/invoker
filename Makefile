WORKER_IMAGE_TAG?=invoker-worker

.PHONY: build
build:
	go build -o bin/invoker main.go

.PHONY: build-and-push-worker-image
build-and-push-worker-image: build-worker-image push-worker-image

.PHONY: build-worker-image
build-worker-image:
	docker build -t $(WORKER_IMAGE_TAG) -f ./worker/Dockerfile .

.PHONY: push-worker-image
push-worker-image:
	docker push $(WORKER_IMAGE_TAG)
