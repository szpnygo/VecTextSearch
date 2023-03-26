.PHONY: build push run init

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

run:
	@echo "Running application locally..."
	export $(grep -v '^#' .env | xargs) && go run ./cmd/main.go

init:
	@echo "Creating .env template..."
	@if [ ! -f .env ]; then \
		echo "VECTEXTSEARCH_OPENAI_KEY=your_openai_api_key_here" >> .env; \
		echo "VECTEXTSEARCH_API_PORT=8000" >> .env; \
		echo "VECTEXTSEARCH_WEAVIATE_URL=localhost:8888" >> .env; \
		echo ".env template created successfully."; \
	else \
		echo ".env file already exists. No changes were made."; \
	fi
