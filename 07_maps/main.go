package main

func main() {
	// entry point
}

type Dictionary map[string]string

func (dictionary Dictionary) Delete(word string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		 delete(dictionary, word)
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(word string, newDefinition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Add(word string, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string { return string(e) }

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
