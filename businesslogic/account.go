// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package businesslogic

import (
	"errors"
	"log"
	"time"

	"github.com/DancesportSoftware/das/util"
	"github.com/google/uuid"
)

// Account is the base account data for all users in DAS. Some fields are required with others are not
type Account struct {
	ID                    int    // userID will be account ID, too
	UUID                  string // uuid that will be used in communicating with client
	accountRoles          map[int]AccountRole
	AccountStatusID       int
	UserGenderID          int
	FirstName             string
	MiddleNames           string
	LastName              string
	DateOfBirth           time.Time
	Email                 string // can be used as login user name
	Phone                 string // for raw input
	EmailVerified         bool
	PhoneVerified         bool
	HashAlgorithm         string
	PasswordSalt          []byte // TODO: this should be refactored
	PasswordHash          []byte
	DateTimeCreated       time.Time
	DateTimeModified      time.Time
	ToSAccepted           bool // users who do not accept ToS cannot proceed anything until accepted
	PrivacyPolicyAccepted bool
	ByGuardian            bool
	Signature             string
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

// CreateAccountStrategy can create an account that only has record in IAccountRepository
type CreateAccountStrategy struct {
	AccountRepo          IAccountRepository
	RoleRepository       IAccountRoleRepository
	PreferenceRepository IUserPreferenceRepository
}

// CreateAccount creates a non-organizer account
func (strategy CreateAccountStrategy) CreateAccount(account Account, password string) error {
	if err := createAccount(&account, password, strategy.AccountRepo); err != nil {
		return err
	}

	// create default role for the account, which is athlete
	var createRoleErr error
	if strategy.RoleRepository != nil {
		defaultRole := AccountRole{
			AccountID:       account.ID,
			AccountTypeID:   AccountTypeAthlete,
			CreateUserID:    account.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    account.ID,
			DateTimeUpdated: time.Now(),
		}
		createRoleErr = strategy.RoleRepository.CreateAccountRole(&defaultRole)
	}
	if createRoleErr != nil {
		return createRoleErr
	}

	// initiate account preference data
	var createPrefErr error
	if strategy.PreferenceRepository != nil {
		defaultPreference := UserPreference{
			AccountID:       account.ID,
			DefaultRole:     AccountTypeAthlete,
			CreateUserID:    account.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    account.ID,
			DateTimeUpdated: time.Now(),
		}
		createPrefErr = strategy.PreferenceRepository.CreatePreference(&defaultPreference)
	}
	if createPrefErr != nil {
		return createPrefErr
	}

	return nil
}

// CreateOrganizerAccountStrategy creates an organizer account, which follows a different procedure from other accounts
type CreateOrganizerAccountStrategy struct {
	AccountRepo   IAccountRepository
	ProvisionRepo IOrganizerProvisionRepository
	HistoryRepo   IOrganizerProvisionHistoryRepository
}

// CreateAccount creates an organizer account
func (strategy CreateOrganizerAccountStrategy) CreateAccount(account Account, password string) error {
	if strategy.AccountRepo == nil {
		return errors.New("account repository is null")
	}
	if strategy.HistoryRepo == nil {
		return errors.New("organizer history repository is null")
	}
	if strategy.ProvisionRepo == nil {
		return errors.New("organizer repository is null")
	}
	if err := createAccount(&account, password, strategy.AccountRepo); err != nil {
		return err
	}
	provision, history := initializeOrganizerProvision(account.ID)
	if err := strategy.ProvisionRepo.CreateOrganizerProvision(&provision); err != nil {
		return err
	}
	if err := strategy.HistoryRepo.CreateOrganizerProvisionHistory(&history); err != nil {
		return err
	}
	return nil
}

func createAccount(account *Account, password string, repo IAccountRepository) error {
	if err := validateAccountRegistration(*account, repo); err != nil {
		return err
	}
	salt := util.GenerateSalt([]byte(password))
	hash := util.GenerateHash(salt, []byte(password))
	account.PasswordHash = hash
	account.PasswordSalt = salt
	account.UUID = uuid.New().String()
	account.HashAlgorithm = "SHA256"

	// TODO: email and phone verification should be performed before account can be activated
	account.AccountStatusID = AccountStatusActivated

	return repo.CreateAccount(account)
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

// GetAccountByUUID will retrieve account from repo by UUID. This function will return either a matched account
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

func checkEmailUsed(email string, repo IAccountRepository) bool {
	accounts, err := repo.SearchAccount(SearchAccountCriteria{
		Email: email,
	})
	if err != nil {
		log.Println(err)
	}
	if len(accounts) != 0 {
		return true
	}
	return false
}

// IAccountValidationStrategy specifies the function that should be implemented to be used to validate accounts that
// are about to be created
type IAccountValidationStrategy interface {
	Validate(account Account, accountRepo IAccountRepository) error
}

// validateAccountRegistration is for use
func validateAccountRegistration(account Account, accountRepo IAccountRepository) error {

	if checkEmailUsed(account.Email, accountRepo) {
		return errors.New("this email address is already used")
	}
	/*if (time.Now().Year() - account.DateOfBirth.Year()) > 120 {
		return errors.New("invalid date of birth")
	}
	if age.AgeAt(account.DateOfBirth, time.Now()) < 13 && !account.ByGuardian {
		return errors.New("must be 13 years old to register")
	}
	if age.AgeAt(account.DateOfBirth, time.Now()) < 13 && len(account.Signature) < 3 {
		return errors.New("must have consent from legal guardian")
	}*/
	if !account.ToSAccepted {
		return errors.New("terms of services must be accepted")
	}
	if !account.PrivacyPolicyAccepted {
		return errors.New("privacy policy must be accepted")
	}
	return nil
}
