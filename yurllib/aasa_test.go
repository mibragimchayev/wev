package yurllib

import (
	"reflect"
	"testing"
)

// TestGetDomain is used for
func TestGetDomain(t *testing.T) {

	testCases := []struct {
		input           string
		expectedOutput  string
		expectedMessage []string
	}{
		{
			"www.chayev.com",
			"www.chayev.com",
			[]string{"WARNING: The URL must use HTTPS, trying HTTPS instead. \n\n"},
		},
		{
			"https://www.chayev.com",
			"www.chayev.com",
			nil,
		},
	}

	for i, tc := range testCases {

		actual, message := getDomain(tc.input)

		if !reflect.DeepEqual(actual, tc.expectedOutput) || !reflect.DeepEqual(message, tc.expectedMessage) {
			t.Errorf("GetDomain test[%d]: expected %v, actual %v", i, tc.expectedOutput, actual)
			t.Errorf("GetDomain test[%d]: expected %v, actual %v", i, tc.expectedMessage, message)
		}
	}
}

// TestEvaluateAASA is used for
func TestEvaluateAASA(t *testing.T) {

}

// TestVerifyJSONformat is used for
func TestVerifyJSONformat(t *testing.T) {

}

// TestVerifyBundleIdentifierIsPresent is used for
func TestVerifyBundleIdentifierIsPresent(t *testing.T) {

}

// TestCheckDomain is used for
func TestCheckDomain(t *testing.T) {
	// Uses HTTP Get - need to figure out how to test properly
	// calls loadAASAContents() which calls makeRequest()
}

// TestLoadAASAContents is used for
func TestLoadAASAContents(t *testing.T) {
	// Uses HTTP Get - need to figure out how to test properly
	// calls makeRequest()
}

// TestMakeRequest is used for
func TestMakeRequest(t *testing.T) {
	// Uses HTTP Get - need to figure out how to test properly
}
