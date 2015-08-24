package segment

import "github.com/mitch000001/go-hbci/domain"

var SupportedHBCIVersions = map[int]HBCIVersion{
	220: HBCI220,
	300: FINTS300,
}

type HBCIVersion struct {
	version                   int
	PinTanEncryptionHeader    func(clientSystemId string, keyName domain.KeyName) *EncryptionHeaderSegment
	RDHEncryptionHeader       func(clientSystemId string, keyName domain.KeyName, key []byte) *EncryptionHeaderSegment
	SignatureHeader           func() *SignatureHeaderSegment
	PinTanSignatureHeader     func(controlReference string, clientSystemId string, keyName domain.KeyName) *SignatureHeaderSegment
	RDHSignatureHeader        func(controlReference string, signatureId int, clientSystemId string, keyName domain.KeyName) *SignatureHeaderSegment
	SignatureEnd              func() *SignatureEndSegment
	SynchronisationRequest    func(modus int) *SynchronisationRequestSegment
	AccountTransactionRequest func(account domain.AccountConnection, allAccounts bool) *AccountTransactionRequestSegment
}

func (v HBCIVersion) Version() int {
	return v.version
}
