Feature: Delete a pipeline

  The ability to delete pipelines.

  Scenario: Delete a simple pipeline
    When we delete a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully deleted the pipeline" in the file "reports/pipeline.log"

  Scenario: Delete a simple pipeline with a missing token
    When we try to delete a pipeline with a missing token
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: open test: no such file or directory" in the file "reports/pipeline.log"

  Scenario: Delete a missing pipeline
    When we try to delete a missing pipeline
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: rest: missing pipeline" in the file "reports/pipeline.log"
