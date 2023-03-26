.PHONY: build push

IMAGE_NAME := neosu/vec-text-search
IMAGE_TAG := latest

build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

push: build
	@echo "Logging into Docker Hub..."
	docker login
	@echo "Pushing Docker image to Docker Hub..."
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
