# frozen_string_literal: true

When('we get a pipeline') do
  @status = Idpctl.start_process('client', '--get', '1')
end

When('we try to get a missing pipeline') do
  @status = Idpctl.start_process('client', '--get', '2')
end

When('we try to get a pipeline with a missing token') do
  @status = Idpctl.start_process('missing', '--get', '1')
end
