package dictionary

type Dictionary map[string]string

const (
	ErrNotFound   = errDictionary("could not find the word you were looking for")
	ErrWordExists = errDictionary("cannot add word because it already exists")
)

type errDictionary string

func (e errDictionary) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, defination string) {
	d[word] = defination
}
