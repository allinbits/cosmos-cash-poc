# State

## Validator

---

Validators objects should be primarily stored and accessed by the `Name`, or `Address` an SDK validator address for the operator of the validator.

- Validators: `0x51 | Name -> amino(Validator)`

```go
type Validator struct {
    Name	     string
    Address	     sdk.ValAddress
    PubKey	     crypto.PubKey
    Jailed           bool
    Description      Description
}

type Description struct {
    Moniker          string                 // name
    Identity         string                 // optional identity signature
    Website          string
    SecurityContact  string
    Details          string
}
```

## Vote

---

Vote objects should be primarily stored and accessed by the `CandidateAddress` & `VoterAddress` to allows the application to handle duplication votes

- Vote: `0x52 | VoterAddr | CandidateAddr -> amino(Vote)`

Each vote state is stored in a Vote struct:

```go
type Vote struct {
    Voter         sdk.ValAddress
    Name          uint32        // Name | Candidate of the v
    InFavor       bool          // Has voted in favor of a new validator
}
```
