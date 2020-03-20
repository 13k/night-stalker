# frozen_string_literal: true

require_relative 'webapp_tasks'

namespace :webapp do
  desc 'Run webapp linters'
  task lint: WebAppTasks[:lint]

  desc 'Run webapp tests'
  task test: WebAppTasks[:test]

  desc 'Build webapp'
  task build: ['webapp:build:scripts', 'webapp:build:proto', 'webapp:build:app']

  namespace :build do
    desc 'Compile webapp scripts'
    task scripts: WebAppTasks[:build_scripts]

    desc 'Compile JS protobufs'
    task proto: WebAppTasks[:build_proto]

    desc 'Build webapp'
    task app: WebAppTasks[:build_app]
  end

  namespace :assets do
    desc 'Download hero images'
    task hero_images: WebAppTasks[:assets_hero_images]
  end
end
