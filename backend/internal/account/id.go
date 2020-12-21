package account

// ID is a unique indentifier for a User
type ID string

// String returns the string value of a given ID
func (i ID) String() string {
	return string(i)
}
