.PHONY: build push run init start-dependencies

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
	@./run.sh

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

start-dependencies:
	@echo "Starting Weaviate dependency..."
	@docker run -d \
	  --name weaviate \
	  -p 8888:8080 \
	  --restart on-failure:0 \
	  -e QUERY_DEFAULTS_LIMIT=25 \
	  -e AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true \
	  -e PERSISTENCE_DATA_PATH='/var/lib/weaviate' \
	  -e DEFAULT_VECTORIZER_MODULE='none' \
	  -e ENABLE_MODULES='' \
	  -e AUTOSCHEMA_ENABLED=true \
	  -e CLUSTER_HOSTNAME='node1' \
	  semitechnologies/weaviate:1.18.1 \
	  --host 0.0.0.0 \
	  --port 8080 \
	  --scheme http