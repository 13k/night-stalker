# frozen_string_literal: true

require_relative 'paths'
require_relative 'shell'

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

  def self.instrospect_pkg(pkg, *args)
    json = Shell.capture(GO_CMD, 'list', '-json', *args, pkg)
    JSON.parse(json)
  end

  def self.install_pkg(pkg, *args, output_path: nil, **options)
    options[:env] ||= {}
    options[:env][:GOBIN] = output_path if output_path

    Shell.run(GO_CMD, 'install', pkg, *args, **options)
  end

  def self.build_pkg(pkg, *args, **options)
    Shell.run(GO_CMD, 'build', *args, pkg, **options)
  end

  def self.test(*args, **options)
    Shell.sh(GO_CMD, 'test', *args, **options)
  end

  def self.format(path, **options)
    Shell.sh(GOFMT_CMD, '-s', '-w', path, **options)
  end
end
