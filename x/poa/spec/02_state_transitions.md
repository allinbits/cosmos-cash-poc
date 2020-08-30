# State Transitions

This document describes the state transition operations pertaining to:

1. [Validators](./02_state_transitions.md#validators)
2. [Votes](./02_state_transitions.md#votes)

## Validators

State transitions in validators are performed on every `EndBlock` and are used update the active `ValidatorSet`.

### `Accepted` **||** `Not Accepted`

When a validator is accepted it is added the Tendermint set. The following operations occur:

- set `Validator.Accepted` to true
- update the `Validator` object for this validator

## Votes

Vote counts are performed on every `BeginBlock` and are used update the active `ValidatorSet`.

