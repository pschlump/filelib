//
// Copyright (C) Philip Schlump, 2013-2016.
// Version: 1.0.3
// Tested on Mon Jun 20 18:01:48 MDT 2016
//
package filelib

import "os"

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
