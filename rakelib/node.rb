# frozen_string_literal: true

require_relative 'paths'
require_relative 'shell'

module Node
  include Paths

  def self.run(*args, chdir: nil, **env)
    Shell.run(NODE_CMD, *args, env: env, chdir: chdir)
  end
end
