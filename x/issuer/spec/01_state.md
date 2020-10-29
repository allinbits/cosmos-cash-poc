# State

## Issuer

---

Issuers objects should be primarily stored and accessed by the `Name`, or `Token`.

- Issuers: `0x51 | Name -> amino(Issuer)`
- Issuers: `0x52 | Token -> amino(Issuer)`

```go
type Issuer struct {
	Name    string         `json:"name" yaml:"name"`
	Token   string         `json:"token" yaml:"token"`
	Fee     uint16         `json:"fee" yaml:"fee"`
	State   IssuerState    `json:"state" yaml:"state"`
	Address sdk.AccAddress `json:"address" yaml:"address"`
}
```

