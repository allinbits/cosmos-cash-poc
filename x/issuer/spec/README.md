# issuer module specification

## Abstract

This paper specifies a `issuer` module of the Cosmos-SDK. The `issuer` module will define 

The `issuer` module will have the following characteristics:

- Issuers can be created by admins in the system
- Issuers can mint tokens 
- Issuers can burn tokens 
- Issuers can freeze accounts and freeze token transfers 

One main data structure is used to allow the `issuer` module to function correctly:

- `Issuer`: Issuer represents an actor in the system that can manage tokens

## Contents

1. **[State](01_state.md)**
1. **[State Transitions](02_state_transitions.md)**
3. **[Messages](03_messages.md)**
7. **[Events](04_events.md)**
