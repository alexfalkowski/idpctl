Feature: Create a pipeline

  The ability to create pipelines.

  Scenario: Create a simple pipeline
    When we create a new pipeline
    Then it should run successfully
    And I should see a log entry of "successfully created the pipeline" in the file "reports/pipeline.log"

  Scenario: Create a simple pipeline with a missing token
    When we try to create a new pipeline with a missing token
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: open test: no such file or directory" in the file "reports/pipeline.log"
