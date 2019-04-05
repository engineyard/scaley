Feature: Custom Stop Script
  So as to ensure that a server that I'm stopping during a downscale
  event is actually ready to be stopped, I'd like to configure a custom
  script that will connect to the target server to perform pre-shutdown
  processes.

  Background:
    Given I have a scaley config
    And I have a group named mygroup
    And I have a script that determines if I should scale up or down
    And conditions dictate that downscaling is necessary
    And there is capacity for the group to downscale

  Scenario: No stop script (default behavior)
    Given my group does not use a custom stop script
    When I run `scaley scale mygroup`
    Then the group is scaled down
    And it exits successfully
    And the group is unlocked
    But the stop script is not executed

  Scenario: With a custom stop script
    Given my group uses a custom stop script that always succeeds
    When I run `scaley scale mygroup`
    Then the stop script is executed for each target server
    And the group is scaled down
    And it exits successfully
    And the group is unlocked

  Scenario: With a crashy custom stop script
    Given my group uses a custom stop script that fails for the first server
    When I run `scaley scale mygroup`
    Then the stop script is executed for each target server
    But a stop script failure is logged for the first server
    And all applicable servers but the first server are stopped
    And it exits with an error
    And the group is unlocked

