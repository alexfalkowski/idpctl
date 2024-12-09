# frozen_string_literal: true

When('we update a pipeline') do
  @status = Idpctl.start_process('client', '--update', '1:pipeline')
end

When('we try to update a missing pipeline') do
  @status = Idpctl.start_process('client', '--update', '2:pipeline')
end

When('we update an invalid pipeline') do
  @status = Idpctl.start_process('client', '--update', 'invalid')
end

When('we try to update a pipeline with a missing token') do
  @status = Idpctl.start_process('missing', '--update', '1:pipeline')
end
