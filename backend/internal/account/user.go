package account

// User represents data associated with a single Endpoint user
type User struct {
	ID      ID        `json:"id"`
	APIKeys []*APIKey `json:"apiKeys"`
}

// NewUser returns an initialized user with the given ID.  The ID should correspond to the user's ID in the external authentication system.
func NewUser(id ID) *User {
	return &User{
		ID: id,
		APIKeys: []*APIKey{
			NewAPIKey(false),
			NewAPIKey(true),
		},
	}
}

// ID is a unique indentifier for a User
type ID string

// String returns the string value of a given ID
func (i ID) String() string {
	return string(i)
}

// Datastore is a persistence layer for Users
type Datastore interface {
	SaveUser(*User) error
	GetUser(ID) (*User, error)
}
