package utils

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCloneMap(t *testing.T) {
	a := assert.New(t)

	languages := map[string]string{
		"zh": "你好",
		"en": "Hello",
		"fr": "Bonjour",
	}
	clone := CloneMap(languages)
	a.EqualNow(len(clone), len(languages))

	for k, v := range languages {
		cv, ok := languages[k]
		a.TrueNow(ok)
		a.EqualNow(cv, v)
	}
}

func TestCloneMapNil(t *testing.T) {
	a := assert.New(t)

	var oldMap map[string]string
	newMap := CloneMap(oldMap)
	a.NilNow(newMap)
}

func TestCopyMap(t *testing.T) {
	a := assert.New(t)

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
		a.TrueNow(ok)
		a.EqualNow(cv, v)
	}
}

func TestCopyMapNil(t *testing.T) {
	a := assert.New(t)

	var src map[string]string
	var dest map[string]string

	err := CopyMap(dest, src)
	a.NilNow(err)

	src = map[string]string{
		"zh": "你好",
		"en": "Hello",
		"fr": "Bonjour",
	}
	err = CopyMap(dest, src)
	a.NotNilNow(err)
}
