package user

// ID is contains the unique identifier for a user.
type ID struct {
	id string
}

// NewID returns a new ID corresponding to the given ID string.
func NewID(uid string) *ID {
	return &ID{id: uid}
}

// Value returns the internal user ID's string representation
func (i *ID) Value() string {
	return i.id
}
