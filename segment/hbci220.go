package segment

var HBCI220 = Version{
	version:                220,
	PinTanEncryptionHeader: NewPinTanEncryptionHeaderSegment,
	RDHEncryptionHeader:    NewEncryptionHeaderSegment,
	PinTanSignatureHeader:  NewPinTanSignatureHeaderSegment,
	RDHSignatureHeader:     NewRDHSignatureHeaderSegment,
	SignatureEnd:           NewSignatureEndSegment,
	SynchronisationRequest: NewSynchronisationSegmentV2,
}
