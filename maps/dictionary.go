package maps

import "errors"

type Dictionary map[string]string

var (
	ErrDidNotFindKey = errors.New("could not find the word you were looking for")
	ErrWordExists    = errors.New("Word already exists in the map")
)

func (d Dictionary) Find(k string) (string, error) {
	v, ok := d[k]
	if !ok {
		return "", ErrDidNotFindKey
	}
	return v, nil
}

func (d Dictionary) Add(k, v string) error {
	_, ok := d[k]
	if ok {
		return ErrWordExists
	}
	d[k] = v
	return nil
}
