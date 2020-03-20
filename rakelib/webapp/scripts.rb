# frozen_string_literal: true

require_relative '../paths'

module WebApp
  module Scripts
    include Paths

    def self.specs
      @specs ||= WEBAPP_SCRIPTS_SRC_PATH.glob('**/*.ts').map do |src_path|
        rel_path = src_path.relative_path_from(WEBAPP_SCRIPTS_SRC_PATH)
        rel_out = rel_path.sub_ext('.js')
        output = WEBAPP_SCRIPTS_BUILD_PATH / rel_out

        {
          input: src_path,
          output: output,
        }
      end
    end
  end
end
