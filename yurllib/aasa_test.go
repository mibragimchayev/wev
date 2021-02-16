package yurllib

import (
	"reflect"
	"strings"
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
			t.Errorf("getDomain test[%d]: expected %v, actual %v", i, tc.expectedOutput, actual)
			t.Errorf("getDomain test[%d]: expected %v, actual %v", i, tc.expectedMessage, message)
		}
	}
}

// TestEvaluateAASA is used for
func TestEvaluateAASA(t *testing.T) {

	testCases := []struct {
		input       []byte
		contentType []string
		bundleID    string
		teamID      string
		encrypted   bool
	}{
		{
			[]byte{123, 34, 97, 112, 112, 108, 105, 110, 107, 115, 34, 58, 123, 34, 97, 112, 112, 115, 34, 58, 91, 93, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 91, 123, 34, 97, 112, 112, 73, 68, 34, 58, 34, 81, 52, 51, 75, 83, 56, 88, 77, 50, 65, 46, 99, 111, 109, 46, 107, 111, 104, 108, 115, 46, 109, 99, 111, 109, 109, 101, 114, 99, 101, 46, 105, 112, 104, 111, 110, 101, 34, 44, 34, 112, 97, 116, 104, 115, 34, 58, 91, 34, 47, 102, 74, 89, 114, 47, 42, 34, 93, 125, 44, 123, 34, 97, 112, 112, 73, 68, 34, 58, 34, 50, 90, 81, 76, 72, 65, 81, 69, 57, 90, 46, 99, 111, 109, 46, 107, 111, 104, 108, 115, 46, 101, 110, 116, 101, 114, 112, 114, 105, 115, 101, 46, 105, 112, 104, 111, 110, 101, 34, 44, 34, 112, 97, 116, 104, 115, 34, 58, 91, 34, 47, 102, 74, 89, 114, 47, 42, 34, 93, 125, 44, 123, 34, 97, 112, 112, 73, 68, 34, 58, 34, 81, 52, 51, 75, 83, 56, 88, 77, 50, 65, 46, 99, 111, 109, 46, 107, 111, 104, 108, 115, 46, 109, 99, 111, 109, 109, 101, 114, 99, 101, 46, 105, 112, 97, 100, 34, 44, 34, 112, 97, 116, 104, 115, 34, 58, 91, 34, 47, 107, 55, 70, 50, 47, 42, 34, 93, 125, 44, 123, 34, 97, 112, 112, 73, 68, 34, 58, 34, 50, 90, 81, 76, 72, 65, 81, 69, 57, 90, 46, 99, 111, 109, 46, 107, 111, 104, 108, 115, 46, 101, 110, 116, 101, 114, 112, 114, 105, 115, 101, 46, 116, 99, 111, 109, 119, 114, 97, 112, 112, 101, 114, 34, 44, 34, 112, 97, 116, 104, 115, 34, 58, 91, 34, 47, 107, 55, 70, 50, 47, 42, 34, 93, 125, 93, 125, 125},
			[]string{"application/pkcs7-mime"},
			"Q43KS8XM2A",
			"com.kohls.mcommerce.iphone",
			false,
		},
	}

	for i, tc := range testCases {

		messages, _ := evaluateAASA(tc.input, tc.contentType, tc.bundleID, tc.teamID, tc.encrypted)

		message := strings.Join(messages, " ")

		if !strings.Contains(message, "applinks") {
			t.Errorf("evaluateAASA test[%d]: does not contain 'applinks' key.", i)
		}
	}
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
