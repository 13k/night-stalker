FIND := $(shell command -v gfind)
FIND ?= $(shell command -v find)

TOOLS_PATH := bin
PROTOBUFS_PATH := internal/protocol

protobufs := $(shell $(FIND) $(PROTOBUFS_PATH) -type f -name '*.proto' -printf '%P\n')
goprotos := $(patsubst %.proto,%.pb.go,$(protobufs))

%.pb.go: %.proto
	protoc -I . --go_out=. $<

.PHONY: proto
proto: $(goprotos)

$(TOOLS_PATH): tools.go
	bash hack/install-tools.bash
	@touch $(TOOLS_PATH)

.PHONY: tools
tools: $(TOOLS_PATH)
