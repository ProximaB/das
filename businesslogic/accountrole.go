package businesslogic

import "time"

// AccountRole defines the role that an account can be associated with
type AccountRole struct {
	ID              int
	AccountID       int
	AccountTypeID   int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// SearchAccountRoleCriteria specifies the parameters that can be used to search Account Roles in a repository
type SearchAccountRoleCriteria struct {
	AccountID     int
	AccountTypeID int
}

// IAccountRoleRepository defines the functions that an account role repository should implement
type IAccountRoleRepository interface {
	CreateAccountRole(role *AccountRole) error
	DeleteAccountRole(role AccountRole) error
	SearchAccountRole(criteria SearchAccountRoleCriteria) ([]AccountRole, error)
	UpdateAccountRole(role AccountRole) error
}
