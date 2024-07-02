package empty

import (
	"encoding/json"
)

// List is a type that allows empty slices to be rendered as an empty array.
// This changes the default behaviour from `null` to `[]` when the list is empty.
// It does not change anything when unmarshalling.
type List[T any] []T

// MarshalJSON is replacing the default marshal of an empty list.
func (e List[T]) MarshalJSON() ([]byte, error) {
	if len(e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal([]T(e))
}
