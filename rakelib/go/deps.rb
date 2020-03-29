# frozen_string_literal: true

require 'json'

require_relative '../paths'

module Go
  module Deps
    include Paths

    def self.list(direct: false, update: false)
      args = ['-m', 'all']
      args = ['-u', *args] if update

      out = Go.list(*args, json: true, trace: true, chdir: ROOT_PATH)

      pkgs = out
        .scan(/^\{.+?^\}/m)
        .map(&JSON.method(:parse))
        .reject { |pkg| pkg['Main'] }

      pkgs = pkgs.reject { |pkg| pkg['Indirect'] } if direct
      pkgs = pkgs.select { |pkg| pkg['Update'] } if update

      pkgs
    end
  end
end
