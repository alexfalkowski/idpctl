Feature: Update a pipeline

  The ability to update pipelines.

  Scenario: Update a simple pipeline
    When we update a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully updated the pipeline" in the file "reports/pipeline.log"

  Scenario: Update an invalid pipeline
    When we update an invalid pipeline
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: invalid update format, this should be id:path" in the file "reports/pipeline.log"

  Scenario: Update a simple pipeline with a missing token
    When we try to update a pipeline with a missing token
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: open test: no such file or directory" in the file "reports/pipeline.log"

  Scenario: Update a missing pipeline
    When we try to update a missing pipeline
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: rest: missing pipeline" in the file "reports/pipeline.log"
