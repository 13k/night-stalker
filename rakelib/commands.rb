# frozen_string_literal: true

require_relative 'paths'
require_relative 'shell'

module Commands
  include Paths

  def self.version
    return @version if defined?(@version)

    @version = Shell.capture(
      GIT_CMD, 'describe', '--tags', '--exact-match', 'HEAD',
      raise_error: false, printer: :null,
    )

    @version ||= 'dev'
  end

  def self.revision
    @revision ||= Shell.capture(GIT_CMD, 'rev-parse', '--short', 'HEAD')
  end

  def self.specs
    @specs ||= CMD_PKG_PATH.glob('*').map do |pkg_path|
      name = pkg_path.basename.to_s
      outname = name
      outname += "-#{revision}" if ENV['CMD_REV']

      pkg_rel_path = pkg_path.relative_path_from(CMD_PKG_PATH)
      pkg_rel_dir = pkg_rel_path.dirname

      {
        name: name,
        input: pkg_path,
        output: CMD_OUT_PATH / pkg_rel_dir / outname,
        version: version,
        revision: revision,
      }
    end
  end

  def self.built
    @built ||= CMD_PKG_PATH.glob('*').map do |pkg_path|
      name = pkg_path.basename.to_s
      pkg_rel_path = pkg_path.relative_path_from(CMD_PKG_PATH)
      pkg_rel_dir = pkg_rel_path.dirname
      patterns = [pkg_rel_dir / name, pkg_rel_dir / "#{name}-*"]
      artifacts = patterns.flat_map(&CMD_OUT_PATH.method(:glob))
      [name, artifacts]
    end.to_h
  end
end
