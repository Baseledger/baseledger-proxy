package types

//Put here our Types needed for the proxy elements?
type OffchainProcessMessageReferenceType string

type OffchainProcessMessage struct {
	senderID                         string
	receiverID                       string
	topic                            string
	offchainProcessMessageID         string
	referencedOffchainProcessMessage string
	referenceType                    OffchainProcessMessageReferenceType
	//BusinessObject todo
	hash                                 string
	blockchainTransactionIdOfStoredProof string
	baseledgerTransactionIdOfStoredProof string
}
