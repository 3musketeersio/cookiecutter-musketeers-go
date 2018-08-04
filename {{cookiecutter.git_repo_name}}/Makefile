# DOCKER_RUN_ALPINE has been added for creating .env file because
#		- trying to run a docker-compose command will fail if .env is not there
#				as service musketeers requires it
#		- when TravisCI generates this example with cookiecutter, the file owner stays root
#				and running `cp -f $(ENVFILE) .env` will result a permission denied error
DOCKER_RUN_ALPINE = docker run -v $(PWD):/opt/app -w /opt/app alpine
APP_NAME={{ cookiecutter.git_repo_name }}
VERSION = 0.0.1
IMAGE_NAME ?= {{ cookiecutter.git_username }}/$(APP_NAME):$(VERSION)
BIN_DIR = bin
BIN_ARTIFACT = bin.zip
EXECUTABLE = $(BIN_DIR)/$(APP_NAME)
GOLANG_DEPS_DIR = vendor
GOLANG_DEPS_ARTIFACT = golang_vendor.zip
PROFILE_NAME ?= profile.out
COMPOSE_RUN_GOLANG = docker-compose run --rm golang
COMPOSE_RUN_MUSKETEERS = docker-compose run --rm musketeers
# ENVFILE is .env.template by default but can be overwritten
ENVFILE ?= .env.template

# all is the default Make target. it installs the dependencies, tests, and builds the application and cleans everything.
all:
	ENVFILE=.env.example $(MAKE) envfile deps test build run dockerBuild dockerTest clean
.PHONY: all

############
# env file #
############

# envfile creates or overwrites .env with $(ENVFILE)
envfile:
	$(DOCKER_RUN_ALPINE) sh -c 'cp -f $(ENVFILE) .env'
.PHONY: envfile

removeEnvfile:
	$(DOCKER_RUN_ALPINE) sh -c 'rm -f .env'
.PHONY: removeEnvfile

###################
# Artifacts #
###################

# if there is no vendor directory then unzip from golang_vendor.zip artifact
$(GOLANG_DEPS_DIR): | $(GOLANG_DEPS_ARTIFACT)
	$(COMPOSE_RUN_MUSKETEERS) make _unzipGoDeps

# if bin directory is not present, it unzips all the zip binaries into bin directory
$(BIN_DIR): | $(BIN_ARTIFACT)
	$(COMPOSE_RUN_MUSKETEERS) make _unzipBinArtifact

# _zipGoDeps zips the go dependencies so they can be passed along with a single zip file
_zipGoDeps:
	zip -rq $(GOLANG_DEPS_ARTIFACT) $(GOLANG_DEPS_DIR)/
.PHONY: _zipGoDeps

# _unzipGoDeps unzips the go dependencies zip file
_unzipGoDeps:
	unzip -qo -d . $(GOLANG_DEPS_ARTIFACT)
.PHONY: _unzipGoDeps

# _zipBinArtifact zips the go dependencies so they can be passed along with a single zip file
_zipBinArtifact:
	zip -rq $(BIN_ARTIFACT) $(BIN_DIR)/
.PHONY: _zipBinArtifact

# _unzipBinArtifact unzips the artifact that contains the binary packages
_unzipBinArtifact:
	unzip -qo -d . $(BIN_ARTIFACT)
.PHONY: _unzipBinArtifact

##################
# Public Targets #
##################

# deps installs all dependencies for testing/building/deploying. This example only has golang dependencies
deps:
	$(COMPOSE_RUN_GOLANG) make _deps
	$(COMPOSE_RUN_MUSKETEERS) make _zipGoDeps
.PHONY: deps

# test tests the application
test: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _test
.PHONY: test

build: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _build
	$(COMPOSE_RUN_MUSKETEERS) make _zipBinArtifact
.PHONY: build

run: $(BIN_DIR)
	$(COMPOSE_RUN_GOLANG) make _run
.PHONY: run

shell:
	$(COMPOSE_RUN_GOLANG) bash
.PHONY: shell

# clean removes build artifacts and cleans docker environment
clean:
	$(COMPOSE_RUN_GOLANG) make _clean
	docker-compose down --remove-orphans
	-$(MAKE) dockerRemove removeEnvfile
.PHONY: clean

mock: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _mock
.PHONY: mock

dockerBuild:
	docker build --no-cache -t $(IMAGE_NAME) .
.PHONY: dockerBuild

dockerTest:
	docker run --rm $(IMAGE_NAME)
	docker run --rm $(IMAGE_NAME) version
.PHONY: dockerTest

dockerRemove:
	docker rmi -f $(IMAGE_NAME)
.PHONY: dockerRemove

###################
# Private Targets #
###################

_deps:
	dep ensure
.PHONY: _deps

_test:
	go test -coverprofile=$(PROFILE_NAME) ./...
.PHONY: _test

_build:
	APP_NAME=$(APP_NAME) VERSION=$(VERSION) ./scripts/build.sh
.PHONY: _build

_run:
	./$(EXECUTABLE)
.PHONY: _run

_install:
	go install
.PHONY: _install

_mock:
	go get -u github.com/vektra/mockery/.../
	mockery -dir=pkg -all -case=underscore -output=pkg/mocks
.PHONY: _mock

_htmlCover:
	go tool cover -html=$(PROFILE_NAME)
.PHONY: _htmlCover

_clean:
	rm -fr bin vendor
.PHONY: _clean

_tag:
	-git tag -d $(VERSION)
	-git push origin :refs/tags/$(VERSION)
	git tag $(VERSION)
	git push origin $(VERSION)
.PHONY: _tag