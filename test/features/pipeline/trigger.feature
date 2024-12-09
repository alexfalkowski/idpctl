Feature: Trigger a pipeline

  The ability to trigger pipelines.

  Scenario: Trigger a pipeline
    When we trigger a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully triggered the pipeline" in the file "reports/pipeline.log"

  Scenario: Trigger a simple pipeline with a missing token
    When we try to trigger pipeline with a missing token
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: open test: no such file or directory" in the file "reports/pipeline.log"

  Scenario: Trigger a missing pipeline
    When we try to trigger a missing pipeline
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: rest: missing pipeline" in the file "reports/pipeline.log"
