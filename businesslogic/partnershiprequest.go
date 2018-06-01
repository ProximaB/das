package businesslogic

import (
	"errors"
	"time"
)

const (
	PARTNERSHIP_REQUEST_STATUS_ACCEPTED = 1
	PARTNERSHIP_REQUEST_STATUS_PENDING  = 2
	PARTNERSHIP_REQUEST_STATUS_DECLINED = 3
)

type PartnershipRequest struct {
	PartnershipRequestID int
	SenderID             int
	RecipientID          int
	SenderRole           string
	RecipientRole        string
	Message              string
	Status               int
	CreateUserID         int
	DateTimeCreated      time.Time
	UpdateUserID         int
	DateTimeUpdated      time.Time
}

type PartnershipRequestResponse struct {
	RequestID       int
	RecipientID     int
	Response        int
	DateTimeCreated time.Time
}

type SearchPartnershipRequestCriteria struct {
	RequestID       int `schema:"id"`
	Type            int `schema:"type"`
	Sender          int `schema:"sender"`
	Recipient       int `schema:"recipient"`
	RequestStatusID int `schema:"status"`
}

type IPartnershipRequestRepository interface {
	CreatePartnershipRequest(request *PartnershipRequest) error
	SearchPartnershipRequest(criteria *SearchPartnershipRequestCriteria) ([]PartnershipRequest, error)
	DeletePartnershipRequest(request PartnershipRequest) error
	UpdatePartnershipRequest(request PartnershipRequest) error
}

func (request PartnershipRequest) validateRoles() error {
	if request.SenderRole != PARTNERSHIP_ROLE_LEAD && request.SenderRole != PARTNERSHIP_ROLE_FOLLOW {
		return errors.New("sender's role is not specified")
	}
	if request.RecipientRole != PARTNERSHIP_ROLE_LEAD && request.RecipientRole != PARTNERSHIP_ROLE_FOLLOW {
		return errors.New("recipient's role is not specified")
	}
	if request.RecipientRole == request.SenderRole {
		return errors.New("sender and recipient have identical roles")
	}
	if request.SenderID == request.RecipientID {
		return errors.New("cannot send partnership request to yourself")
	}
	return nil
}

func (request PartnershipRequest) hasValidSenderAndRecipient(accountRepo IAccountRepository) error {
	senderAccounts, seErr := accountRepo.SearchAccount(&SearchAccountCriteria{ID: request.SenderID})
	recipientAccounts, recErr := accountRepo.SearchAccount(&SearchAccountCriteria{ID: request.RecipientID})
	if seErr != nil {
		return seErr
	}
	if recErr != nil {
		return recErr
	}
	if len(senderAccounts) != 1 {
		return errors.New("sender account cannot be found")
	}
	if len(recipientAccounts) != 1 {
		return errors.New("recipient account cannot be found")
	}
	if senderAccounts[0].AccountTypeID != ACCOUNT_TYPE_ATHLETE {
		return errors.New("sender is not an athlete")
	}
	if recipientAccounts[0].AccountTypeID != ACCOUNT_TYPE_ATHLETE {
		return errors.New("recipient is not an athlete")
	}
	return nil
}

func (request PartnershipRequest) senderBlockedByRecipient(blacklistRepo IPartnershipRequestBlacklistRepository) bool {
	recipientBlacklist, _ := blacklistRepo.SearchPartnershipRequestBlacklist(SearchPartnershipRequestBlacklistCriteria{ReporterID: request.RecipientID})
	for _, each := range recipientBlacklist {
		if each.BlockedUserID == request.SenderID {
			return true
		}
	}
	return false
}

// hasExistingPartnership checks if there is already a partnership between the two dancers
func (request PartnershipRequest) hasExistingPartnership(accountRepo IAccountRepository, partnershipRepo IPartnershipRepository) bool {
	// configure search partnershipCriteria
	senderAccount := GetAccountByID(request.SenderID, accountRepo)
	recipientAccount := GetAccountByID(request.RecipientID, accountRepo)

	partnershipCriteria := new(SearchPartnershipCriteria)
	if request.SenderRole == PARTNERSHIP_ROLE_LEAD {
		partnershipCriteria.LeadID = senderAccount.ID
		partnershipCriteria.FollowID = recipientAccount.ID
	} else {
		partnershipCriteria.FollowID = senderAccount.ID
		partnershipCriteria.LeadID = recipientAccount.ID
	}

	// check if sender is already in a partnership with recipient
	partnerships, _ := partnershipRepo.SearchPartnership(partnershipCriteria)
	if len(partnerships) != 0 {
		return true
	}
	return false
}

