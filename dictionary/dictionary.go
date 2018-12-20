//Package dictionary is for exploring maps
package dictionary

//Dictionary custom map type
type Dictionary map[string]string

const (
	//ErrNotFound is the not found error
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	//ErrWordExists returned when the value already exists for an add
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	//ErrWordDoesNotExist returned when trying to update a non-existent word
	ErrWordDoesNotExist = DictionaryErr("cannot find the word you want to update")
)

//DictionaryErr is the custom type errors for Dictionary Type
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search looks for the requested entry in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add puts a new definition in the Dictionary for the provided word
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

//Update changes a definition to a new value
func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return nil
	}

	return nil
}

//Delete removes the desired word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
