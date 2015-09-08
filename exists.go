//
// Copyright (C) Philip Schlump, 2013-2015.
// Version: 1.0.2
// Tested on Wed Sep  2 21:28:25 MDT 2015
//
package filelib

import "os"

// -------------------------------------------------------------------------------------------------
// Tested
// Exists reports whether the named file or directory exists.
// -------------------------------------------------------------------------------------------------
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
