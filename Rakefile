# frozen_string_literal: true

$VERBOSE = nil

require 'json'
require 'pathname'

require 'tty-command'
require 'tty-logger'
require 'tty-which'

NOWRITE = nowrite
VERBOSE = verbose
LOG = TTY::Logger.new

class Shell
  CommandNotFound = Class.new(StandardError)

  CMD_OPTIONS = %i[output color uuid printer].freeze

  def self.which(cmd)
    TTY::Which.which(cmd)
  end

  def self.which!(cmd)
    path = which(cmd)

    raise CommandNotFound, format('Command %<cmd>p not found', cmd: cmd) if path.nil?

    path
  end

  def self.run(*args, trace: true, quiet: false, **options)
    cmd_options = options.slice(*CMD_OPTIONS)
    cmd_options[:printer] = trace ? :pretty : :quiet

    options = options.reject { |k| CMD_OPTIONS.include?(k) }
    options[:only_output_on_error] = quiet

    TTY::Command.new(**cmd_options).run(*args, **options)
  end

  def self.capture(*args, trace: false, quiet: true, strip: true, **options)
    options = options.merge(trace: trace, quiet: quiet)
    result = run(*args, **options)
    strip ? result.out.strip : result.out
  end
end

begin
  GIT = Shell.which!(ENV.fetch('GIT', 'git'))
  PROTOC = Shell.which!(ENV.fetch('PROTOC', 'protoc'))
  GO = Shell.which!(ENV.fetch('GO', 'go'))
  NODE = Shell.which!(ENV.fetch('NODE', 'node'))
  YARN = Shell.which!(ENV.fetch('YARN', 'yarn'))
rescue Shell::CommandNotFound => e
  LOG.fatal(e.message)
  exit 1
end

ROOT_PATH = Pathname.new(__dir__)
WEBAPP_PATH = ROOT_PATH / 'balanar'

CMD_PKG_PATH = ROOT_PATH / 'cmd'
CMD_OUT_PATH = ROOT_PATH / 'bin'

PROTO_SRC_PATH = ROOT_PATH / 'protobuf'
PROTO_SRC_RELPATH = PROTO_SRC_PATH.relative_path_from(ROOT_PATH)
PROTO_GO_OUT_PATH = ROOT_PATH / 'internal' / 'protobuf'
PROTO_GO_OUT_RELPATH = PROTO_GO_OUT_PATH.relative_path_from(ROOT_PATH)
PROTO_JS_OUT_PATH = WEBAPP_PATH / 'src' / 'protocol'

TOOLS_SRC_PATH = ROOT_PATH / 'tools'
TOOLS_SRC_RELPATH = TOOLS_SRC_PATH.relative_path_from(ROOT_PATH)
TOOLS_OUT_PATH = TOOLS_SRC_PATH / 'bin'

WEBAPP_SCRIPTS_SRC_PATH = WEBAPP_PATH / 'scripts'
WEBAPP_SCRIPTS_TSCONFIG = WEBAPP_SCRIPTS_SRC_PATH / 'tsconfig.json'
WEBAPP_SCRIPTS_BUILD_PATH = WEBAPP_SCRIPTS_SRC_PATH / 'build'
WEBAPP_SCRIPT_HERO_IMAGES = WEBAPP_SCRIPTS_BUILD_PATH / 'gen_hero_images.js'

WEBAPP_ASSETS_PATH = WEBAPP_PATH / 'src' / 'assets'
WEBAPP_IMAGES_HEROES_CDN_PATH = WEBAPP_ASSETS_PATH / 'images' / 'heroes' / 'cdn'

def require_env!(var)
  return ENV[var] unless ENV[var].nil?

  LOG.fatal(format(<<~FMT, var: var))
    Environment variable %<var>p not set.
  FMT

  exit 1
end

def require_go_tool!(cmd)
  Shell.which!(cmd)
rescue Shell::CommandNotFound
  LOG.fatal(format(<<~FMT, cmd: cmd, path: TOOLS_OUT_PATH))
    Go tool %<cmd>p not found in PATH.

    Make sure to install tools with `rake install:tools` and add tools installation directory
    %<path>s to PATH.
  FMT

  exit 1
end

def install_go_tool(pkg:, reqs:, output:, **)
  file output => reqs do
    Shell.run(GO, 'install', pkg, env: { GOBIN: TOOLS_OUT_PATH })
  end
end

def run_go_tests
  sh(GO, 'test', './...')
end

def run_go_linters
  sh('golangci-lint', 'run')
end

def compile_go_package(pkg_path:, output:)
  reqs = pkg_path.glob('*.go')

  file output => reqs do
    input = pkg_path.absolute? ? pkg_path.relative_path_from(ROOT_PATH) : pkg_path
    output = output.relative_path_from(ROOT_PATH) if output.absolute?

    Shell.run(GO, 'build', '-o', output, "./#{input}")
  end
end

def compile_go_proto_file(input:, output:)
  file output => input do
    Shell.run(
      PROTOC,
      '-I', PROTO_SRC_RELPATH,
      "--go_out=paths=source_relative:#{PROTO_GO_OUT_RELPATH}",
      input.relative_path_from(ROOT_PATH),
    )
  end
end

