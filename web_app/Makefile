FIND := $(shell command -v gfind)
ifeq ($(strip $(FIND)),)
FIND := $(shell command -v find)
endif

YARN := $(shell command -v yarn)

PROTO_SRC = ../proto
PROTO_JSOUT = src/protocol
PROTOS = $(shell $(FIND) "$(PROTO_SRC)" -type f -name '*.proto')
PROTOS_JS_COMBINED = $(PROTO_JSOUT)/proto.js

$(PROTOS_JS_COMBINED): $(PROTOS)
	$(YARN) run pbjs -t static-module -w es6 --keep-case --force-long -o "$@" $(PROTOS)

.PHONY: proto-js
proto-js: $(PROTOS_JS_COMBINED)

.PHONY: proto
proto: proto-js
