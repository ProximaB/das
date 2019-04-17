package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

const (
	CompetitionStatusPreRegistration    = 1
	CompetitionStatusOpenRegistration   = 2
	CompetitionStatusClosedRegistration = 3
	CompetitionStatusInProgress         = 4
	CompetitionStatusProcessing         = 5
	CompetitionStatusClosed             = 6
	CompetitionStatusCancelled          = 7
)

// CompetitionStatus defines the data that is required to label the status of a Competition
type CompetitionStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// ICompetitionStatusRepository defines the function that a CompetitionStatusRepository should implement
type ICompetitionStatusRepository interface {
	GetCompetitionAllStatus() ([]CompetitionStatus, error)
}

// Competition provides the base data structure for a competitive ballroom dance. All competitions in
// DAS must have some affiliation with a dancesport federation (Not Affiliated/Independent is also a Federation,)
type Competition struct {
	ID                        int
	FederationID              int
	Name                      string
	Street                    string
	City                      City
	State                     State
	Country                   Country
	StartDateTime             time.Time
	EndDateTime               time.Time
	CreateUserID              int
	DateTimeCreated           time.Time
	UpdateUserID              int
	DateTimeUpdated           time.Time
	ContactName               string
	ContactEmail              string
	ContactPhone              string
	statusID                  int
	Website                   string
	Attendance                int
	RegistrationOpenDateTime  time.Time
	RegistrationCloseDateTime time.Time
	officials                 map[int][]Account // the integer key is the type of official: adjudicator, scrutineer, emcee, and deck captain
}

// UpdateStatus will attempt to change the status of the caller competition to statusID, if the change is in logical order
func (comp *Competition) UpdateStatus(statusID int) error {
	if comp.statusID >= statusID && comp.statusID != 0 {
		return errors.New("cannot revert competition status")
	}
	if comp.statusID == CompetitionStatusClosed || comp.statusID == CompetitionStatusCancelled {
		return errors.New("competition status is locked")
	}
	comp.statusID = statusID
	return nil
}

func (comp *Competition) AddOfficial(role int, official Account) error {
	if comp.officials == nil {
		comp.officials = make(map[int][]Account)
	}

	if !official.HasRole(role) {
		return errors.New("This user does not have the appointed role provisioned.")
	}

	if comp.officials[role] == nil {
		comp.officials[role] = make([]Account, 0)
	}

	// check if this official is already added
	for _, each := range comp.officials[role] {
		if each.ID == official.ID {
			return errors.New("This user is already added to the list of officials")
		}
	}
	comp.officials[role] = append(comp.officials[role], official)

	return nil
}

func (comp Competition) GetOfficials(role int) ([]Account, error) {
	if comp.officials == nil {
		return nil, errors.New("No official has been assigned to this competition")
	}
	return comp.officials[role], nil
}

func (comp Competition) GetStatus() int {
	return comp.statusID
}

// SearchCompetitionCriteria specifies the parameters that can be used to search a Competition
type SearchCompetitionCriteria struct {
	ID            int       `schema:"id"`
	Name          string    `schema:"name"`
	FederationID  int       `schema:"federation"`
	StateID       int       `schema:"state"`
	CountryID     int       `schema:"country"`
	StartDateTime time.Time `schema:"start"`
	OrganizerID   int
	StatusID      int `schema:"status"`
}

type OrganizerUpdateCompetition struct {
	CompetitionID int       `json:"competitionId"`
	Name          string    `json:"name"`
	Website       string    `json:"website"`
	Status        int       `json:"statusId"`
	Address       string    `json:"street"`
	CityID        int       `json:"cityId"`
	StateID       int       `json:"stateId"`
	CountryID     int       `json:"countryId"`
	FederationID  int       `json:"federationId"`
	ContactName   string    `json:"contact"`
	ContactEmail  string    `json:"email"`
	ContactPhone  string    `json:"phone"`
	StartDate     time.Time `json:"start"`
	EndDate       time.Time `json:"end"`
	UpdateUserID  int
}

