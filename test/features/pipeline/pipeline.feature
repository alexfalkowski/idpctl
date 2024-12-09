Feature: Pipeline

  The client manages all aspects of the pipeline.

  Scenario: Create a pipeline
    When we create a new pipeline
    Then it should run successfully
    And I should see a log entry of "successfully created the pipeline" in the file "reports/pipeline.log"

  Scenario: Get a pipeline
    When we get a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully retrieved the pipeline" in the file "reports/pipeline.log"

  Scenario: Update a pipeline
    When we update a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully updated the pipeline" in the file "reports/pipeline.log"

  Scenario: Delete a pipeline
    When we delete a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully deleted the pipeline" in the file "reports/pipeline.log"

  Scenario: Trigger a pipeline
    When we trigger a pipeline
    Then it should run successfully
    And I should see a log entry of "successfully triggered the pipeline" in the file "reports/pipeline.log"

  Scenario: Missing token
    When we try to create a new pipeline with a missing token
    Then it should run unsuccessfully
    And I should see a log entry of "pipeline: open test: no such file or directory" in the file "reports/pipeline.log"
