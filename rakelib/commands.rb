# frozen_string_literal: true

require_relative 'paths'
require_relative 'shell'

module Commands
  include Paths

  def self.specs
    @specs ||= CMD_PKG_PATH.glob('*').map do |pkg_path|
      name = pkg_path.basename.to_s
      outname = name

      if ENV['CMD_TAG']
        tag = Shell.capture(GIT_CMD, 'rev-parse', '--short', 'HEAD')
        outname += "-#{tag}"
      end

      pkg_rel_path = pkg_path.relative_path_from(CMD_PKG_PATH)
      pkg_rel_dir = pkg_rel_path.dirname

      {
        name: name,
        input: pkg_path,
        output: CMD_OUT_PATH / pkg_rel_dir / outname,
      }
    end
  end
end
