Feature: Pipeline

  The client manages all aspects of the pipeline.

  Scenario: Create a pipeline
    When we create a new pipeline
    Then it should run successfully
    And I should see a log entry of "successfully created pipeline" in the file "reports/pipeline.log"
