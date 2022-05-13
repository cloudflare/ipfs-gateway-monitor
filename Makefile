# The import path is where your repository can be found.
# To import subpackages, always prepend the full import path.
# If you change this, run `make clean`. Read more: https://git.io/vM7zV
IMPORT_PATH := github.com/cloudflare/ipfs-gateway-monitor
GOPATH := $(shell go env GOPATH)

GO := go

# V := 1 # When V is set, print commands and build progress.

# Space separated patterns of packages to skip in list, test, format.
IGNORED_PACKAGES := /vendor/

SRC_FILES := $(shell find . -type f -name '*.go' $(foreach p, $(IGNORED_PACKAGES), -not -path ".$p*"))

.PHONY: all
all: build

# docker-compose and other utility targets.

.PHONY: compose-build
compose-build:
	docker-compose -f deployments/docker-compose.yaml build \
		--build-arg VERSION="$(VERSION)" \
		--build-arg DATE="$(DATE)"

.PHONY: compose-up
compose-up:
	docker-compose -f deployments/docker-compose.yaml up --no-build monitor

.PHONY: compose-down
compose-down:
	docker-compose -f deployments/docker-compose.yaml down

# Add all projects in the cmds directory (this is a work around for cmd/... not working).
CMDS := $(notdir $(wildcard $(CURDIR)/cmd/*))
.PHONY: %.gocmd
%.gocmd: ; $Q $(GO) build -o bin/$* $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/$*

.PHONY: %.goinstall
%.goinstall: %.gocmd ; cp bin/$* $(GOPATH)/bin

.PHONY: build
build: $(CMDS:%=%.gocmd)

.PHONY: install
install: $(CMDS:%=%.goinstall)

##### ^^^^^^ EDIT ABOVE ^^^^^^ #####

##### =====> Utility targets <===== #####

.PHONY: pretty
pretty:
	# reorganize imports and reformat
	GO111MODULE=off $(GO) get golang.org/x/tools/cmd/goimports
	$(GOPATH)/bin/goimports -w -d $(SRC_FILES)

.PHONY: clean test list

clean:
	$Q rm -rf bin target

test:
ifndef CI
	$Q $(GO) vet $(allpackages)
	$Q GODEBUG=cgocheck=2 go test -race $(allpackages)
else
	$Q ( $(GO) vet $(allpackages); echo $$? ) | \
	    tee target/vet.txt | sed '$$ d'; exit $$(tail -1 target/vet.txt)
	$Q ( GODEBUG=cgocheck=2 $(GO) test -v -race $(allpackages); echo $$? ) | \
	    tee target/vet.txt | sed '$$ d'; exit $$(tail -1 target/vet.txt)
endif

list:
	@echo $(allpackages)

##### =====> Internals <===== #####

VERSION          ?= $(shell git describe --tags --always --dirty="-dev")
DATE             ?= $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

_allpackages = $(shell ( $(GO) list ./... 2>&1 1>&3 | \
    grep -v -e "^$$" $(addprefix -e ,$(IGNORED_PACKAGES)) 1>&2 ) 3>&1 | \
    grep -v -e "^$$" $(addprefix -e ,$(IGNORED_PACKAGES)))

# memoize allpackages, so that it's executed only once and only if used
allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)

# Based on https://github.com/cloudflare/hellogopher - v1.1 - MIT License
#
# Copyright (c) 2017-2022 Cloudflare
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
