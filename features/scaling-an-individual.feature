Feature: Scaling An Individual
  For some workflows, it makes sense to increase/decrease the servers
  involved one-at-a-time when there's work. This is how the individual
  strategy works. When a scaling event occurs, only one of the servers
  that can scale in the desired direction is brought up or down,
  depending on which operation is requested by the scaling script.

  Background:
    Given I have a scaley config
    And I have a group named mygroup
    And I have a script that determines if I should scale up or down
    And my group is configured to use the individual strategy

  Scenario: Scaling Up
    Given conditions dictate that upscaling is necessary
    And there is capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then only one server in the group is started

  Scenario: Scaling Down
    Given conditions dictate that downscaling is necessary
    And there is capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then only one server in the group is stopped
