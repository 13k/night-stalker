# frozen_string_literal: true

require_relative 'paths'
require_relative 'shell'

module YARN
  include Paths

  def self.run(task, *args, chdir: nil, **env)
    cmd = [
      YARN_CMD, 'run', task,
      *args,
    ]

    Shell.sh(*cmd, chdir: chdir, env: env)
  end
end
