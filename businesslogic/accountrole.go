package businesslogic

import (
	"errors"
	"time"
)

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

func NewAccountRole(user Account, accountType int) AccountRole {
	return AccountRole{
		AccountID:       user.ID,
		AccountTypeID:   accountType,
		CreateUserID:    user.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    user.ID,
		DateTimeUpdated: time.Now(),
	}
}

// SearchAccountRoleCriteria specifies the parameters that can be used to search Account Roles in a repository
type SearchAccountRoleCriteria struct {
	ID            int
	AccountID     int
	AccountTypeID int
}

// IAccountRoleRepository defines the functions that an account role repository should implement
type IAccountRoleRepository interface {
	CreateAccountRole(role *AccountRole) error
	SearchAccountRole(criteria SearchAccountRoleCriteria) ([]AccountRole, error)
}

type AccountRoleProvisionService struct {
	accountRepo               IAccountRepository
	accountRoleRepo           IAccountRoleRepository
	organizerProvisionService OrganizerProvisionService
}

func NewAccountRoleProvisionService(
	accountRepo IAccountRepository,
	accountRoleRepo IAccountRoleRepository) *AccountRoleProvisionService {
	service := AccountRoleProvisionService{
		accountRepo:     accountRepo,
		accountRoleRepo: accountRoleRepo,
	}
	return &service
}

func (service AccountRoleProvisionService) GrantRole(currentUser Account, account Account, usertype int) error {
	if account.HasRole(usertype) {
		return errors.New("this account already has this role")
	}

	newRole := AccountRole{
		AccountID:       account.ID,
		AccountTypeID:   usertype,
		CreateUserID:    currentUser.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUser.ID,
		DateTimeUpdated: time.Now(),
	}
	return service.accountRoleRepo.CreateAccountRole(&newRole)
}