// ICompetitionRepository specifies the interface that a competition repository needs to implement to provide CRUD
// operations in the data repository
type ICompetitionRepository interface {
	CreateCompetition(competition *Competition) error
	SearchCompetition(criteria SearchCompetitionCriteria) ([]Competition, error)
	UpdateCompetition(competition Competition) error
	DeleteCompetition(competition Competition) error
}

// CreateCompetition creates competition in competitionRepo, update records in provisionRepo, and
// add a new record to historyRepo
func CreateCompetition(competition Competition, competitionRepo ICompetitionRepository,
	provisionRepo IOrganizerProvisionRepository, historyRepo IOrganizerProvisionHistoryRepository) error {
	// check if data received is validationErr
	if validationErr := competition.validateCreateCompetition(); validationErr != nil {
		return validationErr
	}

	if competition.statusID == 0 {
		competition.statusID = CompetitionStatusPreRegistration
	}

	// check if organizer is provisioned with available competitions
	provisions, _ := provisionRepo.SearchOrganizerProvision(SearchOrganizerProvisionCriteria{
		OrganizerID: competition.CreateUserID,
	})
	if len(provisions) != 1 {
		return errors.New("no organizer record is found")
	}
	provision := provisions[0]
	if provision.Available < 1 {
		return errors.New("no available competition slot")
	}

	newProvision := provision.updateForCreateCompetition(competition)
	historyEntry := newProvisionHistoryEntry(newProvision, competition)
	updateOrganizerProvision(newProvision, historyEntry, provisionRepo, historyRepo)

	err := competitionRepo.CreateCompetition(&competition)
	if err != nil {
		// refund competition organizer's provision
		refundProvision := newProvision
		refundProvision.Available += 1
		refundProvision.Hosted -= 1

		refundEntry := OrganizerProvisionHistoryEntry{
			OrganizerRoleID: refundProvision.OrganizerRoleID,
			Amount:          1,
			Note:            fmt.Sprintf("Refund for failing in creating competition %v %v", competition.Name, competition.StartDateTime),
			CreateUserID:    competition.CreateUserID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    competition.UpdateUserID,
			DateTimeUpdated: time.Now(),
		}
		updateOrganizerProvision(refundProvision, refundEntry, provisionRepo, historyRepo)
	}
	return err
}

func (comp Competition) validateCreateCompetition() error {
	if comp.FederationID < 1 {
		return errors.New("invalid federation")
	}
	if len(comp.Name) < 3 {
		return errors.New("comp name is too short")
	}
	if len(comp.Website) < 7 { // requires "http://"
		return errors.New("official comp website is required")
	}
	if comp.GetStatus() > CompetitionStatusClosedRegistration {
		return errors.New("cannot create comp that no longer allows new registration")
	}
	if comp.StartDateTime.After(comp.EndDateTime) {
		return errors.New("start date must be ahead of end date")
	}
	if comp.StartDateTime.Before(time.Now()) {
		return errors.New("comp must starts in a future time")
	}
	if comp.StartDateTime.After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("cannot create far-future comp")
	}
	if len(comp.ContactName) < 3 {
		return errors.New("contact name is too short")
	}
	if len(comp.ContactEmail) < 5 {
		return errors.New("contact email is too short")
	}
	if len(comp.ContactPhone) < 9 {
		return errors.New("contact phone is too short")
	}
	if len(comp.Street) < 4 {
		return errors.New("street address is too short")
	}
	if comp.City.ID < 1 {
		return errors.New("city is required")
	}
	if comp.State.ID < 1 {
		return errors.New("state is required")
	}
	if comp.Country.ID < 1 {
		return errors.New("country is required")
	}
	if comp.CreateUserID < 1 {
		return errors.New("unauthorized")
	}
	if comp.UpdateUserID < 1 {
		return errors.New("unauthorized")
	}
	return nil
}

