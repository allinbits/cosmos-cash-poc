Feature: Token Management
  As an Issuer
  I want to manage the token pool
  So that I can <do something useful>

  Scenario: Create token/issuer
    Given I am an admin called "validator"
    And I create an account with name "token-issuer"
    Then I create an issuer "token-issuer" with "100000000000" "tokens"

    # TODO: update commands for minting and burning tokens 
    #And my account has 1000000 tokens
    #When I create 1000000 tokens
    #Then my account will have 2000000 tokens

