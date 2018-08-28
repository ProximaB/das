package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
	"time"
)

type RoleApplicationServer struct {
	auth    authentication.IAuthenticationStrategy
	service businesslogic.RoleProvisionService
}

func NewRoleApplicationServer(authStrat authentication.IAuthenticationStrategy, service businesslogic.RoleProvisionService) RoleApplicationServer {
	return RoleApplicationServer{
		auth:    authStrat,
		service: service,
	}
}

// CreateRoleApplicationHandler handles the request:
//	POST /api/v1.0/account/role/apply
// Accepted JSON payload:
//	{
//		"role": 2,
//		"description": "I am applying this role because..."
//	}
// Sample response payload:
//	{
//		"status": 200,
//		"message":
//		"data": null
//	}
func (server RoleApplicationServer) CreateRoleApplicationHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, authErr := server.auth.GetCurrentUser(r)
	if authErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	applicationDTO := new(viewmodel.SubmitRoleApplication)
	parseErr := util.ParseRequestBodyData(r, applicationDTO)

	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad application data, please try again", nil)
		return
	}
	if applicationDTO.Validate() != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, applicationDTO.Validate().Error(), nil)
		return
	}

	application := businesslogic.RoleApplication{
		AccountID:       currentUser.ID,
		AppliedRoleID:   applicationDTO.AppliedRoleID,
		Description:     applicationDTO.Description,
		StatusID:        businesslogic.RoleApplicationStatusPending,
		ApprovalUserID:  nil,
		CreateUserID:    currentUser.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUser.ID,
		DateTimeUpdated: time.Now(),
	}

	createErr := server.service.CreateRoleApplication(currentUser, &application)
	if createErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, createErr.Error(), nil)
		return
	} else {
		util.RespondJsonResult(w, http.StatusOK, "role application has been submitted successfully", nil)
		return
	}
}

func (server RoleApplicationServer) SearchRoleApplicationHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, userErr := server.auth.GetCurrentUser(r)
	if userErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
	}

	criteria := new(businesslogic.SearchRoleApplicationCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad search criteria, please try again", nil)
		return
	}

	applications, searchErr := server.service.SearchRoleApplication(currentUser, *criteria)
	if searchErr != nil {
		log.Println(searchErr.Error())
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot search role application", nil)
	}

	dtos := make([]viewmodel.RoleApplication, 0)
	for _, each := range applications {
		dtos = append(dtos, viewmodel.RoleApplication{
			ID:                each.ID,
			RoleApplied:       each.AppliedRoleID,
			Description:       each.Description,
			Status:            each.StatusID,
			DateTimeSubmitted: each.DateTimeCreated,
			DateTimeResponded: each.DateTimeApproved,
		})
	}

	output, _ := json.Marshal(dtos)
	w.Write(output)
}

func (server RoleApplicationServer) ProvisionRoleApplicationHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, userErr := server.auth.GetCurrentUser(r)
	if userErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	responseDTO := new(viewmodel.RespondRoleApplication)
	parseErr := util.ParseRequestBodyData(r, responseDTO)

	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad response data, please try again", nil)
		return
	}
	if responseDTO.Validate() != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad response data, please try again", nil)
		return
	}

	applications, searchErr := server.service.SearchRoleApplication(currentUser, businesslogic.SearchRoleApplicationCriteria{
		ID: responseDTO.ApplicationID,
	})
	if searchErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot search this application", nil)
		return
	}
	if len(applications) != 1 {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot find this application", nil)
		return
	}

	updateErr := server.service.UpdateApplication(currentUser, &applications[0], responseDTO.Response)
	if updateErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot process this action", nil)
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "done", nil)
}
