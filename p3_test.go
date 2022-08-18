package filelib

import "testing"

func Test_ReplaceExt(t *testing.T) {
	// func ReplaceExt(fn string, newExt string) (out string) {
	a := ReplaceExt("bob.jpg", ".png")
	if a != "bob.png" {
		t.Errorf("expected bob.png got ->%s<-", a)
	}
	a = ReplaceExt("./abc/bob.jpg", ".png")
	if a != "./abc/bob.png" {
		t.Errorf("expected ./abc/bob.png got ->%s<-", a)
	}
	a = ReplaceExt("/abc/bob.jpg", ".png")
	if a != "/abc/bob.png" {
		t.Errorf("expected /abc/bob.png got ->%s<-", a)
	}
}

/* vim: set noai ts=4 sw=4: */
