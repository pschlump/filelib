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

func Test_ListFilesAndRemoveItem(t *testing.T) {
	fns, dirs := GetFilenames("./test_data_2")
	if len(fns) != 3 {
		t.Errorf("Error, Invalid Length expected 3 got %d\n", len(fns))
	}
	if len(dirs) != 1 {
		t.Errorf("Error, Invalid Length expected 1 got %d\n", len(dirs))
	}
	fns = RemoveMatch("^bbb.bbb$", fns)
	if len(fns) != 2 {
		t.Errorf("Error, Invalid Length expected 2 got %d\n", len(fns))
	}
}
