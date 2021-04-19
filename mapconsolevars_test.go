package mapconsolevars

import (
	"testing"
)

//
//  need to run:  go test -args key1:value1 keyDots1=valueDots1
//

func TestInexistetKey(t *testing.T) {
	_, found := GetValue("ThisKeyShouldNotExixts")
	if found {
		t.Errorf("Found a inexistent key")
	}
}

func TestExistentKey(t *testing.T) {
	value1, found := GetValue("key1")
	if !found {
		t.Errorf("Not found key 'key1'. Did u run go test -args key1=value1 ?")
	}
	if value1 != "value1" {
		t.Errorf("Found key 'key1', but value is diferent. Did u run go test -args key1=value1 ?")
	}
}

func TestExistentKeyWithDotSeparator(t *testing.T) {
	ChangeSeparatorTo(":")
	value1, found := GetValue("keyDots1")
	if !found {
		t.Errorf("Not found key 'key1'. Did u run go test -args keyDots1:valueDots1 ?")
		return
	}
	if value1 != "valueDots1" {
		t.Errorf("Found key 'keyDots1', but value is diferent. Did u run go test -args keyDots1:valueDots1 ?")
	}
}
