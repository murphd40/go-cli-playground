
## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

TMP ?= $(shell pwd)/.tmp
$(TMP):
	mkdir -p $(TMP)

PROTOC ?= $(LOCALBIN)/protoc
PROTOC_VERSION ?= 21.12
.PHONY: protoc
protoc: $(PROTOC) ## Download yq locally if necessary.
$(PROTOC): $(LOCALBIN) $(TMP)
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-osx-aarch_64.zip -o $(TMP)/protoc.zip
	unzip -uq $(TMP)/protoc.zip -d $(TMP)
	mv -f $(TMP)/bin/protoc $(LOCALBIN)

PROTOC_GEN_GO ?= $(LOCALBIN)/protoc-gen-go
PROTOC_GEN_GO_VERSION ?= v1.28
.PHONY: protoc-gen-go
protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)

PROTOC_GEN_GO_GRPC ?= $(LOCALBIN)/protoc-gen-go-grpc
PROTOC_GEN_GO_GRPC_VERSION ?= v1.2
.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

.PHONY: generate
generate: protoc protoc-gen-go protoc-gen-go-grpc
	@PATH=$(PATH):$(LOCALBIN) $(PROTOC) --go_out=. --go_opt=paths=source_relative examples/grpc/helloworld/proto/helloworld.proto --go-grpc_out=. --go-grpc_opt=paths=source_relative
	@PATH=$(PATH):$(LOCALBIN) $(PROTOC) --go_out=. --go_opt=paths=source_relative examples/grpc/pullserver/proto/pullserver.proto --go-grpc_out=. --go-grpc_opt=paths=source_relative


