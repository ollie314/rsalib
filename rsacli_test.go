package rsalib

import "testing"

func testRsaKey(t *testing.T) {
	want := "test"
	if got := RsaKey(); got != want {
		t.Errorf("%q obtained, expect %q", got, want)
	}
}
