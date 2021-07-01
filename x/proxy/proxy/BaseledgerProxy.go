package proxy

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/x/proxy/messaging"
	"github.com/unibrightio/baseledger/x/proxy/types"
	"github.com/unibrightio/baseledger/x/proxy/workgroups"

	// "github.com/cosmos/cosmos-sdk/client/tx"
	common "github.com/unibrightio/baseledger/common"
)

type workgroupMock struct {
	BaselineWorkgroupID uuid.UUID
	Description         string
	PrivatizeKey        string
}

type IBaseledgerProxy interface {
	CreateBaseledgerTransactionPayload(synchronizationRequest *types.SynchronizationRequest) (string, string)
	SendOffchainProcessMessage(message types.OffchainProcessMessage, txHash string)
}

type BaseledgerProxy struct {
	config          BaseledgerProxyConfig
	messagingClient messaging.IMessagingClient
	workgroupClient workgroups.IWorkgroupClient
}

type BaseledgerProxyConfig struct {
	connectionString string
}

func NewBaseledgerProxy() BaseledgerProxy {
	proxy := BaseledgerProxy{}
	proxy.config = BaseledgerProxyConfig{"das connection string"}

	proxy.messagingClient = &messaging.NatsMessagingClient{}
	// proxy.messagingClient.Subscribe("local server conn string", "token", "baseledger", receiveOffchainProcessMessage)

	proxy.workgroupClient = &workgroups.PostgresWorkgroupClient{}

	return proxy
}

func CreateBaseledgerTransactionPayload(
	synchronizationRequest *types.SynchronizationRequest,
	offchainProcessMessage *types.OffchainProcessMessage,
) string {
	workgroup := findWorkgroupMock(synchronizationRequest.WorkgroupId)
	// workgroup := workgroupClient.FindWorkgroup(synchronizationRequest.WorkgroupId)

	payload := &types.BaseledgerTransactionPayload{
		// TODO proper identifier BAS-33
		SenderId:                             "123",
		TransactionType:                      "Suggest",
		OffchainMessageId:                    offchainProcessMessage.Id.String(),
		ReferencedOffchainMessageId:          offchainProcessMessage.ReferencedOffchainProcessMessageId.String(),
		ReferencedBaseledgerTransactionId:    synchronizationRequest.ReferencedBaseledgerTransactionId,
		BaseledgerTransactionId:              offchainProcessMessage.BaseledgerTransactionIdOfStoredProof.String(),
		Proof:                                offchainProcessMessage.BusinessObjectProof,
		BaseledgerBusinessObjectId:           synchronizationRequest.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: synchronizationRequest.ReferencedBaseledgerBusinessObjectId,
	}

	fmt.Printf("\n payload %v \n", *payload)
	enc := privatizePayload(payload, workgroup.PrivatizeKey)
	fmt.Printf("enc %s\n\n", enc)
	dec := deprivatizePayload(enc, workgroup.PrivatizeKey)
	fmt.Printf("dec %s\n", dec)

	return enc
}

func CreateBaseledgerTransactionFeedbackPayload(
	synchronizationFeedback *types.SynchronizationFeedback,
	offchainProcessMessage *types.OffchainProcessMessage,
) string {
	workgroup := findWorkgroupMock(synchronizationFeedback.WorkgroupId)
	// workgroup := workgroupClient.FindWorkgroup(synchronizationRequest.WorkgroupId)

	feedbackMsg := "Approve"
	if !synchronizationFeedback.Approved {
		feedbackMsg = "Reject"
	}
	payload := &types.BaseledgerTransactionPayload{
		// TODO proper identifier BAS-33
		SenderId:                             "123",
		TransactionType:                      feedbackMsg,
		OffchainMessageId:                    offchainProcessMessage.Id.String(),
		ReferencedOffchainMessageId:          offchainProcessMessage.ReferencedOffchainProcessMessageId.String(),
		ReferencedBaseledgerTransactionId:    synchronizationFeedback.OriginalBaseledgerTransactionId,
		BaseledgerTransactionId:              offchainProcessMessage.BaseledgerTransactionIdOfStoredProof.String(),
		Proof:                                offchainProcessMessage.BusinessObjectProof,
		BaseledgerBusinessObjectId:           offchainProcessMessage.BaseledgerBusinessObjectId.String(),
		ReferencedBaseledgerBusinessObjectId: offchainProcessMessage.ReferencedBaseledgerBusinessObjectId.String(),
	}

	fmt.Printf("\n payload %v \n", *payload)
	enc := privatizePayload(payload, workgroup.PrivatizeKey)
	fmt.Printf("enc %s\n\n", enc)
	dec := deprivatizePayload(enc, workgroup.PrivatizeKey)
	fmt.Printf("dec %s\n", dec)

	return enc
}

