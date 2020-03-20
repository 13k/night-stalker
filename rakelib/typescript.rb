# frozen_string_literal: true

require_relative 'yarn'

module TypeScript
  def self.compile_project(tsconfig, chdir: nil)
    YARN.run('tsc', '-b', tsconfig, chdir: chdir)
  end
end
