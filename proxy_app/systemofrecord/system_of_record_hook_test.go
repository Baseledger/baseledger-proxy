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
