package filelib

import "testing"

// InPatternArray seqrches a set of regular expression patterns in 'pat' and sees if any match.
// -1 is returned if no match.
// func InPatternArray(s string, pat []string) (rv int, err error) {
func Test_InPatternArray(t *testing.T) {
	pat := []string{"abc*d", "aabbcc[0-9]*dd"}

	rv, err := InPatternArray("xxx", pat)
	if err != nil {
		t.Errorf("Unexpeced error %s", err)
	} else if rv != -1 {
		t.Errorf("Expected -1 (not found) got %d\n", rv)
	}

	rv, err = InPatternArray("cc3dd", pat)
	if err != nil {
		t.Errorf("Unexpeced error %s", err)
	} else if rv != -1 {
		t.Errorf("Expected 1 got %d\n", rv)
	}

}
