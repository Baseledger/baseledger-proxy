package systemofrecord

import (
	"strings"
	"testing"

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/types"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHook(t *testing.T) {
	trustmeshEntry := &types.TrustmeshEntry{
		BaseledgerBusinessObjectId: "test",
		BusinessObjectType:         "test",
		SorBusinessObjectId:        "test",
		BaseledgerTransactionId:    uuid.UUID{},
		ReceiverOrgId:              uuid.UUID{},
		TrustmeshId:                uuid.UUID{},
	}

	want := "https://test.url/test"

	result := buildWebhookRequestUrl("https://test.url/{{bo_id}}", "bo_id:BaseledgerBusinessObjectId", trustmeshEntry)
	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}
