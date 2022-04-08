package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary(t *testing.T) {
	d := Dictionary{"k1": "v1"}
	want := "v1"
	got, _ := d.Find("k1")
	assert.Equal(t, want, got, "did not get expected result, got: %s, want: %s", got, want)
}

func TestDictionaryTable(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value string
		err   error
	}{
		{name: "find, key present", key: "k2", value: "v2", err: nil},
		{name: "find, key not present", key: "nk3", value: "", err: ErrDidNotFindKey},
	}
	d := Dictionary{"k2": "v2"}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := d.Find(tC.key)
			assert.Equal(t, tC.value, got, "did not get expected result, got: %s, want: %s", got, tC.value)
			assert.Equal(t, err, tC.err, "did not get expected error, got: %s, want: %s", err, tC.err)
		})
	}
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
		assert.Nil(t, err, "Error is not nil, want: %#v, got: %#v", nil, err)
	})
	t.Run("add existing word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def1 := "original"
		def2 := "updated"
		err := dict.Add(word, def1)
		assert.Nil(t, err, "Expected first error to be nil but got: %#v", err)
		err = dict.Add(word, def2)
		assertDefinition(t, dict, word, def1)
		assert.EqualError(t, err, ErrWordExists.Error(), "errors does not match")
	})
}

func TestUpdate(t *testing.T) {
	dict := Dictionary{}
	word := "test"
	definition := "this is original test"
	updatedDef := "this is updated test"
	dict.Add(word, definition)
	dict.Update(word, updatedDef)
	assertDefinition(t, dict, word, updatedDef)
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	got, _ := dict.Find(word)
	assert.Equal(t, definition, got, "Not matching, want: %s, got: %s", definition, got)
}
