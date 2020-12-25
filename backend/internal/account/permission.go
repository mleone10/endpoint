package account

// Permission contains information regarding allowed operations for a given account ID.
type Permission struct {
	ID       ID
	ReadOnly bool
}
