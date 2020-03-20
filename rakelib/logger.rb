# frozen_string_literal: true

require 'forwardable'

require 'tty-logger'

module Logger
  class << self
    extend Forwardable

    delegate TTY::Logger::LOG_TYPES.keys => :default_logger
  end

  def self.default_logger
    @default_logger ||= TTY::Logger.new
  end

  def self.fail(msg, code = 1)
    fatal(msg)
    exit code
  end
end
