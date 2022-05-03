package bot

import (
	"testing"
)

func TestLogin(t *testing.T) {
	testCases := []struct {
		sessionid1 string
		sessionid2 string
	}{
		{"12345678912%1234567812%90876543212", "12345678912%1234567812%908765432123"},
	}
	got := Login("username", "password")
	for _, tc := range testCases {
		if len(got) != len(tc.sessionid1) && len(got) != len(tc.sessionid2) {
			t.Errorf("TestLogin() returned %v; wanted %v or %v", len(got), len(tc.sessionid1), len(tc.sessionid2))
		}
	}
}
