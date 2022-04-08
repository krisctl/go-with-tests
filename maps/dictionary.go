package maps

type Dictionary map[string]string

type DictionaryError string

func (de DictionaryError) Error() string {
	return string(de)
}

const (
	ErrDidNotFindKey    = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("word already exists in the map")
	ErrWordDoesNotExist = DictionaryError("word does not exist")
)

func (d Dictionary) Find(k string) (string, error) {
	v, ok := d[k]
	if !ok {
		return "", ErrDidNotFindKey
	}
	return v, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Find(k)

	switch err {
	case ErrDidNotFindKey:
		d[k] = v
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Find(k)
	switch err {
	case ErrDidNotFindKey:
		return ErrWordDoesNotExist
	case nil:
		d[k] = v
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
