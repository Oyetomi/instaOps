package bot

import "testing"

const (
	pathToSettings = "settings"
)

func TestLogin(t *testing.T) {
	sessionid := Login(pathToSettings)
	want := len(sessionid)
	got := len(sessionid)
	if want != got {
		t.Fatalf("wanted %d got %d", want, got)
	}

}
