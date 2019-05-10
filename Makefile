DOCKER_DIR = docker

.PHONY: dc-%
dc-%:
	$(DOCKER_DIR)/$*.sh

.PHONY: up-build
up-build:
	$(DOCKER_DIR)/up.sh --build

.PHONY: up
up: dc-up

.PHONY: down
down:
	$(DOCKER_DIR)/down.sh -v --rmi all --remove-orphans

.PHONY: start
start: dc-start

.PHONY: stop
stop: dc-stop

test:
	./docker/run.sh core go test -timeout 30s -v ./...

test_cover:
	./docker/run.sh core go test -timeout 30s -v -cover ./...
