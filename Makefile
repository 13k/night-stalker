FIND := $(shell command -v gfind)
ifeq ($(strip $(FIND)),)
FIND := $(shell command -v find)
endif

PROTOC := $(shell command -v protoc)

SUBDIRS = web_app
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
proto-js: web_app

.PHONY: proto
proto: proto-go proto-js

$(TOOLS_PATH): tools.go
	bash hack/install-tools.bash
	@touch $(TOOLS_PATH)

.PHONY: install-tools
install-tools: $(TOOLS_PATH)
