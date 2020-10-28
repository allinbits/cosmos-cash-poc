Feature: Token Management
  As an Issuer
  I want to manage the token pool
  So that I can <do something useful>

  Scenario: Create token/issuer
    Given I am an "admin" called "validator"
    And I create an account with name "token-issuer"
    Then I create an issuer "token-issuer" with "100000000000" "tokens"
    And I wait for my transaction to be completed

  Scenario: Mint token
    Given I am an "issuer" called "token-issuer"
    And my account "token-issuer" has "100000000000" tokens
    When I create "100000000000" "tokens" tokens
    And I wait for my transaction to be completed
    Then my account "token-issuer" has "200000000000" tokens

  Scenario: Burn token
    Given I am an "issuer" called "token-issuer"
    And my account "token-issuer" has "200000000000" tokens
    When I burn "100000000000" "tokens" tokens
    And I wait for my transaction to be completed
    Then my account "token-issuer" has "100000000000" tokens