// UpdateCompetition updates the existing competition with the information provided in OrganizerUpdateCompetition
func UpdateCompetition(user *Account, competition OrganizerUpdateCompetition, repo ICompetitionRepository) error {
	// check if user is the owner of the original competition
	competitions, err := repo.SearchCompetition(SearchCompetitionCriteria{ID: competition.CompetitionID})
	if err != nil {
		return err
	}
	if competitions == nil || len(competitions) != 1 || competitions[0].ID == 0 {
		return errors.New("cannot find this competition")
	}
	if validationErr := validateUpdateCompetition(user, competitions[0], &competition); validationErr != nil {
		return validationErr
	}

	if competitions[0].GetStatus() == CompetitionStatusOpenRegistration ||
		competitions[0].GetStatus() == CompetitionStatusClosedRegistration {
		// TODO: reimplement event update
		/*if updateEventErr := dataaccess.UpdateCompetitionEventStatus(dataaccess.DATABASE, competition.ID, competitions[0].StatusID); updateEventErr != nil {
			return updateEventErr
		}*/
	}

	// if competition is cancelled, refund the slot. competition cannot be cancelled unless it is done by site administrator

	return repo.UpdateCompetition(competitions[0])
}

func validateUpdateCompetition(user *Account,
	competition Competition,
	updateDTO *OrganizerUpdateCompetition) error {
	if user.ID != competition.CreateUserID {
		return errors.New("not authorized to update this competition")
	}

	if competition.GetStatus() > updateDTO.Status {
		return errors.New("cannot change competition status back")
	}

	if competition.GetStatus() == CompetitionStatusClosed {
		return errors.New("competition is closed")
	}
	if len(updateDTO.Name) < 3 {
		return errors.New("invalid competition name")
	}
	if len(updateDTO.Website) < 3 {
		// TODO: need a better url validation mechanics
		return errors.New("website link is too short")
	}
	if updateDTO.StartDate.After(updateDTO.EndDate) {
		return errors.New("competition must start before it ends")
	}

	if updateDTO.StartDate.Before(time.Now()) {
		return errors.New("cannot start competition in the past")
	}

	if updateDTO.StartDate.After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("cannot create competition that starts a year later")
	}

	return nil
}

type IEventMetaRepository interface {
	GetEventUniqueFederations(competition Competition) ([]Federation, error)
	GetEventUniqueDivisions(competition Competition) ([]Division, error)
	GetEventUniqueAges(competition Competition) ([]Age, error)
	GetEventUniqueProficiencies(competition Competition) ([]Proficiency, error)
	GetEventUniqueStyles(competition Competition) ([]Style, error)
}

// Get a list of unique federations that a competition has
func (comp Competition) GetEventUniqueFederations(eventRepository IEventMetaRepository) ([]Federation, error) {
	return eventRepository.GetEventUniqueFederations(comp)
}
func (comp Competition) GetEventUniqueDivisions(eventRepository IEventMetaRepository) ([]Division, error) {
	return eventRepository.GetEventUniqueDivisions(comp)
}
func (comp Competition) GetEventUniqueAges(eventRepository IEventMetaRepository) ([]Age, error) {
	return eventRepository.GetEventUniqueAges(comp)
}
func (comp Competition) GetEventUniqueProficiencies(eventRepository IEventMetaRepository) ([]Proficiency, error) {
	return eventRepository.GetEventUniqueProficiencies(comp)
}
func (comp Competition) GetEventUniqueStyles(eventRepository IEventMetaRepository) ([]Style, error) {
	return eventRepository.GetEventUniqueStyles(comp)
}

