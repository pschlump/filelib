# filelib

 [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/pschlump/Go-FTL/master/LICENSE)

Go (golang) miscellaneous file and templating routines

A library of convenience functions for dealing with files.

```
func Exists(name string) bool {
```

Return true if file exists.


```
func GetFilenames(dir string) (filenames, dirs []string) {
```

GetFilenames gets a list of file names and directorys.

```
func ExistsIsDir(name string) bool {
```

Returns true if name exists and is a directory.


```
func InArray(lookFor string, inArr []string) bool {
```

Returns true if `lookFor` is in the array `inArr`.  A more general form of this is in the `pluto` library that uses generics.

```
func InArrayN(lookFor string, inArr []string) int {
```

Searches the array `inArr` for `lookFor` returning the position if found, or -1 if not found.

```
func InArrayInt(lookFor int, inArr []int) bool {
```

Searches the array  of int `inArr`  for the value `lookFor` returning true if found.


```
func AllFiles(path, match string) (fns, dirs []string) {
```

Recursively searches for all the files in a path.

