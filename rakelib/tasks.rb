# frozen_string_literal: true

require 'securerandom'

require 'rake'

require_relative 'go'
require_relative 'paths'
require_relative 'shell'
require_relative 'typescript'
require_relative 'yarn'

module Tasks
  include Rake::DSL
  include Paths

  def self.extended(base)
    base.class_eval do
      @tasks = {}

      def self.gen_task(name, &block)
        @tasks[name] = block.call
      end

      def self.[](name)
        @tasks[name]
      end
    end
  end

  def anon_task(*reqs, &block)
    task(rand_task_id => reqs, &block)
  end

  def anon_multitask(*reqs, &block)
    multitask(rand_task_id => reqs, &block)
  end

  def require_env(var)
    anon_task do
      Shell.require_env!(var)
    end
  end

  def require_command(cmd)
    anon_task do
      Shell.require_command!(cmd)
    end
  end

  def run_command(*args, **options)
    anon_task do
      Shell.run(*args, **options)
    end
  end

  def exec_command(*args, **options)
    anon_task do
      Shell.sh(*args, **options)
    end
  end

  def file_rm(*paths)
    anon_task do
      rm_f(paths)
    end
  end

  def require_go_tool(cmd)
    anon_task do
      Go.require_go_tool!(cmd)
    end
  end

  def install_go_tool(pkg:, reqs:, output:, **)
    file output => reqs do
      Go.install_pkg(pkg, output_path: TOOLS_OUT_PATH)
    end
  end

  def get_go_pkg(pkg)
    anon_task do
      Go.get_pkg(pkg, chdir: ROOT_PATH)
    end
  end

  def compile_go_command(input, output, *args, **options)
    task output do
      input_arg = input.absolute? ? input.relative_path_from(ROOT_PATH) : input
      args = [*args, '-o', output]
      Go.build_pkg("./#{input_arg}", *args, **options)
    end
  end

  def compile_proto_go(input:, output:, **)
    file output => input do
      Protobuf.compile_go(input, PROTO_SRC_PATH, PROTO_GO_OUT_PATH, paths: 'source_relative')
    end
  end

  def compile_proto_js(inputs:, output:)
    file output => inputs do
      Protobuf.compile_js(inputs, output, chdir: WEBAPP_PATH)
    end
  end

  def compile_ts(input:, output:)
    file output => [WEBAPP_SCRIPTS_TSCONFIG, input] do
      TypeScript.compile_project(WEBAPP_SCRIPTS_TSCONFIG, chdir: WEBAPP_PATH)
    end
  end

  def yarn_run(task)
    anon_task do
      YARN.run(task, chdir: WEBAPP_PATH)
    end
  end

  protected

  def rand_task_id
    SecureRandom.hex(8)
  end
end
