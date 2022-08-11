PKGS = $(shell go list ./... | grep -v /vendor/ | grep -v dbmodels)
SRCS = $(shell git ls-files '*.go')
SERVICES = $(shell cd ./cmd && go list -f '{{.Dir}}' ./...)
BRANCH := $(shell git symbolic-ref --short HEAD)

build:
	@ $(foreach file,$(SERVICES),echo $(file) && $(MAKE) -C $(file) build || exit 1;)