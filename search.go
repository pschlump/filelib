//
// Copyright (C) Philip Schlump, 2013-2015.
// Version: 1.0.3
// Tested on Mon Sep  7 21:05:15 MDT 2015
//
package filelib

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

func SearchPathApp(rawFileName string, appName string, searchPath string) (fullFileName string, ok bool) {

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error(10020): Unable to get the hostname (%v)\n", err)
		os.Exit(1)
	}

	mdata := make(map[string]string, 30)
	mdata["HostName"] = hostname
	mdata["AppName"] = appName
	mdata["IS_WINDOWS"] = ""
	ps := string(os.PathSeparator)
	if ps != "/" {
		mdata["IS_WINDOWS"] = ""
	} else {
		mdata["IS_WINDOWS"] = "ms"
	}
	mdata["HOME"] = os.Getenv("HOME")
	mdata["FILENAMERAW"] = rawFileName
	mdata["FILENAME"] = RmExt(rawFileName)
	mdata["FILEEXT"] = filepath.Ext(rawFileName)
	if ps != "/" {
		ps = ps + ps
	}
	mdata["OS_SEP"] = ps

	sp := strings.Split(searchPath, string(os.PathListSeparator))
	ok = false
	tmpl_arr := []string{
		"%{CUR_PATH%}%{OS_SEP%}%{FILENAME%}-%{AppName%}-%{HostName%}%{FILEEXT%}",
		"%{CUR_PATH%}%{OS_SEP%}%{FILENAME%}-%{AppName%}%{FILEEXT%}",
		"%{CUR_PATH%}%{OS_SEP%}%{FILENAME%}-%{HostName%}%{FILEEXT%}",
		"%{CUR_PATH%}%{OS_SEP%}%{FILENAME%}%{FILEEXT%}",
		"%{CUR_PATH%}%{OS_SEP%}%{FILENAMERAW%}",
	}
	for _, p := range sp {
		mdata["CUR_PATH"] = p

		for _, tmpl := range tmpl_arr {
			fullFileName = Qt(tmpl, mdata)
			fullFileName, _ = SubstitueUserInFilePath(fullFileName, mdata)
			fullFileName = Qt(fullFileName, mdata)
			// fmt.Printf("1: %s\n", fullFileName)
			if Exists(fullFileName) {
				ok = true
				return
			}
		}

	}
	fullFileName = rawFileName
	fullFileName, _ = SubstitueUserInFilePath(fullFileName, mdata)
	fullFileName = Qt(fullFileName, mdata)
	ok = Exists(fullFileName)
	return
}

func RmExt(filename string) string {
	var extension = filepath.Ext(filename)
	var name = filename[0 : len(filename)-len(extension)]
	return name
}

var hasUserPat *regexp.Regexp
var replUserPat *regexp.Regexp
var homeDir string

func init() {
	ps := string(os.PathSeparator)
	if ps != "/" {
		ps = ps + ps
	}

	hasUserPat = regexp.MustCompile("~([a-zA-Z][^" + ps + "]*)" + ps)
	replUserPat = regexp.MustCompile("(~[a-zA-Z][^" + ps + "]*)")

	homeDir = os.Getenv("HOME")
}

func SubstitueUserInFilePath(s string, mdata map[string]string) (rs string, has bool) {
	has = false
	x := hasUserPat.FindStringSubmatch(s)
	// fmt.Printf("x=%s\n", SVar(x))
	rs = s
	if len(x) > 1 {
		has = true
		p := x[1]
		ud, err := user.Lookup(p)
		if err != nil {
			fmt.Printf("Error (13922): unable to lookup %s as a username, error=%s\n", p, err)
		} else {
			mdata["USER_"+ud.Username] = ud.HomeDir
			rs = replUserPat.ReplaceAllLiteralString(rs, "%{USER_"+ud.Username+"%}")
		}
	} else if strings.HasPrefix(rs, "~") {
		// fmt.Printf("Before last substitue rs [%s]\n", rs)
		rs = strings.Replace(rs, "~", "%{HOME%}", 1)
		// fmt.Printf("At bottom rs [%s]\n", rs)
	}
	return
}
