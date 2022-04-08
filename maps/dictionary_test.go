package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	word       = "test"
	definition = "this is original definition"
	updatedDef = "this is updated definition"
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
		err := dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
		assert.Nil(t, err, "Error is not nil, want: %#v, got: %#v", nil, err)
	})
	t.Run("add existing word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add(word, definition)
		assert.Nil(t, err, "Expected first error to be nil but got: %#v", err)
		err = dict.Add(word, updatedDef)
		assertDefinition(t, dict, word, definition)
		assert.EqualError(t, err, ErrWordExists.Error(), "errors does not match")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_no_existing_key", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update(word, updatedDef)
		assert.EqualError(t, err, ErrWordDoesNotExist.Error(), "Erros not equal, got: %#v, want: %#v", err, ErrDidNotFindKey)
		assertDefinition(t, dict, word, "")
	})
	t.Run("update_existing_key", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add(word, definition)
		err := dict.Update(word, updatedDef)
		assert.Nil(t, err, "Expected nil error, got: %#v", err)
		assertDefinition(t, dict, word, updatedDef)
	})
}

func TestDelete(t *testing.T) {
	dict := Dictionary{word: definition}
	dict.Delete(word)
	value, err := dict.Find(word)
	assert.Empty(t, value, "Expected empty value, got: %s", value)
	assert.EqualError(t, err, ErrDidNotFindKey.Error(), "Erros not equal, got: %#v, want: %#v", err, ErrDidNotFindKey)
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	got, _ := dict.Find(word)
	assert.Equal(t, definition, got, "Not matching, want: %s, got: %s", definition, got)
}
