# frozen_string_literal: true

When('we trigger a pipeline') do
  @status = Idpctl.start_process('client', '--trigger', '1')
end

When('we try to trigger a missing pipeline') do
  @status = Idpctl.start_process('client', '--trigger', '2')
end

When('we try to trigger pipeline with a missing token') do
  @status = Idpctl.start_process('missing', '--trigger', '1')
end
