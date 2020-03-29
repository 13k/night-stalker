# frozen_string_literal: true

require_relative '../paths'

module Go
  module Tools
    include Paths

    def self.pkg_info
      @pkg_info ||= Go.list(
        '-tags', 'tools', "./#{TOOLS_SRC_RELPATH}",
        json: true, json_parse: true, chdir: ROOT_PATH,
      )
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
end
