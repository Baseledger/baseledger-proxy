package systemofrecord

import (
	"fmt"
	"strings"
	"testing"

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/types"
)

func TestGivenATrutmeshUrlWithAtTrustmeshParamWhenBuildWebhookRequestUrlParamCorrectlyInjected(t *testing.T) {
	trustmeshEntry := &types.TrustmeshEntry{
		BaseledgerBusinessObjectId: "test",
		BusinessObjectType:         "test",
		SorBusinessObjectId:        "test",
		BaseledgerTransactionId:    uuid.NewV4(),
		ReceiverOrgId:              uuid.NewV4(),
		TrustmeshId:                uuid.NewV4(),
	}

	want := "https://test.url/" + trustmeshEntry.TrustmeshId.String()

	result := buildWebhookRequestUrl("https://test.url/{{trustmesh_id}}", "trustmesh_id:TrustmeshId", trustmeshEntry)
	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}

func TestGivenAUrlWhenhandleBasicAuthBasicAuthAddedTourl(t *testing.T) {
	username := "testuser"
	password := "testpass"
	targetUrl := "https://www.test.url/"

	want := fmt.Sprintf("https://%s:%s@www.test.url/", username, password)

	result := handleBasicAuth(targetUrl, username, password)

	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}

func TestGivenACreateWebhookWithSpecialParamPayloadAndOriginWhenBuildWebhookRequestBodyCorrectRequestBodyCreated(t *testing.T) {
	trustmeshEntry := &types.TrustmeshEntry{
		BaseledgerBusinessObjectId: "test",
		BusinessObjectType:         "test",
		SorBusinessObjectId:        "test",
		BaseledgerTransactionId:    uuid.NewV4(),
		ReceiverOrgId:              uuid.NewV4(),
		TrustmeshId:                uuid.NewV4(),
	}

	webhook := &types.SorWebhook{
		Body:        "{ 'sorPayload': '{{business_object_json_payload}}', 'origin':'{{origin}}' }",
		BodyParams:  "",
		WebhookType: types.CreateObject,
	}

	payload := `{ "testproperty" : "testvalue" }`

	want := "{ 'sorPayload': '{ \\\"testproperty\\\" : \\\"testvalue\\\" }', 'origin':'proxy' }"

	result := buildWebhookRequestBody(
		webhook.Body,
		webhook.BodyParams,
		webhook.WebhookType,
		trustmeshEntry,
		payload,
		"irelevant",
		"irelevant",
		"proxy",
	)

	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}

func TestGivenACreateWebhookWithBodyParamsWhenBuildWebhookRequestBodyCorrectRequestBodyCreated(t *testing.T) {
	trustmeshEntry := &types.TrustmeshEntry{
		BaseledgerBusinessObjectId: "test",
		BusinessObjectType:         "test",
		SorBusinessObjectId:        "test",
		BaseledgerTransactionId:    uuid.NewV4(),
		ReceiverOrgId:              uuid.NewV4(),
		TrustmeshId:                uuid.NewV4(),
	}

	webhook := &types.SorWebhook{
		Body:        "{ 'sorPayload': '{{business_object_json_payload}}', 'workflow_id':'{{trustmesh_id}}' }",
		BodyParams:  "trustmesh_id:TrustmeshId",
		WebhookType: types.CreateObject,
	}

	payload := `{ "testproperty" : "testvalue" }`

	want := "{ 'sorPayload': '{ \\\"testproperty\\\" : \\\"testvalue\\\" }', 'workflow_id':'" + trustmeshEntry.TrustmeshId.String() + "' }"

	result := buildWebhookRequestBody(
		webhook.Body,
		webhook.BodyParams,
		webhook.WebhookType,
		trustmeshEntry,
		payload,
		"irelevant",
		"irelevant",
		"irelevant",
	)

	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}

func TestGivenAUpdateWebhookWithSpecialParamsApprovedAndMessageWhenBuildWebhookRequestBodyCorrectRequestBodyCreated(t *testing.T) {
	trustmeshEntry := &types.TrustmeshEntry{
		BaseledgerBusinessObjectId: "test",
		BusinessObjectType:         "test",
		SorBusinessObjectId:        "test",
		BaseledgerTransactionId:    uuid.NewV4(),
		ReceiverOrgId:              uuid.NewV4(),
		TrustmeshId:                uuid.NewV4(),
	}

	webhook := &types.SorWebhook{
		Body:        "{ 'isApproved': '{{approved}}', 'feedbackMessage': '{{message}}' }",
		BodyParams:  "",
		WebhookType: types.UpdateObject,
	}

	want := "{ 'isApproved': 'APPROVED', 'feedbackMessage': 'this is a feedback message' }"

	result := buildWebhookRequestBody(
		webhook.Body,
		webhook.BodyParams,
		webhook.WebhookType,
		trustmeshEntry,
		"",
		"APPROVED",
		"this is a feedback message",
		"irelevant",
	)

	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}
