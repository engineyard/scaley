Feature: Scaling Down
  I have a pool of instances that I want to use for scaling some aspect of
  my app. I don't want them to always be turned up, though, because that
  would be some expensive overkill. So, I want to be able to scale the pool
  down when it makes sense to do so.

  Background:
    Given I have a group named mygroup
    And I have a script that determines if I should scale up or down
    And conditions dictate that downscaling is necessary

  Scenario: Scaling with sufficient capacity
    Given there is capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then the group is scaled down

  Scenario: Scaling with insufficient capacity
    Given there is not capacity for the group to downscale
    When I run `scaley scale mygroup`
    Then no changes are made
    And no messages are logged

  Scenario: Attempting to downscale while a scaling event is in progress
    Given a scaling lockfile exists for the group
    When I run `scaley scale mygroup`
    Then no changes are made
