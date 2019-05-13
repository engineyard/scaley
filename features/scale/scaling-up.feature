Feature: Scaling Up
  In order to ensure that I have all of the resources that I need for my app
  to run well, I would like to automatically scale those resources up as
  needed.

  Background:
    Given I have a scaley config
    And I have a group named mygroup
    And I have a script that determines if I should scale up or down
    And conditions dictate that upscaling is necessary

  Scenario: Scaling with sufficient capacity
    Given there is capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then the group is scaled up
    And the group is unlocked
    And it exits successfully

  Scenario Outline: Scaling with insufficient capacity
    Given my group is configured to use the <Strategy> strategy
    Given there is not capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then a warning is logged regarding the insufficient capacity
    And it exits successfully
    But no changes are made
    And the group is unlocked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario: Attempting to upscale while a scaling event is in progress
    Given a scaling lockfile exists for the group
    And there is capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a locking failure is logged
    But no changes are made
    And the group remains locked

    @failure
  Scenario Outline: Start server yields an invalid API response
    Given my group is configured to use the <Strategy> strategy
    Given there is capacity for the group to upscale
    But the API is erroring on server start requests
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a scaling failure is logged
    But the group is unlocked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario Outline: Server start failure
    Given my group is configured to use the <Strategy> strategy
    And there is capacity for the group to upscale
    But the servers cannot be started successfully
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a scaling failure is logged
    But the group is unlocked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario Outline: Chef run yields an API error
    Given my group is configured to use the <Strategy> strategy
    And there is capacity for the group to upscale
    But the API is erroring on environment configuration requests
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a chef failure is logged
    And the group remains locked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario Outline: Chef run failure
    Given my group is configured to use the <Strategy> strategy
    Given there is capacity for the group to upscale
    But the environment cannot run chef successfully
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a chef failure is logged
    And the group remains locked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |
