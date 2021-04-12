package mapconsolevars

import (
	"os"
	"testing"
)

func setArgsWithEqualCharacter() {
	os.Args = append(os.Args, "key1=value1")
}
func setArgsWithDotCharacter() {
	os.Args = append(os.Args, "key1:value1")
}

func TestInexistetKey(t *testing.T) {
	setArgsWithEqualCharacter()
	_, found := GetValue("ThisKeyShouldNotExixts")
	if found {
		t.Errorf("Found a inexistent key")
	}
}

func TestExistentKey(t *testing.T) {
	setArgsWithEqualCharacter()
	value1, found := GetValue("key1")
	if !found {
		t.Errorf("Not found key 'key1', but its present at os.Args")
	}
	if value1 != "value1" {
		t.Errorf("Found key 'key1', but value is diferent")
	}
}

func TestExistentKeyWithDotSeparator(t *testing.T) {
	setArgsWithDotCharacter()
	ChangeSeparatorTo(":")
	value1, found := GetValue("key1")
	if !found {
		t.Errorf("Not found key 'key1', but its present at os.Args")
	}
	if value1 != "value1" {
		t.Errorf("Found key 'key1', but value is diferent")
	}
}
