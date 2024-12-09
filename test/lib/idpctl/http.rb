# frozen_string_literal: true

module Idpctl
  class Application < Sinatra::Application
    configure do
      set :server_settings, log_requests: true
    end

    post '/pipelines' do
      pipeline = JSON.parse(request.body.read)['pipeline']
      pipeline['id'] = 1

      res = {
        meta: {
          id: '1234'
        },
        pipeline: pipeline
      }

      res.to_json
    end

    get '/pipelines/:id' do |id|
      pipeline = JSON.parse(File.read('pipeline'))['pipeline']
      pipeline['id'] = id

      res = {
        meta: {
          id: '1234'
        },
        pipeline: pipeline
      }

      res.to_json
    end

    put '/pipelines/:id' do |id|
      pipeline = JSON.parse(request.body.read)['pipeline']
      pipeline['id'] = id

      res = {
        meta: {
          id: '1234'
        },
        pipeline: pipeline
      }

      res.to_json
    end
  end

  class HTTPServer < Nonnative::HTTPServer
    def app
      Application.new
    end
  end
end
