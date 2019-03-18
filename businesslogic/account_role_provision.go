package businesslogic

import (
	"errors"
	"log"
	"time"
)

const (
	// RoleApplicationStatusApproved marks the role application as "approved"
	RoleApplicationStatusApproved = 1
	// RoleApplicationStatusDenied marks the role application as "denied"
	RoleApplicationStatusDenied = 2
	// RoleApplicationStatusPending marks the role application as "pending"
	RoleApplicationStatusPending = 3
)

var notAuthorizedToApproveUserRoleApplicationError = errors.New("not authorized to approve user's role application")

// SearchRoleApplicationCriteria specifies the search criteria for role application
type SearchRoleApplicationCriteria struct {
	ID             int
	AccountID      int
	AppliedRoleID  int
	StatusID       int
	ApprovalUserID int
	Responded      bool
}

// RoleApplication is an application for restricted roles, including adjudicator, scrutineer, and organizer.
// Non-restrictive roles such as emcee and deck captain can be approved by competition organizers
type RoleApplication struct {
	ID               int
	AccountID        int
	Account          Account
	AppliedRoleID    int
	Description      string
	StatusID         int
	ApprovalUserID   *int
	DateTimeApproved time.Time
	CreateUserID     int
	DateTimeCreated  time.Time
	UpdateUserID     int
	DateTimeUpdated  time.Time
}

// IRoleApplicationRepository specifies the interface that a Role Application Repository should implement
type IRoleApplicationRepository interface {
	CreateApplication(application *RoleApplication) error
	SearchApplication(criteria SearchRoleApplicationCriteria) ([]RoleApplication, error)
	UpdateApplication(application RoleApplication) error
}

// RoleProvisionService is a service that handles Role Application and provision
type RoleProvisionService struct {
	accountRepo               IAccountRepository
	roleApplicationRepo       IRoleApplicationRepository
	roleRepo                  IAccountRoleRepository
	organizerProvisionService OrganizerProvisionService
}

// NewRoleProvisionService create a service that serves Role Provision
func NewRoleProvisionService(
	accountRepo IAccountRepository,
	roleApplicationRepo IRoleApplicationRepository,
	roleRepo IAccountRoleRepository,
	organizerProvisionRepo IOrganizerProvisionRepository,
	organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository) *RoleProvisionService {
	service := RoleProvisionService{
		accountRepo:               accountRepo,
		roleApplicationRepo:       roleApplicationRepo,
		roleRepo:                  roleRepo,
		organizerProvisionService: NewOrganizerProvisionService(accountRepo, roleRepo, organizerProvisionRepo, organizerProvisionHistoryRepo),
	}
	return &service
}

// CreateRoleApplication check the validity of the role application and create it if it's valid
func (service RoleProvisionService) CreateRoleApplication(currentUser Account, application *RoleApplication) error {
	// check if current user has the role
	if currentUser.HasRole(application.AppliedRoleID) {
		return errors.New("current user already has the applied role")
	}

	// check if has a pending application
	searchResults, err := service.roleApplicationRepo.SearchApplication(SearchRoleApplicationCriteria{
		AccountID:     currentUser.ID,
		AppliedRoleID: application.AppliedRoleID,
		StatusID:      RoleApplicationStatusPending,
	})
	if err != nil {
		return err
	}
	if len(searchResults) != 0 {
		return errors.New("previous application has not been responded")
	}

	// check what role that user is applying for
	if application.AppliedRoleID == AccountTypeAthlete {
		return errors.New("athlete role should be granted when the account was created")
	}
	if application.AppliedRoleID > AccountTypeEmcee {
		return errors.New("invalid role")
	}

	return service.roleApplicationRepo.CreateApplication(application)
}

func (service RoleProvisionService) respondRoleApplication(currentUser Account, application *RoleApplication, action int) error {
	application.StatusID = action
	application.ApprovalUserID = &currentUser.ID
	application.DateTimeApproved = time.Now()
	if updateErr := service.roleApplicationRepo.UpdateApplication(*application); updateErr != nil {
		return updateErr
	}
	if action == RoleApplicationStatusApproved {
		role := AccountRole{
			AccountID:       application.AccountID,
			AccountTypeID:   application.AppliedRoleID,
			CreateUserID:    currentUser.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    currentUser.ID,
			DateTimeUpdated: time.Now(),
		}
		return service.roleRepo.CreateAccountRole(&role)
	}
	return nil
}

