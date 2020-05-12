//
// Copyright (C) Philip Schlump, 2013-2016.
// Version: 1.0.3
// Tested on Mon Jun 20 18:01:48 MDT 2016
//
package filelib

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
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

// GetFilenames gets a list of file names and directorys.
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

type ApplyFilenamesFunc func(ty string, fn string, fstat os.FileInfo)

func ApplyFilenames(dir string, fx ApplyFilenamesFunc) (filenames, dirs []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, nil
	}
	for _, fstat := range files {
		if !strings.HasPrefix(string(fstat.Name()), ".") {
			if fstat.IsDir() {
				dirs = append(dirs, fstat.Name())
				fx("dir", fstat.Name(), fstat)
			} else {
				filenames = append(filenames, fstat.Name())
				fx("file", fstat.Name(), fstat)
			}
		}
	}
	return
}

func CleanupOldFiles(dir string, dt time.Duration) {

	ApplyFilenames(dir, func(ty, fn string, fstat os.FileInfo) {
		if ty == "file" {
			modifiedtime := fstat.ModTime()
			duration := time.Since(modifiedtime)
			fmt.Printf("Fn: %s Seconds %v compare to: %v\n", fn, duration, dt)
			// if duration.Seconds() > dt.Seconds() {
			if duration > dt {
				fmt.Printf("	Do Cleanup %s becomes ,%s\n", fn, fn)
				// os.Rename(dir+"/"+fn, dir+"/"+","+fn)
				os.Remove(dir + "/" + fn)
			}
		}
	})

	// nw := time.Now()
	// modifiedtime := file.ModTime()
	// duration := time.Since(then)
	// fmt.Println(duration.Hours())
}

// ExistsIsDir returns true if the directory exists.
func ExistsIsDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if fi.IsDir() {
		return true
	}
	return false
}

// InArray returns true if `lookFor` is found in `inArr`.
func InArray(lookFor string, inArr []string) bool {
	for _, v := range inArr {
		if lookFor == v {
			return true
		}
	}
	return false
}

// InArrayN returns the position if `lookFor` is found in `inArr`, else -1.
func InArrayN(lookFor string, inArr []string) int {
	for i, v := range inArr {
		if lookFor == v {
			return i
		}
	}
	return -1
}

func FilterArray(re string, inArr []string) (outArr []string) {
	var validID = regexp.MustCompile(re)

	outArr = make([]string, 0, len(inArr))
	for k := range inArr {
		if validID.MatchString(inArr[k]) {
			outArr = append(outArr, inArr[k])
		}
	}
	// fmt.Printf ( "output = %v\n", outArr )
	return
}

func AllFiles(path, match string) (fns, dirs []string) {
	fns, dirs = GetFilenames(path)
	fns = FilterArray(match, fns)
	// xyzzy400 - redursive
	for _, aDir := range dirs {
		tFn, tDir := AllFiles(path+"/"+aDir, match)
		for _, x := range tFn {
			fns = append(fns, aDir+"/"+x)
		}
		_ = tDir
	}
	return
}

/* vim: set noai ts=4 sw=4: */
