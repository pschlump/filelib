package filelib

func ReplaceExt(fn string, newExt string) (out string) {
	t := RmExt(fn)
	out = t + newExt
	return
}

/* vim: set noai ts=4 sw=4: */
