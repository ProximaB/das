package account

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

// ProfileSearchServer defines a virtual server that handles requests from search user profiles, such as competitors,
// couples, competition officials (adjudicators, scrutineers, organizers, etc).
type ProfileSearchServer struct {
	accountRepository businesslogic.IAccountRepository
}

func NewProfileSearchServer(accountRepo businesslogic.IAccountRepository) ProfileSearchServer {
	if accountRepo == nil {
		log.Fatal("[fatal] provided IAccountRepository for ProfileSearchServer is nil")
	}
	return ProfileSearchServer{
		accountRepository: accountRepo,
	}
}

// SearchDancerProfileHandler handles the request
//	GET /api/v1.0/profile/dancers
func (server ProfileSearchServer) SearchDancerProfileHandler(w http.ResponseWriter, r *http.Request) {
	form := new(viewmodel.SearchAthleteProfileForm)
	if err := util.ParseRequestData(r, form); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, err.Error())
		return
	}

	if validationErr := validator.Validate(form); validationErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, validationErr.Error())
		return
	}
}

func (server ProfileSearchServer) SearchPartnershipProfileHandler(w http.ResponseWriter, r *http.Request) {
	form := new(viewmodel.SearchPartnershipProfileForm)
	if err := util.ParseRequestData(r, form); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, err.Error())
		return
	}

	if validationErr := validator.Validate(form); validationErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, validationErr.Error())
		return
	}
}
