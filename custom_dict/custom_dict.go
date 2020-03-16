package custom_dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errExists = errors.New("word exists")

// Search
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add
func (d Dictionary) Add(word string, value string) error {
	_, err := d.Search(word)
	if err != nil {
		d[word] = value
		return nil
	}
	return errExists
}

// Update
func (d Dictionary) Update(word string, value string) error {
	_, err := d.Search(word)
	if err == nil {
		return errNotFound
	}
	d[word] = value
	return nil
}

// Delete
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
		return nil
	}
	return errNotFound
}
