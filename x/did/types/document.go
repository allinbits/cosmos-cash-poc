package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
This represents a minimal self-managed did document
definition: https://w3c.github.io/did-core/

{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:example:123456789abcdefghi",
  "authentication": [{
    "id": "did:example:123456789abcdefghi#keys-1",
    "type": "Ed25519VerificationKey2018",
    "controller": "did:example:123456789abcdefghi",
    "publicKeyBase58": "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
  }],
  "service": [{
    "id":"did:example:123456789abcdefghi#vcs",
    "type": "VerifiableCredentialService",
    "serviceEndpoint": "https://example.com/vc/"
  }]
}
*/

// DidDocument is the data model defined by w3c
type DidDocument struct {
	Context        string   `json:"@context"`
	ID             string   `json:"id"`
	Authentication PubKeys  `json:"authentication"`
	Service        Services `json:"service,omitempty"`
}

func NewDidDocument(context string, id string, authentication PubKeys, services Services) DidDocument {
	return DidDocument{
		Context:        context,
		ID:             id,
		Authentication: authentication,
		Service:        services,
	}
}

type PubKeys []PubKey

// PubKey contains the information of a public key contained inside a Did Document
type PubKey struct {
	ID              string         `json:"id"`
	Type            string         `json:"type"`
	Controller      sdk.AccAddress `json:"controller"`
	PublicKeyBase58 string         `json:"publicKeyBase58"`
}

func NewPubKey(id string, pubKeyType string, controller sdk.AccAddress, base58Value string) PubKey {
	return PubKey{
		ID:              id,
		Type:            pubKeyType,
		Controller:      controller,
		PublicKeyBase58: base58Value,
	}
}

type Services []Service

// Service represents a service type needed for DidDocument
type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

func NewService(id string, serviceType string, serviceEndpoint string) Service {
	return Service{
		ID:              id,
		Type:            serviceType,
		ServiceEndpoint: serviceEndpoint,
	}
}
