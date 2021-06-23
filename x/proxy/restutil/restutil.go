package restutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	txutil "github.com/unibrightio/baseledger/txutil"
	types "github.com/unibrightio/baseledger/x/proxy/types"
)

// need to create these because amino json validation
type baseReq struct {
	From    string `json:"from"`
	ChainId string `json:"chain_id"`
}

// keep in sync with handler struct
type feedbackReq struct {
	BaseReq                                    baseReq `json:"base_req"`
	WorkgroupId                                string  `json:"workgroup_id"`
	BusinessObjectType                         string  `json:"business_object_type"`
	Recipient                                  string  `json:"recipient"`
	Approved                                   bool    `json:"approved"`
	BaseledgerBusinessObjectIdOfApprovedObject string  `json:"baseledger_business_object_id_of_approved_object"`
	HashOfObjectToApprove                      string  `json:"hash_of_object_to_approve"`
	OriginalBaseledgerTransactionId            string  `json:"original_baseledger_transaction_id"`
	OriginalOffchainProcessMessageId           string  `json:"original_offchain_process_message_id"`
	FeedbackMessage                            string  `json:"feedback_message"`
	BaseledgerProvenBusinessObjectJson         string  `json:"baseledger_proven_business_object_json"`
}

func RejectFeedback(offchainMessage types.OffchainProcessMessage, workgroupId string) {
	fromAddress := getFromAddress()

	feedback := feedbackReq{
		BaseReq: baseReq{
			From:    fromAddress,
			ChainId: "baseledger",
		},
		FeedbackMessage:                    "Rejected because Hashes do not match",
		Approved:                           false,
		BaseledgerProvenBusinessObjectJson: offchainMessage.BaseledgerSyncTreeJson,
		BaseledgerBusinessObjectIdOfApprovedObject: offchainMessage.BaseledgerBusinessObjectId.String(),
		WorkgroupId: workgroupId,
		Recipient:   offchainMessage.SenderId.String(),
		// TODO: Which proof to send here?
		HashOfObjectToApprove:            offchainMessage.BusinessObjectProof,
		OriginalBaseledgerTransactionId:  offchainMessage.BaseledgerTransactionIdOfStoredProof.String(),
		OriginalOffchainProcessMessageId: offchainMessage.Id.String(),
		BusinessObjectType:               offchainMessage.BusinessObjectType,
	}

	jsonValue, err := json.Marshal(feedback)

	if err != nil {
		fmt.Println("Error marshaling json feedback")
		return
	}

	_, err = http.Post("http://localhost:1317/proxy/feedback", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("error while sending feedback request %v\n", err.Error())
	}
}

func getFromAddress() string {
	kr, err := txutil.NewKeyringInstance()
	if err != nil {
		panic("Keyring not found")
	}

	keysList, err := kr.List()
	if err != nil || len(keysList) == 0 {
		panic("Keyring keys list empty")
	}

	return keysList[0].GetAddress().String()
}
