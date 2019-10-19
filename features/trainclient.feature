Feature: Train Client

Scenario: Get departure board for a station
    Given I have a Train Client with real API credentials
    And I have entered a valid station
    When I get the departure board
    Then I should receive a response
