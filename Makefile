FIND := $(shell command -v gfind)
ifeq ($(strip $(FIND)),)
FIND := $(shell command -v find)
endif

PROTOC := $(shell command -v protoc)

TOOLS_PATH := bin
PROTO_SRC := proto
PROTO_GOOUT := internal/protocol
PROTO_JSOUT := web_app/src/protocol

protobufs := $(shell $(FIND) "$(PROTO_SRC)" -type f -name '*.proto' -printf '%P\n')
protos_go := $(patsubst %.proto,$(PROTO_GOOUT)/%.pb.go,$(protobufs))
protos_js := $(PROTO_JSOUT)/enums_pb.js

$(PROTO_GOOUT)/%.pb.go: $(PROTO_SRC)/%.proto
	$(PROTOC) -I "$(PROTO_SRC)" "--go_out=$(PROTO_GOOUT)" "$<"

.PHONY: proto-go
proto-go: $(protos_go)

$(PROTO_JSOUT)/%_pb.js: $(PROTO_SRC)/%.proto
	$(PROTOC) -I "$(PROTO_SRC)" "--js_out=import_style=commonjs_strict,binary:$(PROTO_JSOUT)" "$<"

.PHONY: proto-js
proto-js: $(protos_js)

.PHONY: proto
proto: proto-go proto-js

$(TOOLS_PATH): tools.go
	bash hack/install-tools.bash
	@touch $(TOOLS_PATH)

.PHONY: install-tools
install-tools: $(TOOLS_PATH)
