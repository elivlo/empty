## empty [![GoDoc](https://godoc.org/github.com/elivlo/empty?status.svg)](https://godoc.org/github.com/elivlo/empty)
`import "github.com/elivlo/empty"`

empty is a library to only supports marshalling empty lists with square brackets instead of null.
This type won't be needed for communication between Go services but helps a lot with Pydantic in Python.
There is no need for declaring a type in python as `Optional` with this package.

#### Interfaces

- All types also implement `json.Marshaler` and `json.Unmarshaler`, so you can marshal them to their native JSON representation.

#### empty.List[`T`]
Represents a slice of type T.

It marshals an empty list to brackets `[]` instead of `null` but unmarshalls like a normal slice.

### License
MIT
