# frozen_string_literal: true

When('we create a new pipeline') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', 'file:.config/client.yml', '--create', 'pipeline')
  pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we get a pipeline') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', 'file:.config/client.yml', '--get', '1')
  pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('it should run successfully') do
  expect(@status.exitstatus).to eq(0)
end
