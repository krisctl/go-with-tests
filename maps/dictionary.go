package maps

import "errors"

type Dictionary map[string]string

var errDidNotFindKey = errors.New("could not find the word you were looking for")

func (d Dictionary) Find(k string) (string, error) {
	v, ok := d[k]
	if !ok {
		return "", errDidNotFindKey
	}
	return v, nil
}

func (d Dictionary) Add(k, v string) {
	d[k] = v
}
