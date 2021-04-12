package mapconsolevars

import (
	"os"
	"testing"
)

func TestInexistetKey(t *testing.T) {
	_, found := GetArgsValue("key1")
	if found {
		t.Errorf("Found a inexistent key")
	}
}

func TestExistentKey(t *testing.T) {
	os.Args = append(os.Args, "key1=value1")
	value1, found := GetArgsValue("key1")
	if !found {
		t.Errorf("Not found key 'key1', but its present at os.Args")
	}
	if value1 != "value1" {
		t.Errorf("Found key 'key1', but value is diferent")
	}
}
