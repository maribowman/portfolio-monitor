NAME			:=marschine/signal-transmitter
GIT_BRANCH		:= $(shell git rev-parse --abbrev-ref HEAD)
GIT_HASH		:= $(shell git rev-parse --short HEAD)
TAG				:= $(GIT_BRANCH)_$(GIT_HASH)
IMAGE			:= $(NAME):$(TAG)

### docker
build:
	docker build -t $(IMAGE) .
	docker image prune -f --filter label=stage=builder

run: build
	@docker run -d -p 9876:9876 --name $(TAG) $(IMAGE) .

### testing
tests:
	@go test ./...

coverage:
	@go test -cover ./...

smoke: build
	@docker run -d --rm -p 9876:9876 --name test-runner $(IMAGE) .
	@bash ./test/smoke.sh
	@docker stop test-runner
	@docker rmi $(IMAGE)

# TODO create dockerfiles for ingestion and faust
### docker
build:
	docker build -t $(IMAGE) .
	docker image prune -f --filter label=stage=builder

run: build
	@docker run -d -p 8080:8080 --name $(TAG) $(IMAGE) .

