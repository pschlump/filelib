package filelib

import (
	"fmt"
	"testing"
)

func Test_AllFiles(t *testing.T) {
	tests := []struct {
		pth           string   //
		expectedFiles []string //
	}{
		{
			pth:           "tmpl1",
			expectedFiles: []string{""},
		},
	}
	for ii, test := range tests {
		fn, dr := AllFiles(test.pth, "((.html)|(.tmpl))$") // func AllFiles(path, match string) (fns, dirs []string) {
		if false {
			fmt.Printf("fn ->%s<- dr ->%s<-\n", fn, dr)
		}
		if !InArray("a/b/ab.html", fn) {
			t.Errorf("Error %2d, Invalid Failed to find a/b/ab.html\n", ii)
		}
	}
}
