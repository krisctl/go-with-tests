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
		{name: "key present", key: "k2", value: "v2", err: nil},
		{name: "key not present", key: "nk3", value: "", err: errDidNotFindKey},
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
