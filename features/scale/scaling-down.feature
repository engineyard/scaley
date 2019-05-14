Feature: Scaling Down
  I have a pool of instances that I want to use for scaling some aspect of
  my app. I don't want them to always be turned up, though, because that
  would be some expensive overkill. So, I want to be able to scale the pool
  down when it makes sense to do so.

  Background:
    Given I have a scaley config
    And I have a group named mygroup
    And I have a script that determines if I should scale up or down
    And conditions dictate that downscaling is necessary

  Scenario: Scaling with sufficient capacity
    Given there is capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then the group is scaled down
    And the group is unlocked

  Scenario Outline: Scaling with insufficient capacity
    Given my group is configured to use the <Strategy> strategy
    And there is not capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then no changes are made
    And no messages are logged
    And the group is unlocked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario: Attempting to downscale while a scaling event is in progress
    Given a scaling lockfile exists for the group
    And there is capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a locking failure is logged
    But no changes are made
    And the group remains locked

    @failure
  Scenario Outline: Stop server yields an invalid API response
    Given my group is configured to use the <Strategy> strategy
    And there is capacity for the group to downscale
    But the API is erroring on server stop requests
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a scaling failure is logged
    But the group is unlocked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario Outline: Server stop failure
    Given my group is configured to use the <Strategy> strategy
    And there is capacity for the group to downscale
    But the servers cannot be stopped successfully
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
    And there is capacity for the group to downscale
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
    Given there is capacity for the group to downscale
    But the environment cannot run chef successfully
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a chef failure is logged
    And the group remains locked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

    @failure
  Scenario Outline: Chef run failure (locking_on_failure disabled)
    Given my group is configured to use the <Strategy> strategy
    And my group is configured to unlock on failures
    And there is capacity for the group to downscale
    But the environment cannot run chef successfully
    When I run `scaley scale mygroup`
    Then it exits with an error
    And a chef failure is logged
    But the group is unlocked

    Examples:
      | Strategy    |
      | individual  |
      | legion      |

