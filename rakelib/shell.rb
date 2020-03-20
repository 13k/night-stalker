# frozen_string_literal: true

require 'forwardable'

require 'rake'
require 'tty-command'
require 'tty-which'

require_relative 'logger'

module Shell
  extend FileUtils # Add rake's monkey-patched instance methods

  CommandNotFound = Class.new(StandardError)

  CMD_OPTIONS = %i[output color uuid printer].freeze

  class << self
    extend Forwardable

    # Add stdlib singleton methods
    delegate FileUtils::METHODS => FileUtils

    def which(cmd)
      TTY::Which.which(cmd)
    end

    def which!(cmd)
      path = which(cmd)

      raise CommandNotFound, format('Command %<cmd>p not found', cmd: cmd) if path.nil?

      path
    end

    alias orig_sh sh

    def sh(*args, env: {}, **options)
      orig_sh(env, *args.map(&:to_s), **options)
    end

    def run(*args, trace: true, quiet: false, **options)
      cmd_options = options.slice(*CMD_OPTIONS)
      cmd_options[:printer] ||= trace ? :pretty : :quiet

      options = options.reject { |k| CMD_OPTIONS.include?(k) }
      options[:only_output_on_error] = quiet

      TTY::Command.new(**cmd_options).run(*args, **options)
    end

    def capture(*args, trace: false, quiet: true, strip: true, **options)
      options = options.merge(trace: trace, quiet: quiet)
      result = run(*args, **options)
      strip ? result.out.strip : result.out
    end

    def require_env!(var, msg: nil, code: 1)
      return ENV[var] unless ENV[var].nil?

      msg ||= format('Environment variable %<var>p not set.', var: var)

      Logger.fail(msg, code)
    end

    def require_command!(cmd, msg: nil, code: 1)
      which!(cmd)
    rescue CommandNotFound => e
      msg ||= e.message

      Logger.fail(msg, code)
    end
  end
end
