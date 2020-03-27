# frozen_string_literal: true

require 'pathname'

require_relative 'shell'

module Paths
  ROOT_PATH = Pathname.new(File.expand_path('..', __dir__)).freeze
  WEBAPP_PATH = ROOT_PATH / 'balanar'

  CMD_PKG_PATH = ROOT_PATH / 'cmd'
  CMD_OUT_PATH = ROOT_PATH / 'bin'

  PROTO_SRC_PATH = ROOT_PATH / 'protobuf'
  PROTO_GO_OUT_PATH = ROOT_PATH / 'internal' / 'protobuf'
  PROTO_JS_OUT_PATH = WEBAPP_PATH / 'src' / 'protocol'
  PROTO_JS_OUT_COMBINED = PROTO_JS_OUT_PATH / 'proto.js'

  TOOLS_SRC_PATH = ROOT_PATH / 'tools'
  TOOLS_SRC_RELPATH = TOOLS_SRC_PATH.relative_path_from(ROOT_PATH)
  TOOLS_OUT_PATH = TOOLS_SRC_PATH / 'bin'

  WEBAPP_SCRIPTS_SRC_PATH = WEBAPP_PATH / 'scripts'
  WEBAPP_SCRIPTS_TSCONFIG = WEBAPP_SCRIPTS_SRC_PATH / 'tsconfig.json'
  WEBAPP_SCRIPTS_BUILD_PATH = WEBAPP_SCRIPTS_SRC_PATH / 'build'
  WEBAPP_SCRIPT_HERO_IMAGES = WEBAPP_SCRIPTS_BUILD_PATH / 'gen_hero_images.js'

  WEBAPP_ASSETS_PATH = WEBAPP_PATH / 'src' / 'assets'
  WEBAPP_IMAGES_HEROES_CDN_PATH = WEBAPP_ASSETS_PATH / 'images' / 'heroes' / 'cdn'

  GIT_CMD = Shell.require_command!(ENV.fetch('GIT', 'git'))
  PROTOC_CMD = Shell.require_command!(ENV.fetch('PROTOC', 'protoc'))
  GO_CMD = Shell.require_command!(ENV.fetch('GO', 'go'))
  GOFMT_CMD = Shell.require_command!(ENV.fetch('GOFMT', 'gofmt'))
  NODE_CMD = Shell.require_command!(ENV.fetch('NODE', 'node'))
  YARN_CMD = Shell.require_command!(ENV.fetch('YARN', 'yarn'))
end
