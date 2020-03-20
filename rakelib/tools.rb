# frozen_string_literal: true

require 'json'

require_relative 'go'
require_relative 'paths'

module Tools
  include Paths

  def self.pkg_info
    @pkg_info ||= Go.instrospect_pkg("./#{TOOLS_SRC_RELPATH}", '-tags', 'tools')
  end

  def self.specs
    reqs = pkg_info.fetch('GoFiles').map { |f| TOOLS_SRC_PATH / f }

    pkg_info.fetch('Imports').map do |import_path|
      name = File.basename(import_path)

      {
        name: name,
        reqs: reqs,
        pkg: import_path,
        output: TOOLS_OUT_PATH / name,
      }
    end
  end
end
