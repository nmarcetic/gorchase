# Copyright (c) nmarcetic

DOCKER_IMAGE_NAME_PREFIX ?= gorchase
BUILD_DIR = build
SERVICES = server
DOCKERS = $(addprefix docker_,$(SERVICES))
DOCKERS_DEV = $(addprefix docker_dev_,$(SERVICES))
CGO_ENABLED ?= 0
GOARCH ?= amd64

define compile_service
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -mod=vendor -ldflags "-s -w" -o ${BUILD_DIR}/gorchase-$(1) cmd/$(1)/main.go
endef

define make_docker
	$(eval svc=$(subst docker_,,$(1)))

	docker build \
		--no-cache \
		--build-arg SVC=$(svc) \
		--build-arg GOARCH=$(GOARCH) \
		--build-arg GOARM=$(GOARM) \
		--tag=$(DOCKER_IMAGE_NAME_PREFIX)/$(svc) \
		-f docker/Dockerfile .
endef

define make_docker_ui
	docker build \
		--no-cache \
		--tag=$(DOCKER_IMAGE_NAME_PREFIX)/ui \
		-f ui/docker/Dockerfile ui
endef

define make_docker_dev
	$(eval svc=$(subst docker_dev_,,$(1)))

	docker build \
		--no-cache \
		--build-arg SVC=$(svc) \
		--tag=$(DOCKER_IMAGE_NAME_PREFIX)/$(svc) \
		-f docker/Dockerfile.dev ./build
endef

build: $(SERVICES)

.PHONY: all $(SERVICES) dockers dockers_dev latest release

clean:
	rm -rf ${BUILD_DIR}

cleandocker:
	# Stops containers and removes containers, networks, volumes, and images created by up
	docker-compose -f docker/docker-compose.yml down --rmi all -v --remove-orphans

ifdef pv
	# Remove unused volumes
	docker volume ls -f name=$(DOCKER_IMAGE_NAME_PREFIX) -f dangling=true -q | xargs -r docker volume rm
endif

install:
	cp ${BUILD_DIR}/* $(GOBIN)

test:
	go test -mod=vendor -v -race -count 1 -tags test $(shell go list ./... | grep -v 'vendor\|cmd')

$(SERVICES):
	$(call compile_service,$(@))

$(DOCKERS):
	$(call make_docker,$(@),$(GOARCH))
	$(call make_docker_ui)

$(DOCKERS_DEV):
	$(call make_docker_dev,$(@))

dockers: $(DOCKERS)
dockers_dev: $(DOCKERS_DEV)

define docker_push
	for svc in $(SERVICES); do \
		docker push $(DOCKER_IMAGE_NAME_PREFIX)/$$svc:$(1); \
	done
endef

changelog:
	git log $(shell git describe --tags --abbrev=0)..HEAD --pretty=format:"- %s"

latest: dockers
	$(call docker_push,latest)

run:
	docker-compose -f docker/docker-compose.yml up
