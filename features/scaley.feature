Feature: Scaley
  What is the behavior when I run scaley without a command?

  Scenario: Running scaley
    When I run `scaley`
    Then I see the help description
    Then I see the usage
    Then I see the available commands
    And it exits successfully
