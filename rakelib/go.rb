# frozen_string_literal: true

require 'json'

require_relative 'paths'
require_relative 'shell'
require_relative 'go/deps'
require_relative 'go/tools'

module Go
  include Paths

  def self.require_go_tool!(cmd)
    Shell.require_command!(cmd, msg: format(<<~FMT, cmd: cmd, path: TOOLS_OUT_PATH))
      Go tool %<cmd>p not found in PATH.

      Make sure to install tools with `rake app:install:tools` and add tools installation
      directory %<path>s to PATH.
    FMT
  end

  def self.module(path)
    Shell.capture(GO_CMD, 'list', '-m', chdir: path)
  end

  def self.list(*args, json: false, json_parse: false, **options)
    args = ['-json', *args] if json
    out = Shell.capture(GO_CMD, 'list', *args, **options)
    json_parse ? JSON.parse(out) : out
  end

  def self.install_pkg(pkg, *args, output_path: nil, **options)
    options[:env] ||= {}
    options[:env][:GOBIN] = output_path if output_path

    Shell.run(GO_CMD, 'install', pkg, *args, **options)
  end

  def self.get_pkg(pkg, *args, **options)
    Shell.run(GO_CMD, 'get', *args, pkg, **options)
  end

  def self.build_pkg(pkg, *args, **options)
    Shell.run(GO_CMD, 'build', *args, pkg, **options)
  end

  def self.test(*args, **options)
    Shell.sh(GO_CMD, 'test', *args, **options)
  end

  def self.fmt(path, **options)
    Shell.sh(GOFMT_CMD, '-s', '-w', '-l', path, **options)
  end
end
