package mapconsolevars

import (
	"os"
	"testing"
)

func onInit() {
	os.Args = append(os.Args, "key1=value1")
}
func TestInexistetKey(t *testing.T) {
	onInit()
	_, found := GetValue("ThisKeyShouldNotExixts")
	if found {
		t.Errorf("Found a inexistent key")
	}
}

func TestExistentKey(t *testing.T) {
	onInit()
	value1, found := GetValue("key1")
	if !found {
		t.Errorf("Not found key 'key1', but its present at os.Args")
	}
	if value1 != "value1" {
		t.Errorf("Found key 'key1', but value is diferent")
	}
}
