# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'idpctl/http'

module Idpctl
  class << self
    def client_config
      @client_config ||= Nonnative.configurations('.config/client.yml')
    end
  end
end
