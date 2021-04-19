package mapconsolevars

import (
	"os"
	"strings"
)

var (
	// The map with the args
	// Example: Example: go run *.go key1=value1 key2=value2
	// the mapArgs hold the key/value pairs:
	// mapArgs["key1"] = value1
	// mapArgs["key2"] = value2
	mapArgs map[string]string

	// The character separator between key and value
	// default is =
	separator string = "="
)

func init() {
	scan()
}

func scan() {
	mapArgs = make(map[string]string)
	for _, arg := range os.Args {
		splitArg := strings.Split(arg, separator)
		if len(splitArg) == 2 {
			mapArgs[splitArg[0]] = splitArg[1]
		}
	}
}

// Get the value of a key
//
// Example: go run *.go key1=value1
//
// 	value1, found := mapconsolevars.GetArgsValue("key1")
//	if found {
//		fmt.Println("The value of key1 is", value1)
//	}
func GetValue(key string) (string, bool) {
	value := mapArgs[key]
	if value == "" {
		return "", false
	}
	return value, true
}

// Change tha character separator. The default is "="
//
// Example: go run *.go key1:value1
//
//  mapconsolevars.ChangeSeparatorTo(":")
// 	value1, found := mapconsolevars.GetArgsValue("key1")
//	if found {
//		fmt.Println("The value of key1 is", value1)
//	}
func ChangeSeparatorTo(newSeparator string) {
	separator = newSeparator
	scan()
}
