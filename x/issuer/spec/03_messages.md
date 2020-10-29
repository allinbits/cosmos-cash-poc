# Messages

In this section we describe the processing of the issuer messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./01_state.md) section.

## **MsgCreateIssuer**

---

A issuer is created using the `MsgCreateIssuer` message

```go
type MsgCreateIssuer struct {
	Name    string         `json:"name"`
	Address sdk.AccAddress `json:"issuer"`
	Token   string         `json:"token"`
	Amount  string         `json:"amount"`
	Owner   sdk.AccAddress `json:"owner"`
}
```


## **MsgBurnToken**

---

A `Token` can be burned as an issuer using the `MsgBurnToken` message.

```go
type MsgBurnToken struct {
	Token  string         `json:"token"`
	Amount string         `json:"amount"`
	Issuer sdk.AccAddress `json:"issuer"`
}

```

## **MsgMintToken**

---

A `Token` can be burned as an issuer using the `MsgMintToken` message.

```go
type MsgMintToken struct {
	Token  string         `json:"token"`
	Amount string         `json:"amount"`
	Issuer sdk.AccAddress `json:"issuer"`
}

```

## **MsgFreezeToken**

---

A `Token` can be frozen as an issuer using the `MsgFreezeToken` message.
 
```go
type MsgMintToken struct {
	Token  string         `json:"token"`
	Issuer sdk.AccAddress `json:"issuer"`
}
```
