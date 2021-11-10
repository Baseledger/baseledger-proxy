package proxyutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	uuid "github.com/kthomas/go.uuid"
	"github.com/spf13/viper"

	// "github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/messaging"
	"github.com/unibrightio/proxy-api/types"
	"github.com/unibrightio/proxy-api/workgroups"
)

type workgroupMock struct {
	BaselineWorkgroupID uuid.UUID
	Description         string
	PrivatizeKey        string
}

type IBaseledgerProxy interface {
	CreateBaseledgerTransactionPayload(synchronizationRequest *types.NewSuggestionRequest) (string, string)
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

func CreateNewSuggestionBaseledgerTransactionPayload(
	newSuggestionRequest *types.NewSuggestionRequest,
	offchainProcessMessage *types.OffchainProcessMessage,
) string {
	// Do we need client anymore now when we split apps? Maybe just simple query util like for other entities?
	// should we load workgroup at start and keep it in memory, we are querying it all the time and it won't change
	workgroupClient := &workgroups.PostgresWorkgroupClient{}
	workgroup := workgroupClient.FindWorkgroup(newSuggestionRequest.WorkgroupId.String())

	payload := &types.BaseledgerTransactionPayload{
		SenderId:                   offchainProcessMessage.SenderId.String(),
		TransactionType:            offchainProcessMessage.BaseledgerTransactionType,
		OffchainMessageId:          offchainProcessMessage.Id.String(),
		BaseledgerTransactionId:    offchainProcessMessage.BaseledgerTransactionIdOfStoredProof.String(),
		Proof:                      offchainProcessMessage.BusinessObjectProof,
		BaseledgerBusinessObjectId: newSuggestionRequest.BaseledgerBusinessObjectId,
	}

	logger.Infof("\n payload %v \n", *payload)
	enc := privatizePayload(payload, workgroup.PrivatizeKey)
	logger.Infof("enc %s\n\n", enc)
	dec := deprivatizePayload(enc, workgroup.PrivatizeKey)
	logger.Infof("dec %s\n", dec)

	return enc
}

func CreateExitBaseledgerTransactionPayload(
	workgroupId uuid.UUID,
	baseledgerTransactionId uuid.UUID,
	rootProof string,
) string {
	workgroupClient := &workgroups.PostgresWorkgroupClient{}
	workgroup := workgroupClient.FindWorkgroup(workgroupId.String())
	payload := &types.BaseledgerTransactionExitPayload{
		RootProof:               rootProof,
		BaseledgerTransactionId: baseledgerTransactionId.String(),
	}

	return privatizeExitPayload(payload, workgroup.PrivatizeKey)
}

func CreateNewFeedbackBaseledgerTransactionPayload(
	newFeedbackRequest *types.NewFeedbackRequest,
	offchainProcessMessage *types.OffchainProcessMessage,
) string {
	workgroupClient := &workgroups.PostgresWorkgroupClient{}
	workgroup := workgroupClient.FindWorkgroup(newFeedbackRequest.WorkgroupId.String())

	payload := &types.BaseledgerTransactionPayload{
		SenderId:                             viper.Get("ORGANIZATION_ID").(string),
		TransactionType:                      offchainProcessMessage.BaseledgerTransactionType,
		OffchainMessageId:                    offchainProcessMessage.Id.String(),
		ReferencedBaseledgerTransactionId:    newFeedbackRequest.OriginalBaseledgerTransactionId,
		BaseledgerTransactionId:              offchainProcessMessage.BaseledgerTransactionIdOfStoredProof.String(),
		Proof:                                offchainProcessMessage.BusinessObjectProof,
		BaseledgerBusinessObjectId:           offchainProcessMessage.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: offchainProcessMessage.ReferencedBaseledgerBusinessObjectId,
	}

	logger.Infof("\n payload %v \n", *payload)
	enc := privatizePayload(payload, workgroup.PrivatizeKey)
	logger.Infof("enc %s\n\n", enc)
	dec := deprivatizePayload(enc, workgroup.PrivatizeKey)
	logger.Infof("dec %s\n", dec)

	return enc
}

func SendOffchainMessage(payload []byte, workgroupId string, recipientId string) (err error) {
	workgroupClient := &workgroups.PostgresWorkgroupClient{}

	logger.Infof("trying to find workgroup member - workgroup id: %s recipient id: %s \n", workgroupId, recipientId)
	workgroupMembership := workgroupClient.FindWorkgroupMember(workgroupId, recipientId)

	if workgroupMembership == nil {
		return errors.New("failed to find a workgroup member")
	}

	logger.Infof("trying to message on url: %s with token: %s\n", workgroupMembership.OrganizationEndpoint, workgroupMembership.OrganizationToken)

	messagingClient := &messaging.NatsMessagingClient{}
	messagingClient.SendMessage(payload, workgroupMembership.OrganizationEndpoint, workgroupMembership.OrganizationToken)

	return nil
}

func CreateHashFromBusinessObject(bo string) string {
	hash := md5.Sum([]byte(bo))
	return hex.EncodeToString(hash[:])
}

func DeprivatizeBaseledgerTransactionPayload(payload string, workgroupId uuid.UUID) string {
	workgroupClient := &workgroups.PostgresWorkgroupClient{}
	workgroup := workgroupClient.FindWorkgroup(workgroupId.String())
	return deprivatizePayload(payload, workgroup.PrivatizeKey)
}

func privatizePayload(payload *types.BaseledgerTransactionPayload, key string) string {
	payloadJson, _ := json.Marshal(payload)
	return encrypt(string(payloadJson), key)
}

func privatizeExitPayload(payload *types.BaseledgerTransactionExitPayload, key string) string {
	payloadJson, _ := json.Marshal(payload)
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
