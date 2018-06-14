Feature: Scaling A Legion
  For some workflows, it makes sense to invite the whole team to the
  party when there's work to be done. This is how the legion strategy
  works. When a scaling event occurs, all of the scaling servers in
  the group are brought either up or down, depending on which operation
  is requested by the scaling script.

  Background:
    Given I have a scaley config
    And I have a group named mygroup
    And I have a script that determines if I should scale up or down
    And my group is configured to use the legion strategy

  Scenario: Scaling Up
    Given conditions dictate that upscaling is necessary
    And there is capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then all of the servers in the group are started

  Scenario: Scaling Down
    Given conditions dictate that downscaling is necessary
    And there is capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then all of the servers in the group are stopped