// UpdateApplication attempts to approve the Role application based on the privilege of current user.
// If current user is admin, any application can be approved
// If current user is organizer, only emcee and deck-captain can be approved
// If current user is other roles, current user will be prohibited from performing such action
func (service RoleProvisionService) UpdateApplication(currentUser Account, application *RoleApplication, action int) error {
	// check if action is valid
	if !(action == RoleApplicationStatusApproved || action == RoleApplicationStatusDenied) {
		return errors.New("invalid response to role application")
	}
	// check if application is pending
	if application.StatusID == RoleApplicationStatusApproved || application.StatusID == RoleApplicationStatusDenied {
		return errors.New("role application is already responded")
	}
	// Only an Admin or Organizer user ca update user's role application
	if !(currentUser.HasRole(AccountTypeOrganizer) || currentUser.HasRole(AccountTypeAdministrator)) {
		return errors.New("unauthorized")
	}
	// should not allow users to provision themselves other than Admin
	if currentUser.ID == application.AccountID && !currentUser.HasRole(AccountTypeAdministrator) {
		return errors.New("not authorized to provision your own role application")
	}
	switch application.AppliedRoleID {
	case AccountTypeAthlete:
		return nil // Athlete role does not need to be provisioned
	case AccountTypeAdjudicator:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return notAuthorizedToApproveUserRoleApplicationError
		}
	case AccountTypeScrutineer:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return notAuthorizedToApproveUserRoleApplicationError
		}
	case AccountTypeOrganizer:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return notAuthorizedToApproveUserRoleApplicationError
		}
	case AccountTypeDeckCaptain:
		if !(currentUser.HasRole(AccountTypeAdministrator) || currentUser.HasRole(AccountTypeOrganizer)) {
			return notAuthorizedToApproveUserRoleApplicationError
		}
	case AccountTypeEmcee:
		if !(currentUser.HasRole(AccountTypeAdministrator) || currentUser.HasRole(AccountTypeOrganizer)) {
			return notAuthorizedToApproveUserRoleApplicationError
		}
	default:
		return errors.New("invalid role application")
	}
	roleProvisionErr := service.respondRoleApplication(currentUser, application, action)
	if roleProvisionErr != nil {
		return roleProvisionErr
	}

	if application.AppliedRoleID == AccountTypeOrganizer {
		roleSearch, roleSearchErr := service.roleRepo.SearchAccountRole(SearchAccountRoleCriteria{
			AccountID:     application.AccountID,
			AccountTypeID: AccountTypeOrganizer,
		})
		if roleSearchErr != nil || len(roleSearch) != 1 {
			if roleSearchErr != nil {
				log.Println(roleSearchErr)
			}
			return errors.New("cannot find Organizer role of this account")
		}
		role := roleSearch[0]

		// create organizer provision
		provisionEntry, initErr := service.organizerProvisionService.NewOrganizerProvision(role.ID, currentUser.ID)
		if initErr != nil {
			return initErr
		}
		provisionHistoryEntry := service.organizerProvisionService.NewInitialOrganizerProvisionHistoryEntry(role.ID, currentUser.ID)
		if orgProvErr := service.organizerProvisionService.organizerProvisionRepo.CreateOrganizerProvision(&provisionEntry); orgProvErr != nil {
			return orgProvErr
		}
		if orgProvHistErr := service.organizerProvisionService.organizerProvisionHistoryRepo.CreateOrganizerProvisionHistory(&provisionHistoryEntry); orgProvHistErr != nil {
			return orgProvHistErr
		}
	}
	return nil
}

// SearchRoleApplication searches the available role application based on current user's privilege
func (service RoleProvisionService) SearchRoleApplication(currentUser Account, criteria SearchRoleApplicationCriteria) ([]RoleApplication, error) {
	return service.roleApplicationRepo.SearchApplication(criteria)
}
