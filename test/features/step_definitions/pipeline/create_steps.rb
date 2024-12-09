# frozen_string_literal: true

When('we create a new pipeline') do
  @status = Idpctl.start_process('client', '--create', 'pipeline')
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
