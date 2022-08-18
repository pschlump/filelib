package filelib

import (
	"fmt"
	"regexp"
)

var qtRegEx *regexp.Regexp

func init() {
	qtRegEx = regexp.MustCompile("%{([A-Za-z0-9_]*)%}")
}

// Qt is a string quick template.
// %{name%} gets replace with substitution from map if it is in map, else ""
func Qt(format string, data map[string]string) string {
	// re := regexp.MustCompile("%{([A-Za-z0-9_]*)%}")
	return string(qtRegEx.ReplaceAllFunc([]byte(format), func(s []byte) []byte {
		return []byte(data[string(s[2:len(s)-2])])
	}))
}

// QtR is a string quick template.  It uses the same format as Qt but you get to pass a map[string]interface{}
func QtR(format string, data map[string]interface{}) string {
	return string(qtRegEx.ReplaceAllFunc([]byte(format), func(s []byte) []byte {
		t := string(s[2 : len(s)-2])
		u, ok := data[t]
		if !ok {
			return []byte("")
		}
		switch u.(type) {
		case string:
			return []byte(u.(string))
		default:
			sb := fmt.Sprintf("%v", u)
			return []byte(sb)
		}
	}))
}

/* vim: set noai ts=4 sw=4: */
