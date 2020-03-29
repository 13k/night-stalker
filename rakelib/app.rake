# frozen_string_literal: true

require_relative 'app_tasks'

namespace :app do
  namespace :go do
    desc 'Install app tools'
    task tools: AppTasks[:go_tools]

    desc 'Update direct dependencies'
    task deps: AppTasks[:go_deps]
  end

  desc 'Run app linters'
  task lint: ['app:lint:proto', 'app:lint:go', 'app:lint:ruby']

  namespace :lint do
    desc 'Run protobuf linters'
    task proto: AppTasks[:lint_proto]

    desc 'Run Go linters'
    task go: AppTasks[:lint_go]

    desc 'Run ruby linters'
    task ruby: AppTasks[:lint_ruby]
  end

  desc 'Run app tests'
  task test: ['app:test:go']

  namespace :test do
    desc 'Run Go tests'
    task go: AppTasks[:test_go]
  end

  desc 'Build app'
  task build: ['app:build:proto', 'app:build:commands']

  namespace :build do
    desc 'Build app command binaries'
    task commands: AppTasks[:build_commands]

    desc 'Compile app protobufs'
    task proto: AppTasks[:build_proto]
  end

  namespace :clean do
    desc 'Remove built app command binaries'
    task commands: AppTasks[:clean_commands]
  end

  namespace :format do
    desc 'Format Go files'
    task go: AppTasks[:format_go]
  end
end
