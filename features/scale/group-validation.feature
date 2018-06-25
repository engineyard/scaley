Feature: Group Validation
  In order to ensure proper scaling behavior, we need to also validate that
  the group configuration is sound

  Background:
    Given I have a scaley config
    And I have a group named mygroup
    And I have a script that determines if I should scale up or down

    @failure
  Scenario Outline: Group lacks a scaling script
    Given there is capacity for the group to <Scaling Direction>
    And my group is configured to use the <Strategy> strategy
    But my group lacks a scaling script
    When I run `scaley scale mygroup`
    Then I see an error about the group's missing scaling script
    And it exits with an error

    Examples:
      | Scaling Direction | Strategy    |
      | upscale           | individual  |
      | upscale           | legion      |
      | downscale         | individual  |
      | downscale         | legion      |

    @failure
  Scenario Outline: Group specifies a non-existent scaling script
    Given there is capacity for the group to <Scaling Direction>
    And my group is configured to use the <Strategy> strategy
    But the scaling script for my group does not exist
    When I run `scaley scale mygroup`
    Then I see an error about the non-existent scaling script
    And it exits with an error

    Examples:
      | Scaling Direction | Strategy    |
      | upscale           | individual  |
      | upscale           | legion      |
      | downscale         | individual  |
      | downscale         | legion      |

    @failure
  Scenario Outline: Group has no scaling servers
    Given there is capacity for the group to <Scaling Direction>
    And my group is configured to use the <Strategy> strategy
    But my group has no scaling servers
    When I run `scaley scale mygroup`
    Then I see an error about the missing scaling servers
    And it exits with an error

    Examples:
      | Scaling Direction | Strategy    |
      | upscale           | individual  |
      | upscale           | legion      |
      | downscale         | individual  |
      | downscale         | legion      |

    @failure
  Scenario Outline: Group contains invalid scaling servers
    Given there is capacity for the group to <Scaling Direction>
    And my group is configured to use the <Strategy> strategy
    But my group has a scaling server that doesn't exist
    When I run `scaley scale mygroup`
    Then I see an error about the invalid scaling server
    And it exits with an error

    Examples:
      | Scaling Direction | Strategy    |
      | upscale           | individual  |
      | upscale           | legion      |
      | downscale         | individual  |
      | downscale         | legion      |

