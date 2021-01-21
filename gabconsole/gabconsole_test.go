package gabconsole

import "testing"

func TestGabconsole(t *testing.T) {
	want := "Hello, console"
	if got := GabConsole(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
