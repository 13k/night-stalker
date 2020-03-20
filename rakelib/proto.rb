# frozen_string_literal: true

require_relative 'paths'
require_relative 'shell'
require_relative 'yarn'

module Protobuf
  include Paths

  GO_OUT_EXT = '.pb.go'

  def self.compile_go(input, proto_path, go_out_path, **go_out_options)
    proto_path_args = Array(proto_path).map { |p| "--proto_path=#{p}" }
    go_out_opts = go_out_options.map { |k, v| v == true ? k.to_s : "#{k}=#{v}" }

    cmd = [
      PROTOC_CMD,
      *proto_path_args,
      "--go_out=#{go_out_opts.join(',')}:#{go_out_path}",
      input,
    ]

    Shell.run(*cmd)
  end

  def self.compile_js(inputs, output, chdir: nil)
    cmd = [
      'pbjs',
      '-t', 'static-module',
      '-w', 'es6',
      '--keep-case', '--force-long',
      '-o', output,
      *inputs,
    ]

    YARN.run(*cmd, chdir: chdir)
  end

  def self.sources
    @sources ||= PROTO_SRC_PATH.glob('**/*.proto')
  end

  def self.specs_go
    @specs_go ||= sources.map do |src_path|
      outname = src_path.basename.sub_ext(GO_OUT_EXT)
      rel_path = src_path.relative_path_from(PROTO_SRC_PATH)
      rel_dir = rel_path.dirname

      {
        input: src_path,
        output: PROTO_GO_OUT_PATH / rel_dir / outname,
      }
    end
  end

  def self.specs_js
    @specs_js ||= [
      {
        inputs: sources,
        output: PROTO_JS_OUT_COMBINED,
      },
    ]
  end
end
