package mydict

import "errors"

// type을 통해 alias
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// map의 key에 해당하는 값과 존재하는지에 관한 여부(bool)을 가져올 수 있다.
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }

}

// Update  a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = definition
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {

	delete(d, word)
}
