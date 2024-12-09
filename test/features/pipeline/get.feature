Feature: Get a pipeline

  The ability to get pipelines.

  Scenario: Get a simple pipeline
    When we get a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully retrieved the pipeline" in the file "reports/pipeline.log"

  Scenario: Get a simple pipeline with a missing token
    When we try to get a pipeline with a missing token
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: open test: no such file or directory" in the file "reports/pipeline.log"

  Scenario: Get a missing pipeline
    When we try to get a missing pipeline
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: rest: missing pipeline" in the file "reports/pipeline.log"
