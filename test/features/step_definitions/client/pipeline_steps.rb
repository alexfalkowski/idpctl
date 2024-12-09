# frozen_string_literal: true

When('we create a new pipeline') do
  @status = Idpctl.start_process('client', '--create', 'pipeline')
end

When('we get a pipeline') do
  @status = Idpctl.start_process('client', '--get', '1')
end

When('we update a pipeline') do
  @status = Idpctl.start_process('client', '--update', '1:pipeline')
end

When('we update an pipeline') do
  @status = Idpctl.start_process('client', '--update', 'invalid')
end

When('we delete a pipeline') do
  @status = Idpctl.start_process('client', '--delete', '1')
end

When('we trigger a pipeline') do
  @status = Idpctl.start_process('client', '--trigger', '1')
end

When('we try to create a new pipeline with a missing token') do
  @status = Idpctl.start_process('missing', '--create', 'pipeline')
end

Then('it should run successfully') do
  expect(@status.exitstatus).to eq(0)
end

Then('it should run unsuccessfully') do
  expect(@status.exitstatus).to eq(1)
end
