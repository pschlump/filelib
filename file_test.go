package filelib

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func Test_Exists(t *testing.T) {
	os.Mkdir("./test_data", 0700)
	err := ioutil.WriteFile("./test_data/file-exists", []byte("bob\n"), 0600)
	if err != nil {
		t.Errorf("Unable to setup test file, ./test_data/file-exists")
	}

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

// func CleanupOldFiles(dir string, dt time.Duration) {
func Test_CleanupOldFiles(t *testing.T) {
	CleanupOldFiles("./test-cleanup", time.Duration(30)*time.Second)
}

/* vim: set noai ts=4 sw=4: */
