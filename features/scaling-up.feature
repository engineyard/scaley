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
    And it exits successfully

  Scenario: Scaling with insufficient capacity
    Given there is not capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then a warning is logged regarding the insufficient capacity
    And it exits successfully
    But no changes are made

    @failure
  Scenario: Attempting to upscale while a scaling event is in progress
    Given a scaling lockfile exists for the group
    And there is capacity for the group to upscale
    When I run `scaley scale mygroup`
    Then it exits with an error
    And no changes are made
