Feature: Token Management
  As an Issuer
  I want to manage the token pool
  So that I can <do something useful>

  Scenario: Create tokens
    Given I am an issuer
    And my account has 1000000 tokens
    When I create 1000000 tokens
    Then my account will have 2000000 tokens
