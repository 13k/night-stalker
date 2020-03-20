# frozen_string_literal: true

# disable Ruby 2.7 warnings
$VERBOSE = nil

DRY_RUN = Rake::FileUtilsExt.nowrite
VERBOSE = Rake::FileUtilsExt.verbose

desc 'Run linters'
task lint: ['app:lint', 'webapp:lint']

desc 'Run tests'
task test: ['app:test', 'webapp:test']

desc 'Build app and webapp'
task build: ['app:build', 'webapp:build']

namespace :build do
  desc 'Generate protobufs'
  task proto: ['app:build:proto', 'webapp:build:proto']
end

task default: :test