func OffchainProcessMessageReceived(offchainProcessMessage types.OffchainProcessMessage, txHash string) {
	fmt.Println("OffchainProcessMessageReceived")
	entryType := common.SuggestionReceivedTrustmeshEntryType
	if offchainProcessMessage.EntryType == common.FeedbackSentTrustmeshEntryType {
		entryType = common.FeedbackReceivedTrustmeshEntryType
	}
	trustmeshEntry := &types.TrustmeshEntry{
		TendermintTransactionId:              offchainProcessMessage.BaseledgerTransactionIdOfStoredProof,
		OffchainProcessMessageId:             offchainProcessMessage.Id,
		SenderOrgId:                          offchainProcessMessage.SenderId,
		ReceiverOrgId:                        offchainProcessMessage.ReceiverId,
		WorkgroupId:                          uuid.FromStringOrNil(offchainProcessMessage.Topic),
		WorkstepType:                         offchainProcessMessage.WorkstepType,
		BaseledgerTransactionType:            offchainProcessMessage.BaseledgerTransactionType,
		BaseledgerTransactionId:              offchainProcessMessage.BaseledgerTransactionIdOfStoredProof,
		ReferencedBaseledgerTransactionId:    offchainProcessMessage.ReferencedBaseledgerTransactionId,
		BusinessObjectType:                   offchainProcessMessage.BusinessObjectType,
		BaseledgerBusinessObjectId:           offchainProcessMessage.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: offchainProcessMessage.ReferencedBaseledgerBusinessObjectId,
		ReferencedProcessMessageId:           offchainProcessMessage.ReferencedOffchainProcessMessageId,
		TransactionHash:                      txHash,
		EntryType:                            entryType,
	}

	if !trustmeshEntry.Create() {
		fmt.Printf("error when creating new trustmesh entry")
	}
}

// TODO: skos remove this and read from db
func findWorkgroupMock(workgroupId uuid.UUID) *workgroupMock {
	return &workgroupMock{
		BaselineWorkgroupID: workgroupId,
		Description:         "Mocked workgroup",
		PrivatizeKey:        "0c2e08bc9249fb42568e5a478e9af87a208471c46211a08f3ad9f0c5dbf57314",
	}
}

// TODO: made this public just as a mock, we will integrate with NATS here and implement real logic
func SendOffchainProcessMessage(message types.OffchainProcessMessage, receiver string, txHash string) {
	fmt.Printf("SENDING OFFCHAIN PROCESS MESSAGE WITH ID %v AND TX HASH %v\n", message.Id, txHash)
	// marshal natsMessage to byte array
	// recipientMessagingEndpoint := workgroupClient.FindRecipientMessagingEndpoint(recipientId)
	// recipientMessagingToken := workgroupClient.FindRecipientMessagingToken(recipientId)
	// messagingClient.SendMessage("TODO: convert message to correct payload", recipientMessagingEndpoint, recipientMessagingToken)
}

// func receiveOffchainProcessMessage(sender string, message string) {
// 	fmt.Printf("\n sender %v \n", sender)
// 	fmt.Printf("\n message %v \n", message)
// }

func CreateHashFromBusinessObject(bo string) string {
	hash := md5.Sum([]byte(bo))
	return hex.EncodeToString(hash[:])
}

func DeprivatizeBaseledgerTransactionPayload(payload string, workgroupId uuid.UUID) string {
	workgroup := findWorkgroupMock(workgroupId)
	return deprivatizePayload(payload, workgroup.PrivatizeKey)
}

func privatizePayload(payload *types.BaseledgerTransactionPayload, key string) string {
	payloadJson, _ := json.Marshal(payload)
	fmt.Println("json", string(payloadJson))
	return encrypt(string(payloadJson), key)
}

func deprivatizePayload(payload string, key string) string {
	return decrypt(payload, key)
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(encryptedString string, keyString string) (decryptedString string) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
