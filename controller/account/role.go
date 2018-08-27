package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"time"
)

type RoleApplicationServer struct {
	auth            authentication.IAuthenticationStrategy
	accountRepo     businesslogic.IAccountRepository
	roleAppRepo     businesslogic.IRoleApplicationRepository
	accountRoleRepo businesslogic.IAccountRoleRepository
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
	parseErr := util.ParseRequestData(r, applicationDTO)

	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad application data, please try again", nil)
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

	createErr := server.roleAppRepo.CreateApplication(&application)
	if createErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, createErr.Error(), nil)
		return
	} else {
		util.RespondJsonResult(w, http.StatusOK, "role application has been submitted successfully", nil)
		return
	}
}

func (server RoleApplicationServer) SearchRoleApplicationHandler(w http.ResponseWriter, r *http.Request) {

}
