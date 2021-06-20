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
)

type workgroupMock struct {
	BaselineWorkgroupID string
	Description         string
	PrivatizeKey        string
}

type IBaseledgerProxy interface {
	CreateBaseledgerTransactionPayload(synchronizationRequest *types.SynchronizationRequest) (string, string)
	SendOffchainProcessMessage(message types.OffchainProcessMessage, recipientId string)
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
	proxy.messagingClient.Subscribe("local server conn string", "baseledger", receiveOffchainProcessMessage)

	proxy.workgroupClient = &workgroups.PostgresWorkgroupClient{}

	return proxy
}

func CreateBaseledgerTransactionPayload(synchronizationRequest *types.SynchronizationRequest) (string, string) {
	hash := CreateHashFromBusinessObject(synchronizationRequest.BusinessObject)
	offchainProcessMessage := newOffchainProcessMessage(
		synchronizationRequest.WorkstepType,
		"",
		synchronizationRequest.BusinessObject,
		hash,
		synchronizationRequest.BaseledgerBusinessObjectId,
		synchronizationRequest.ReferencedBaseledgerBusinessObjectId,
		synchronizationRequest.WorkstepType+" suggested")

	workgroup := findWorkgroupMock(synchronizationRequest.WorkgroupId)
	// workgroup := workgroupClient.FindWorkgroup(synchronizationRequest.WorkgroupId)

	transactionIdUuid, _ := uuid.NewV4()
	payload := &types.BaseledgerTransactionPayload{
		// TODO proper identifier
		PhonebookIdentifier:                  "123",
		TransactionType:                      "Suggest",
		OffchainMessageId:                    offchainProcessMessage.Id.String(),
		ReferencedOffchainMessageId:          "",
		ReferencedBaseledgerTransactionId:    "",
		BaseledgerTransactionID:              transactionIdUuid.String(),
		Proof:                                hash,
		BaseledgerBusinessObjectId:           synchronizationRequest.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: synchronizationRequest.ReferencedBaseledgerBusinessObjectId,
	}

	fmt.Printf("\n payload %v \n", *payload)
	enc := privatizePayload(payload, workgroup.PrivatizeKey)
	fmt.Printf("enc %s\n\n", enc)
	dec := deprivatizePayload(enc, workgroup.PrivatizeKey)
	fmt.Printf("dec %s\n", dec)

	return enc, transactionIdUuid.String()
}

func OffchainProcessMessageReceived(processMessage types.OffchainProcessMessage) {
	fmt.Println("OffchainProcessMessageReceived")
}

func getBaseledgerTx(id string) {

}

func findWorkgroupMock(workgroupId string) *workgroupMock {
	newUuid, _ := uuid.NewV4()
	return &workgroupMock{
		BaselineWorkgroupID: newUuid.String(),
		Description:         "Mocked workgroup",
		PrivatizeKey:        "0c2e08bc9249fb42568e5a478e9af87a208471c46211a08f3ad9f0c5dbf57314",
	}
}

// TODO: made this public just as a mock, we will integrate with NATS here and implement real logic
func SendOffchainProcessMessage(message types.OffchainProcessMessage, recipientId string) {
	fmt.Printf("SENDING OFFCHAIN PROCESS MESSAGE WITH ID %v\n", message.Id)
	// recipientMessagingEndpoint := workgroupClient.FindRecipientMessagingEndpoint(recipientId)
	// recipientMessagingToken := workgroupClient.FindRecipientMessagingToken(recipientId)
	// messagingClient.SendMessage("TODO: convert message to correct payload", recipientMessagingEndpoint, recipientMessagingToken)
}

func receiveOffchainProcessMessage(sender string, message string) {
	fmt.Printf("\n sender %v \n", sender)
	fmt.Printf("\n message %v \n", message)
}

func newOffchainProcessMessage(
	workstepType string,
	referencedOffchainProcessMessage string,
	businessObject string,
	hashOfBusinessObject string,
	baseledgerBusinessObjectID string,
	referencedBaseledgerBusinessObjectID string,
	statusTextMessage string) *types.OffchainProcessMessage {
	newUuid, _ := uuid.NewV4()
	return &types.OffchainProcessMessage{
		Id:                                   newUuid,
		WorkstepType:                         workstepType,
		ReferencedOffchainProcessMessageId:   referencedOffchainProcessMessage,
		Hash:                                 hashOfBusinessObject,
		BusinessObject:                       businessObject,
		BaseledgerBusinessObjectId:           baseledgerBusinessObjectID,
		ReferencedBaseledgerBusinessObjectId: referencedBaseledgerBusinessObjectID,
		StatusTextMessage:                    statusTextMessage,
	}
}

// TODO: currently it assumes it is json string, refactor this
func CreateHashFromBusinessObject(bo string) string {
	hash := md5.Sum([]byte(bo))
	return hex.EncodeToString(hash[:])
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
