package mapconsolevars

import (
	"os"
	"strings"
)

var mapArgs map[string]string
var arrayArgs []string
var initializated bool = false

// The character separator between key and value
// default is =
// key=value
var separator string = "="

func initialize() {
	mapArgs = make(map[string]string)
	arrayArgs = os.Args
	for _, arg := range arrayArgs {
		splitArg := strings.Split(arg, separator)
		if len(splitArg) == 2 {
			mapArgs[splitArg[0]] = splitArg[1]
		}
	}
	initializated = true
}

// Get the value of a key
//
// Example: go run *.go key1=value1
//
// 	value1, found := mapconsolevars.GetArgsValue("key1")
//	if found {
//		fmt.Println("The value of key1 is", value1)
//	}
func GetValue(key string) (interface{}, bool) {
	if !initializated {
		initialize()
	}
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
	initialize()
}
