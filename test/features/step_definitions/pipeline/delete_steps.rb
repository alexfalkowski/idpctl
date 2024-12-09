# frozen_string_literal: true

When('we delete a pipeline') do
  @status = Idpctl.start_process('client', '--delete', '1')
end

When('we try to delete a missing pipeline') do
  @status = Idpctl.start_process('client', '--delete', '2')
end

When('we try to delete a pipeline with a missing token') do
  @status = Idpctl.start_process('missing', '--delete', '1')
end
