# frozen_string_literal: true

require_relative 'node'
require_relative 'paths'
require_relative 'proto'
require_relative 'tasks'
require_relative 'webapp/scripts'

module WebAppTasks
  include Paths
  extend Tasks

  gen_task :lint do
    yarn_run('lint')
  end

  gen_task :build_proto do
    anon_task(*Protobuf.specs_js.map { |spec| compile_proto_js(**spec) })
  end

  gen_task :build_scripts do
    anon_task(*WebApp::Scripts.specs.map { |spec| compile_ts(**spec) })
  end

  gen_task :build_app do
    yarn_run('build')
  end

  gen_task :assets_hero_images do
    anon_task(require_env('HEROES_KV'), self[:build_scripts]) do
      Node.run(
        WEBAPP_SCRIPT_HERO_IMAGES,
        ENV.fetch('HEROES_KV'),
        WEBAPP_IMAGES_HEROES_CDN_PATH,
        chdir: WEBAPP_PATH,
      )
    end
  end
end
