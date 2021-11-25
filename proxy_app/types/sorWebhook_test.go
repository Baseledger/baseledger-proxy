package types

import (
	"strings"
	"testing"
)

func TestGivenRequestParamParseRequestParamsIntoStringProducesCorrectFormat(t *testing.T) {
	requestParam := &RequestParam{
		ParamName:       "testName",
		ParamValueField: "testValueField",
	}

	want := "testName:testValueField;"

	result := ParseRequestParamsIntoString([]RequestParam{*requestParam})

	if strings.Compare(want, result) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}

func TestGivenRequestParamStringParseStringIntoRequestParamsProducesCorrectFormat(t *testing.T) {
	requestParamString := "testName:testValueField;"

	want := &RequestParam{
		ParamName:       "testName",
		ParamValueField: "testValueField",
	}

	result := ParseStringIntoRequestParams(requestParamString)

	if strings.Compare(want.ParamName, result[0].ParamName) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}

	if strings.Compare(want.ParamValueField, result[0].ParamValueField) != 0 {
		t.Fatalf(`TestHook = %q,  want match for %#q, nil`, result, want)
	}
}