def compile_js_proto_file(inputs:, output:)
  file output => inputs do
    inputs = inputs.map { |i| i.absolute? ? i.relative_path_from(ROOT_PATH) : i }
    output = output.relative_path_from(ROOT_PATH) if output.absolute?
    cmd = [
      YARN, 'run', 'pbjs',
      '-t', 'static-module',
      '-w', 'es6',
      '--keep-case', '--force-long',
      '-o', output,
      *inputs,
    ]

    Shell.run(*cmd, chdir: WEBAPP_PATH)
  end
end

def compile_ts_script(input:, output:)
  file output => [WEBAPP_SCRIPTS_TSCONFIG, input] do
    cmd = [
      YARN, 'run', 'tsc',
      '-b', WEBAPP_SCRIPTS_TSCONFIG,
    ]

    Shell.run(*cmd, chdir: WEBAPP_PATH)
  end
end

def cmds
  CMD_PKG_PATH.glob('*').map do |pkg_path|
    cmd_name = pkg_path.basename.to_s

    if ENV['CMD_TAG']
      tag = Shell.capture(GIT, 'rev-parse', '--short', 'HEAD')
      cmd_name += "-#{tag}"
    end

    pkg_rel_path = pkg_path.relative_path_from(CMD_PKG_PATH)
    pkg_rel_dir = pkg_rel_path.dirname

    {
      pkg_path: pkg_path,
      output: CMD_OUT_PATH / pkg_rel_dir / cmd_name,
    }
  end
end

@cmd_tasks = cmds.map do |spec|
  compile_go_package(**spec)
end

def proto_sources
  @proto_sources ||= PROTO_SRC_PATH.glob('**/*.proto')
end

def protos_go
  proto_sources.map do |src_path|
    outname = src_path.basename.sub_ext('.pb.go')
    rel_path = src_path.relative_path_from(PROTO_SRC_PATH)
    rel_dir = rel_path.dirname

    {
      input: src_path,
      output: PROTO_GO_OUT_PATH / rel_dir / outname,
    }
  end
end

@proto_go_tasks = protos_go.map do |spec|
  compile_go_proto_file(**spec)
end

def protos_js
  [
    {
      inputs: proto_sources,
      output: PROTO_JS_OUT_PATH / 'proto.js',
    },
  ]
end

@proto_js_tasks = protos_js.map do |spec|
  compile_js_proto_file(**spec)
end

def tools_pkg_info
  @tools_pkg_info ||= begin
      json = Shell.capture(GO, 'list', '-json', '-tags', 'tools', "./#{TOOLS_SRC_RELPATH}")
      JSON.parse(json)
    end
end

def tools
  reqs = tools_pkg_info.fetch('GoFiles').map { |f| TOOLS_SRC_PATH / f }

  tools_pkg_info.fetch('Imports').map do |tool_pkg|
    tool_name = File.basename(tool_pkg)

    {
      name: tool_name,
      reqs: reqs,
      pkg: tool_pkg,
      output: TOOLS_OUT_PATH / tool_name,
    }
  end
end

@tools_tasks = tools.map do |spec|
  install_go_tool(**spec)
end

def webapp_scripts
  WEBAPP_SCRIPTS_SRC_PATH.glob('**/*.ts').map do |src_path|
    rel_path = src_path.relative_path_from(WEBAPP_SCRIPTS_SRC_PATH)
    rel_out = rel_path.sub_ext('.js')
    output = WEBAPP_SCRIPTS_BUILD_PATH / rel_out

    {
      input: src_path,
      output: output,
    }
  end
end

@webapp_scripts_tasks = webapp_scripts.map do |spec|
  compile_ts_script(**spec)
end

namespace :env do
  namespace :require do
    task :heroes_kv do
      require_env!('HEROES_KV')
    end
  end
end

namespace :go do
  namespace :tools do
    task install: @tools_tasks

    namespace :require do
      tools.each do |spec|
        task(spec.fetch(:name)) do
          require_go_tool!(spec.fetch(:name))
        end
      end
    end
  end
end

namespace :install do
  desc 'Install Go tools'
  task tools: 'go:tools:install'
end

namespace :build do
  desc 'Build command binaries'
  multitask commands: @cmd_tasks

  namespace :proto do
    multitask _go: @proto_go_tasks

    desc 'Compile Go protobufs'
    task go: ['go:tools:require:protoc-gen-go', 'build:proto:_go']

    desc 'Compile JS protobufs'
    multitask js: @proto_js_tasks
  end

  task proto: ['build:proto:go', 'build:proto:js']
end

task build: ['build:proto', 'build:commands']

namespace :test do
  desc 'Run Go tests'
  task :go do
    run_go_tests
  end
end

task test: ['test:go']

namespace :lint do
  desc 'Run Go linters'
  task go: ['go:tools:require:golangci-lint'] do
    run_go_linters
  end
end

task lint: ['lint:go']

namespace :webapp do
  namespace :build do
    desc 'Compile webapp scripts'
    task scripts: @webapp_scripts_tasks
  end

  desc 'Download hero images'
  task hero_images: ['env:require:heroes_kv', 'webapp:build:scripts'] do
    Shell.run(
      NODE, WEBAPP_SCRIPT_HERO_IMAGES,
      ENV.fetch('HEROES_KV'),
      WEBAPP_IMAGES_HEROES_CDN_PATH,
    )
  end
end

task default: :test