type CompetitionOfficial struct {
	ID              int
	Competition     Competition
	Official        Account   // the ID for AccountRole
	OfficialRoleID  int       // the ID for AccountType
	EffectiveFrom   time.Time // have privileged access to competition data
	EffectiveUntil  time.Time
	AssignedBy      int // ID of an AccountRole object, must be an organizer. TODO: may use invitation instead of assignment
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// Active checks if the status of this position is still active.
func (official CompetitionOfficial) Active() bool {
	return time.Now().Before(official.EffectiveUntil) && time.Now().After(official.EffectiveFrom)
}

type SearchCompetitionOfficialCriteria struct {
	ID             int
	CompetitionID  int
	OfficialID     string
	OfficialRoleID int
}

type ICompetitionOfficialRepository interface {
	CreateCompetitionOfficial(official *CompetitionOfficial) error
	DeleteCompetitionOfficial(official CompetitionOfficial) error
	SearchCompetitionOfficial(criteria SearchCompetitionOfficialCriteria) ([]CompetitionOfficial, error)
	UpdateCompetitionOfficial(official CompetitionOfficial) error
}

const (
	COMPETITION_INVITATION_STATUS_ACCEPTED = "Accepted"
	COMPETITION_INVITATION_STATUS_REJECTED = "Rejected"
	COMPETITION_INVITATION_STATUS_PENDING  = "Pending"
	COMPETITION_INVITATION_STATUS_REVOKED  = "Revoked"
	COMPETITION_INVITATION_STATUS_EXPIRED  = "Expired"
)

// CompetitionOfficialInvitation is an invitation that can only be sent by organizers to recipients who have provisioned
// competition official role.
type CompetitionOfficialInvitation struct {
	ID                 int
	Sender             Account
	Recipient          Account
	ServiceCompetition Competition // the competition that the recipient will serve at if accepted
	AssignedRoleID     int         // only allow Adjudicator, Scrutineer, Deck Captain, Emcee
	Message            string
	InvitationStatus   string
	ExpirationDate     time.Time
	CreateUserID       int
	DateTimeCreated    time.Time
	UpdateUserID       int
	DateTimeUpdated    time.Time
}

// SearchCompetitionOfficialInvitationCriteria defines the search criteria that can be used to search CompetitionOfficialInvitation
// within a Invitation repository.
type SearchCompetitionOfficialInvitationCriteria struct {
	SenderID             int
	RecipientID          int
	ServiceCompetitionID int
	AssignedRoleID       int
	Status               string
	CreateUserID         int
	UpdateUserID         int
}

// ICompetitionOfficialInvitationRepository defines the interface that a CompetitionOfficialInvitation repository should
// implement, including creating, deleting, searching, and updating the invitation.
type ICompetitionOfficialInvitationRepository interface {
	CreateCompetitionOfficialInvitationRepository(invitation *CompetitionOfficialInvitation) error
	DeleteCompetitionOfficialInvitationRepository(invitation CompetitionOfficialInvitation) error
	SearchCompetitionOfficialInvitationRepository(criteria SearchCompetitionOfficialInvitationCriteria) ([]CompetitionOfficialInvitation, error)
	UpdateCompetitionOfficialInvitationRepository(invitation CompetitionOfficialInvitation) error
}

// CompetitionOfficialInvitationService defines the service that can provide search
type CompetitionOfficialInvitationService struct {
	accountRepo     IAccountRepository
	competitionRepo ICompetitionRepository
	officialRepo    ICompetitionOfficialRepository
	invitationRepo  ICompetitionOfficialInvitationRepository
}

func NewCompetitionOfficialInvitationService(
	accountRepo IAccountRepository,
	competitionRepo ICompetitionRepository,
	officialRep ICompetitionOfficialRepository,
	invitationRepo ICompetitionOfficialInvitationRepository) CompetitionOfficialInvitationService {
	return CompetitionOfficialInvitationService{
		accountRepo:     accountRepo,
		competitionRepo: competitionRepo,
		officialRepo:    officialRep,
		invitationRepo:  invitationRepo,
	}
}

func (service CompetitionOfficialInvitationService) SearchCompetitionOfficialInvitation(criteria SearchCompetitionOfficialInvitationCriteria) ([]CompetitionOfficialInvitation, error) {
	return service.invitationRepo.SearchCompetitionOfficialInvitationRepository(criteria)
}

func (service CompetitionOfficialInvitationService) CreateCompetitionOfficialInvitation(sender, recipient Account, serviceRole int, comp Competition) (CompetitionOfficialInvitation, error) {
	invitation := CompetitionOfficialInvitation{}

	// sender must be the creator of the competition
	if sender.ID != comp.CreateUserID {
		return invitation, errors.New("Not authorized to send competition official invitation.")
	}

	invitation.Sender = sender

	// competition must be prior to running
	if comp.GetStatus() >= CompetitionStatusInProgress {
		return invitation, errors.New("Competition is already running and no more officials can be assigned.")
	}

	invitation.ServiceCompetition = comp

	// recipient must already have the request service role
	if !recipient.HasRole(serviceRole) {
		return invitation, errors.New("Recipient does not have this role provisioned by Administrator.")
	}

	invitation.Recipient = recipient
	invitation.AssignedRoleID = serviceRole
	invitation.DateTimeCreated = time.Now()
	invitation.CreateUserID = sender.ID
	invitation.DateTimeUpdated = time.Now()
	invitation.UpdateUserID = sender.ID

	// invitation will expire either:
	// - 30 days after the invitation, or
	// - after the competition
	thirtyDayLimit := time.Now().AddDate(0, 0, 30)
	if thirtyDayLimit.Before(comp.EndDateTime) {
		invitation.ExpirationDate = thirtyDayLimit
	} else {
		invitation.ExpirationDate = comp.EndDateTime
	}

	// initialize invitation status to pending
	invitation.InvitationStatus = COMPETITION_INVITATION_STATUS_PENDING

	// create the role invitation
	createErr := service.invitationRepo.CreateCompetitionOfficialInvitationRepository(&invitation)

	// TODO: send notification to recipient (requires notification)

	return invitation, createErr
}

func (service CompetitionOfficialInvitationService) UpdateCompetitionOfficialInvitation(currentUser Account, invitation CompetitionOfficialInvitation, response string) error {
	// only the sender and the recipient can make changes to the invitation
	if currentUser.ID != invitation.Sender.ID && currentUser.ID != invitation.Recipient.ID {
		return errors.New("Not authorized to make changes to this invitation")
	}

	// check terminal status
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_EXPIRED {
		return errors.New("Invitation is expired and can no longer be updated")
	}
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_REVOKED {
		return errors.New("Invitation is revoked and can no longer be updated")
	}
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_REJECTED {
		return errors.New("Invitation is rejected and can no longer be updated")
	}

	// for no-terminal status: pending and accepted
	// pending request can be updated by:
	// - recipient to accept/reject
	// - sender to revoke
	// accepted request can be updated by:
	// - recipient to reject
	// - sender to revoke
	canUpdate := false
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_PENDING {
		if currentUser.ID == invitation.Recipient.ID {
			// can accept or reject
			if response != COMPETITION_INVITATION_STATUS_ACCEPTED && response != COMPETITION_INVITATION_STATUS_REJECTED {
				return errors.New("The invitation can only be accepted or rejected.")
			} else {
				canUpdate = true
			}
		} else if currentUser.ID == invitation.Sender.ID {
			if response != COMPETITION_INVITATION_STATUS_REVOKED {
				return errors.New("The invitation can only be revoked")
			} else {
				canUpdate = true
			}
		}

	} else if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_ACCEPTED {
		if currentUser.ID == invitation.Recipient.ID {
			if response != COMPETITION_INVITATION_STATUS_REJECTED {
				return errors.New("The invitation can only be rejected.")
			}
		} else if currentUser.ID == invitation.Sender.ID {
			if response != COMPETITION_INVITATION_STATUS_REVOKED {
				return errors.New("The invitation can only be revoked")
			}
			canUpdate = true
		}
	}

	if canUpdate {
		invitation.InvitationStatus = response
		invitation.DateTimeUpdated = time.Now()
		invitation.UpdateUserID = currentUser.ID
		return service.invitationRepo.UpdateCompetitionOfficialInvitationRepository(invitation)
	}
	return errors.New("An unknown error occurred while processing this invitation. Please report this incident to site administrator.")
}
