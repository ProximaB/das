package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"time"
)

type RoleApplicationServer struct {
	auth    auth.IAuthenticationStrategy
	service businesslogic.RoleProvisionService
}

func NewRoleApplicationServer(authStrat auth.IAuthenticationStrategy, service businesslogic.RoleProvisionService) RoleApplicationServer {
	return RoleApplicationServer{
		auth:    authStrat,
		service: service,
	}
}

// CreateRoleApplicationHandler handles the request:
//	POST /api/v1.0/account/role/application
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

	if errs := validator.Validate(applicationDTO); errs != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, errs.Error(), nil)
		return
	}

	application := businesslogic.RoleApplication{
		AccountID:       currentUser.ID,
		AppliedRoleID:   applicationDTO.RoleID,
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
	}
	util.RespondJsonResult(w, http.StatusOK, "role application has been submitted successfully", nil)
	return
}

// SearchRoleApplicationHandler handles the request:
//	GET /api/v1.0/account/role/application
func (server RoleApplicationServer) SearchRoleApplicationHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, userErr := server.auth.GetCurrentUser(r)
	if userErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	criteria := new(businesslogic.SearchRoleApplicationCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad search criteria, please try again", nil)
		return
	}

	criteria.AccountID = currentUser.ID // only searches applications made by current user

	applications, searchErr := server.service.SearchRoleApplication(currentUser, *criteria)
	if searchErr != nil {
		log.Println(searchErr.Error())
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot search role application", nil)
		return
	}

	dtos := make([]viewmodel.RoleApplicationAdminView, 0)
	for _, each := range applications {
		dtos = append(dtos, viewmodel.RoleApplicationAdminView{
			ID:                each.ID,
			ApplicantName:     each.Account.FullName(),
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

// AdminGetRoleApplicationHandler handles the request:
//	GET /api/v1/admin/role/application
// This will return all role applications to the admin user. This handler is for admin uses only and should enforce role check
func (server RoleApplicationServer) AdminGetRoleApplicationHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, userErr := server.auth.GetCurrentUser(r)
	if userErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	if !currentUser.HasRole(businesslogic.AccountTypeAdministrator) {
		util.RespondJsonResult(w, http.StatusUnauthorized, "admin priviledge is required", nil)
		return
	}

	criteria := new(businesslogic.SearchRoleApplicationCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad search criteria, please double check", nil)
		return
	}
	// admin user can search requests without additional moderation of search criteria

	applications, searchErr := server.service.SearchRoleApplication(currentUser, *criteria)
	if searchErr != nil {
		log.Println(searchErr.Error())
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot search role application", nil)
		return
	}

	dtos := make([]viewmodel.RoleApplicationAdminView, 0)
	for _, each := range applications {
		dtos = append(dtos, viewmodel.RoleApplicationAdminView{
			ID:                each.ID,
			ApplicantName:     each.Account.FullName(),
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

// ProvisionRoleApplicationHandler handles the request:
//	PUT /api/v1.o/account/role/provision
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
	if validator.Validate(responseDTO) != nil {
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
	if len(applications) == 0 {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot find this application", nil)
		return
	}

	updateErr := server.service.UpdateApplication(currentUser, &applications[0], responseDTO.Response)
	if updateErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, updateErr.Error(), nil)
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "done", nil)
}

// RoleServer handles the role information of user
type RoleServer struct {
	Auth        auth.IAuthenticationStrategy
	AccountRepo businesslogic.IAccountRepository
}

// GetAccountRoles get the roles of current user
func (server RoleServer) GetAccountRolesHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, userErr := server.Auth.GetCurrentUser(r)
	if userErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		util.RespondJsonResult(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	roles := currentUser.GetAccountRoles()
	data := make([]viewmodel.AccountRoleDTO, 0)
	for _, each := range roles {
		data = append(data, viewmodel.AccountRoleToAccountRoleDTO(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
