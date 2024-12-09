# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'idpctl/http'

module Idpctl
  class << self
    def start_process(config, flag, param)
      cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', "file:.config/#{config}.yml", flag, param)
      pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

      _, status = Process.waitpid2(pid)

      status
    end
  end
end
