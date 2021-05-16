package types

//Put here our Types needed for the proxy elements?
type OffchainProcessMessageReferenceType string

type OffchainProcessMessage struct {
	SenderID                         string
	ReceiverID                       string
	Topic                            string
	OffchainProcessMessageID         string
	ReferencedOffchainProcessMessage string
	ReferenceType                    OffchainProcessMessageReferenceType
	// todo replace string with proper type?
	BusinessObject                       string
	WorkstepType                         string
	Hash                                 string
	BlockchainTransactionIdOfStoredProof string
	BaseledgerTransactionIdOfStoredProof string
	BaseledgerBusinessObjectID           string
	ReferencedBaseledgerBusinessObjectID string
	StatusTextMessage                    string
}

// TODO rename after clean up
type SynchronizationRequest struct {
	WorkgroupID                          string
	Recipient                            string
	WorkstepType                         string
	BusinessObjectType                   string
	BaseledgerBusinessObjectID           string
	BusinessObject                       string
	ReferencedBaseledgerBusinessObjectID string
}

type BaseledgerTransactionPayload struct {
	PhonebookIdentifier                  string `json:"phonebookIdentifier"`
	BaseledgerTransactionType            string `json:"baseledgerTransactionType"`
	OffchainMessageId                    string `json:"offchainMessageId"`
	ReferencedOffchainMessageId          string `json:"referencedOffchainMessageId"`
	ReferencedBaseledgerTransactionId    string `json:"referencedBaseledgerTransactionId"`
	BaseledgerTransactionID              string `json:"baseledgerTransactionID"`
	Proof                                string `json:"proof"`
	BaseledgerBusinessObjectID           string `json:"baseledgerBusinessObjectID"`
	ReferencedBaseledgerBusinessObjectID string `json:"referencedBaseledgerBusinessObjectID"`
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
