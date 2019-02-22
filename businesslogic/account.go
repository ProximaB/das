package businesslogic

import (
	"errors"
	"time"
)

// Account is the base account data for all users in DAS. Some fields are required with others are not
type Account struct {
	ID                    int    // userID will be account ID, too
	UID                   string // uuid that will be used in communicating with client. This is firebase UID
	accountRoles          map[int]AccountRole
	AccountStatusID       int
	UserGenderID          int
	FirstName             string
	MiddleNames           string
	LastName              string
	DateOfBirth           time.Time
	Email                 string // can be used as login user name
	Phone                 string // for raw input
	DateTimeCreated       time.Time
	DateTimeModified      time.Time
	ToSAccepted           bool // users who do not accept ToS cannot proceed anything until accepted
	PrivacyPolicyAccepted bool
	ByGuardian            bool
	Signature             string
}

func (account Account) MeetMinimalRequirement() error {
	if len(account.FirstName) < 2 || len(account.LastName) < 2 {
		return errors.New("Name is too short")
	}
	if len(account.FirstName) > 18 || len(account.LastName) > 18 {
		return errors.New("Name is too long")
	}
	if len(account.Email) < 5 {
		return errors.New("Invalid email address")
	}
	if len(account.Phone) < 3 {
		return errors.New("Invalid phone number")
	}
	return nil
}

// SetRoles set a list of roles to Account
func (account *Account) SetRoles(roles []AccountRole) {
	account.accountRoles = make(map[int]AccountRole)
	for _, each := range roles {
		account.accountRoles[each.AccountTypeID] = each
	}
}

// HasRole checks if account has a particular role
func (account *Account) HasRole(roleID int) bool {
	if _, ok := account.accountRoles[roleID]; ok {
		return true
	}
	return false
}

// GetRoles returns all the roles that the caller account is associated with
func (account *Account) GetRoles() []int {
	roles := make([]int, 0)
	for k := range account.accountRoles {
		roles = append(roles, k)
	}
	return roles
}

// GetAccountRoles returns slice of AccountRole objects to caller
func (account Account) GetAccountRoles() []AccountRole {
	roles := make([]AccountRole, 0)
	for _, v := range account.accountRoles {
		roles = append(roles, v)
	}
	return roles
}

// IAccountRepository specifies the interface that an account repository needs to implement.
type IAccountRepository interface {
	SearchAccount(criteria SearchAccountCriteria) ([]Account, error)
	CreateAccount(account *Account) error
	UpdateAccount(account Account) error
	DeleteAccount(account Account) error
}

// SearchAccountCriteria provides the parameters that an IAccountRepository can use to search by
type SearchAccountCriteria struct {
	ID            int
	UUID          string
	Email         string
	Phone         string
	FirstName     string
	LastName      string
	DateOfBirth   time.Time
	Gender        int
	AccountType   int
	AccountStatus int
}

// FullName returns the full name of a user (excluding middle name, if any)
func (account Account) FullName() string {
	return account.FirstName + " " + account.LastName
}

// ICreateAccountStrategy specifies the interface that account creation strategy needs to implement.
type ICreateAccountStrategy interface {
	CreateAccount(account Account, password string) error
}

// GetAccountByEmail will retrieve account from repo by email. This function will return either a matched account
// or an empty account
func GetAccountByEmail(email string, repo IAccountRepository) Account {
	accounts, err := repo.SearchAccount(SearchAccountCriteria{
		Email: email,
	})
	if err != nil {
		return Account{}
	}
	if len(accounts) != 1 {
		return Account{}
	}
	return accounts[0]
}

// GetAccountByID will retrieve account from repo by ID. This function will return either a matched account
// or an empty account
func GetAccountByID(accountID int, repo IAccountRepository) Account {
	accounts, _ := repo.SearchAccount(SearchAccountCriteria{
		ID: accountID,
	})
	if len(accounts) != 1 {
		return Account{}
	}
	return accounts[0]
}

// GetAccountByUUID will retrieve account from repo by UID. This function will return either a matched account
// or an empty account
func GetAccountByUUID(uuid string, repo IAccountRepository) Account {
	accounts, _ := repo.SearchAccount(SearchAccountCriteria{
		UUID: uuid,
	})

	if len(accounts) != 1 {
		return Account{} // if no account is find, a null account will be returned
	}
	return accounts[0]
}
