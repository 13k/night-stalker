FIND := $(shell command -v gfind)
ifeq ($(strip $(FIND)),)
FIND := $(shell command -v find)
endif

YARN := $(shell command -v yarn)

PROTO_REL_SRC := ../$(PROTO_SRC)
PROTO_JSOUT = src/protocol
PROTOS = $(shell $(FIND) "$(PROTO_REL_SRC)" -type f -name '*.proto')
PROTOS_JS_COMBINED = $(PROTO_JSOUT)/proto.js
SCRIPTS_SRC = scripts
SCRIPTS_BUILD_DIR = dist/scripts
SCRIPTS = $(shell $(FIND) "$(SCRIPTS_SRC)" -type f -name '*.ts' -printf '%P\n')
SCRIPTS_COMPILED = $(patsubst %.ts,$(SCRIPTS_BUILD_DIR)/%.js,$(SCRIPTS))
SCRIPT_HERO_IMAGES = $(SCRIPTS_BUILD_DIR)/gen_hero_images.js
SCRIPT_HERO_ALIASES = $(SCRIPTS_BUILD_DIR)/gen_hero_aliases.js
HERO_ALIASES_FILE = src/dota2/hero_aliases.json
HERO_CDN_IMAGES_DIR = src/assets/images/heroes/cdn

$(PROTOS_JS_COMBINED): $(PROTOS)
	$(YARN) run pbjs -t static-module -w es6 --keep-case --force-long -o "$@" $(PROTOS)

.PHONY: proto-js
proto-js: $(PROTOS_JS_COMBINED)

.PHONY: proto
proto: proto-js

$(SCRIPTS_BUILD_DIR)/%.js: $(SCRIPTS_SRC)/tsconfig.json $(SCRIPTS_SRC)/%.ts
	$(YARN) run tsc -b $(SCRIPTS_SRC)/tsconfig.json

$(HERO_ALIASES_FILE): $(SCRIPT_HERO_ALIASES) $(HEROES_KV)
	node "$(SCRIPT_HERO_ALIASES)" $(HEROES_KV) "$@"

.PHONY: hero-aliases
hero-aliases: $(HERO_ALIASES_FILE)

.PHONY: hero-images
hero-images: $(SCRIPT_HERO_IMAGES) $(HEROES_KV)
	node "$(SCRIPT_HERO_IMAGES)" $(HEROES_KV) "$(HERO_CDN_IMAGES_DIR)"
