FIND := $(shell command -v gfind)
ifeq ($(strip $(FIND)),)
FIND := $(shell command -v find)
endif

PROTOC := $(shell command -v protoc)

SUBDIRS = balanar
PROTO_SRC = proto
PROTO_GOOUT = internal/protocol
PROTOS = $(shell $(FIND) "$(PROTO_SRC)" -type f -name '*.proto' -printf '%P\n')
PROTOS_GO = $(patsubst %.proto,$(PROTO_GOOUT)/%.pb.go,$(PROTOS))
TOOLS_PATH = bin

.PHONY: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS) FIND=$(FIND)

$(PROTO_GOOUT)/%.pb.go: $(PROTO_SRC)/%.proto
	$(PROTOC) -I "$(PROTO_SRC)" "--go_out=$(PROTO_GOOUT)" "$<"

.PHONY: proto-go
proto-go: $(PROTOS_GO)

.PHONY: proto-js
proto-js: balanar

.PHONY: proto
proto: proto-go proto-js

$(TOOLS_PATH): tools.go
	hack/go-install-tools.sh
	@touch $(TOOLS_PATH)

.PHONY: go-install-tools
go-install-tools: $(TOOLS_PATH)

.PHONY: go-mod-update
go-mod-update:
	hack/go-mod-update.sh
