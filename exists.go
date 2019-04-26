//
// Copyright (C) Philip Schlump, 2013-2016.
// Version: 1.0.3
// Tested on Mon Jun 20 18:01:48 MDT 2016
//
package filelib

import (
	"io/ioutil"
	"os"
	"strings"
)

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// -------------------------------------------------------------------------------------------------
// Get a list of file names and directory.
// -------------------------------------------------------------------------------------------------
func GetFilenames(dir string) (filenames, dirs []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, nil
	}
	for _, fstat := range files {
		if !strings.HasPrefix(string(fstat.Name()), ".") {
			if fstat.IsDir() {
				dirs = append(dirs, fstat.Name())
			} else {
				filenames = append(filenames, fstat.Name())
			}
		}
	}
	return
}
