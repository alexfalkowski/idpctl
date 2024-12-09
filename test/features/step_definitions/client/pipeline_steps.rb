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

When('we update a pipeline') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', 'file:.config/client.yml', '--update', '1:pipeline')
  pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we delete a pipeline') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', 'file:.config/client.yml', '--delete', '1')
  pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we trigger a pipeline') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', 'file:.config/client.yml', '--trigger', '1')
  pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('we try to create a new pipeline with a missing token') do
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../idpctl', 'pipeline', '-i', 'file:.config/missing.yml', '--create', 'pipeline')
  pid = spawn({}, cmd, %i[out err] => ['reports/pipeline.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('it should run successfully') do
  expect(@status.exitstatus).to eq(0)
end

Then('it should run unsuccessfully') do
  expect(@status.exitstatus).to eq(1)
end
