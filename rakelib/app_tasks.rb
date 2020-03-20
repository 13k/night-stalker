# frozen_string_literal: true

require_relative 'commands'
require_relative 'go'
require_relative 'paths'
require_relative 'proto'
require_relative 'shell'
require_relative 'tasks'
require_relative 'tools'

module AppTasks
  include Paths
  extend Tasks

  gen_task :install_tools do
    anon_task(*Tools.specs.map { |spec| install_go_tool(**spec) })
  end

  gen_task :build_proto do
    compile_tasks = anon_multitask(*Protobuf.specs_go.map { |spec| compile_proto_go(**spec) })
    anon_task(require_go_tool('protoc-gen-go'), compile_tasks)
  end

  gen_task :build_commands do
    anon_task(*Commands.specs.map { |spec| compile_go_command(**spec) })
  end

  gen_task :lint_proto do
    anon_task(
      require_command('buf'),
      require_command('protoc-gen-buf-check-lint'),
      exec_command('buf', 'check', 'lint', chdir: ROOT_PATH),
    )
  end

  gen_task :lint_go do
    anon_task(
      require_command('golangci-lint'),
      exec_command('golangci-lint', 'run', chdir: ROOT_PATH),
    )
  end

  gen_task :lint_ruby do
    exec_command('bundle', 'exec', 'rubocop', chdir: ROOT_PATH)
  end

  gen_task :test_go do
    anon_task do
      Go.test('./...', chdir: ROOT_PATH)
    end
  end
end
