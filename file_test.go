package filelib

import "testing"

func Test_Exists(t *testing.T) {
	if Exists("./no-such-file.do-not-create") {
		t.Errorf("Expected false, file should not exist\n")
	}
	if !Exists("./test_data") {
		t.Errorf("Expected true, check of directory\n")
	}
	if !Exists("./test_data/file-exists") {
		t.Errorf("Expected true, check of directory\n")
	}
}
