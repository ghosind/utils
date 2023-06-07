package utils

import "testing"

func TestCloneMap(t *testing.T) {
	languages := map[string]string{
		"zh": "你好",
		"en": "Hello",
		"fr": "Bonjour",
	}
	clone := CloneMap(languages)
	if len(languages) != len(clone) {
		t.Errorf("new map's length expect %d, actually %d", len(languages), len(clone))
	}
	for k, v := range languages {
		cv, ok := languages[k]
		if !ok {
			t.Errorf("expect key %s in the new map", k)
		} else if v != cv {
			t.Errorf("the expecting value of key %s in the new map is %s, actually %s", k, v, cv)
		}
	}
}

func TestCloneMapNil(t *testing.T) {
	var oldMap map[string]string
	newMap := CloneMap(oldMap)
	if newMap != nil {
		t.Errorf("The new map expect nil")
	}
}

func TestCopyMap(t *testing.T) {
	src := map[string]string{
		"zh": "你好",
		"en": "Hello",
		"fr": "Bonjour",
	}
	dest := map[string]string{
		"es": "Hola",
		"fr": "Salut",
	}
	CopyMap(dest, src)
	for k, v := range src {
		cv, ok := dest[k]
		if !ok {
			t.Errorf("expect key %s in the destination map", k)
		} else if v != cv {
			t.Errorf("the expecting value of key %s in the destination map is %s, actually %s", k, v, cv)
		}
	}
}

func TestCopyMapNil(t *testing.T) {
	var src map[string]string
	var dest map[string]string

	if err := CopyMap(dest, src); err != nil {
		t.Errorf("Expect no error, actually %v", err)
	}

	src = map[string]string{
		"zh": "你好",
		"en": "Hello",
		"fr": "Bonjour",
	}
	if err := CopyMap(dest, src); err == nil {
		t.Error("Expect nil destination error, actually no error")
	}
}
