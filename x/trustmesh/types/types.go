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

// all other types for hasing, privacy, off-chain messaging

//Have here the one interface for all proxy methods hasing, privacy, off-chain messaging (IBaseledgerProxy)
//We will implement it with our BaseledgerProxy within this project and open it up for other implementations like Provide's to fulfill the interface as well.
//Assumption is, this would need a component within this project to make RESt/gRPC calls to Provide's or others solutions.
type IBaseledgerProxyInterface interface {
	privatize(text string) string
	deprivatize(textEncrypted string) string

	hash(payload string) string

	//baselineBusinessObjectInitially(businessObject o)
	//giveFeddbackToBusinessobject()
	//......

	//workgroup methods..

	//off-chain messaging methods
}