// hasPendingRequest checks if there is a request between these two dancers that still waits for response
func (request PartnershipRequest) hasPendingRequest(requestRepo IPartnershipRequestRepository) bool {
	// check if there is pending message between sender and recipient
	requests, _ := requestRepo.SearchPartnershipRequest(&SearchPartnershipRequestCriteria{
		Recipient:       request.RecipientID,
		Sender:          request.SenderID,
		RequestStatusID: PARTNERSHIP_REQUEST_STATUS_PENDING,
	})
	if len(requests) == 1 {
		return true
	}
	return false
}

// CreatePartnershipRequest will create the partnership request with validation. Validation includes
// 1. Role validation: must be opposite role
// 2. Blacklist check: sender must not be blacklisted by recipient
// 3. Existing partnership check: sender and recipient must not be in a partnership with specified role
// 4. There is no pending request for the same role (this is applied to request from either party)
// Note: if sender and recipient are in a partnership of opposite role, then it's considered as a different partnership.
// If the request is valid, then request will be created.
func CreatePartnershipRequest(request PartnershipRequest, partnershipRepo IPartnershipRepository,
	requestRepo IPartnershipRequestRepository, accountRepo IAccountRepository,
	blacklistRepo IPartnershipRequestBlacklistRepository) error {

	// validate Roles the request first
	if roleErr := request.validateRoles(); roleErr != nil {
		return roleErr
	}

	// check if accounts exist
	if accountErr := request.hasValidSenderAndRecipient(accountRepo); accountErr != nil {
		return accountErr
	}

	// check if sender is blacklisted by recipient
	if request.senderBlockedByRecipient(blacklistRepo) {
		return errors.New("cannot send partnership request to this user")
	}

	if request.hasExistingPartnership(accountRepo, partnershipRepo) {
		return errors.New("you are already in a partnership with specified role")
	}

	if request.hasPendingRequest(requestRepo) {
		return errors.New("a pending request must be responded first")
	}

	return requestRepo.CreatePartnershipRequest(&request)
}

// verify if the response to a partnership request is valid
func validatePartnershipRequestResponse(response PartnershipRequestResponse, repo IPartnershipRequestRepository) error {
	if response.RecipientID == 0 {
		return errors.New("recipient must be specified")
	}

	if response.RequestID == 0 {
		return errors.New("request must be specified")
	}

	// check if request is valid
	if requests, searchErr := repo.SearchPartnershipRequest(&SearchPartnershipRequestCriteria{
		RequestID: response.RequestID,
		Recipient: response.RecipientID,
	}); searchErr != nil {
		return searchErr
	} else if len(requests) != 1 {
		return errors.New("cannot find request for this recipient")
	} else if requests[0].Status == PARTNERSHIP_REQUEST_STATUS_ACCEPTED || requests[0].Status == PARTNERSHIP_REQUEST_STATUS_DECLINED {
		return errors.New("this request is already responded")
	}

	return nil
}

func RespondPartnershipRequest(response PartnershipRequestResponse,
	requestRepo IPartnershipRequestRepository,
	accountRepo IAccountRepository,
	partnershipRepo IPartnershipRepository) error {

	if validErr := validatePartnershipRequestResponse(response, requestRepo); validErr != nil {
		return validErr
	}

	// respond partnership
	if response.Response == PARTNERSHIP_REQUEST_STATUS_ACCEPTED || response.Response == PARTNERSHIP_REQUEST_STATUS_DECLINED {
		requests, err := requestRepo.SearchPartnershipRequest(&SearchPartnershipRequestCriteria{
			RequestID: response.RequestID,
			Recipient: response.RecipientID,
		})
		if err != nil {
			return err
		}
		requests[0].DateTimeUpdated = time.Now()
		requests[0].Status = response.Response
		if respErr := requestRepo.UpdatePartnershipRequest(requests[0]); respErr != nil {
			return respErr
		}

		// optional: create partnership if accepted
		if response.Response == PARTNERSHIP_REQUEST_STATUS_ACCEPTED {
			partnership := Partnership{}
			requests, _ := requestRepo.SearchPartnershipRequest(&SearchPartnershipRequestCriteria{
				RequestID: response.RequestID,
				Recipient: response.RecipientID,
			})
			request := requests[0]

			if request.RecipientRole == PARTNERSHIP_ROLE_LEAD {
				partnership.LeadID = request.RecipientID
				partnership.FollowID = request.SenderID
			} else {
				partnership.LeadID = request.SenderID
				partnership.FollowID = request.RecipientID
			}

			leadAccount := GetAccountByID(partnership.LeadID, accountRepo)
			followAccount := GetAccountByID(partnership.FollowID, accountRepo)
			if leadAccount.UserGenderID == followAccount.UserGenderID {
				partnership.SameSex = true
			} else {
				partnership.SameSex = false
			}

			partnership.DateTimeCreated = time.Now()
			return partnershipRepo.CreatePartnership(partnership)
		}
	}
	return nil
}
