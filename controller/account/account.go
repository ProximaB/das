package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
)

type Server struct {
	businesslogic.IAccountRepository
	businesslogic.IOrganizerProvisionRepository
	businesslogic.IOrganizerProvisionHistoryRepository
}

// RegisterAccountHandler handle the request
// 	POST /api/account/register
// Accepted JSON parameters:
//	{
//		"accounttype": 1,
//		"email": "awesomeuser@email.com",
//		"phone": 1234567890,
//		"firstname": "John",
//		"middlename": "Adams",
//		"lastname": "Smith",
//		"dateofbirth: "1990-01-01T06:00:00.0000Z",
//		"gender": 2,
//		"password": !@#$1234,
//		"tosaccepted": true,
//		"ppaccepted": true
//	}
// In case an user of under-13 registers an account, parental control will be required and
// additional parameters will be required as well:
// 	"byguardian": true
//	"signature": "John Smith Sr."
func (server Server) RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	createAccount := new(viewmodel.CreateAccount)

	if err := util.ParseRequestBodyData(r, createAccount); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	account := businesslogic.Account{
		AccountTypeID:         createAccount.AccountType,
		FirstName:             createAccount.FirstName,
		MiddleNames:           createAccount.MiddleNames,
		LastName:              createAccount.LastName,
		DateOfBirth:           createAccount.DateOfBirth,
		UserGenderID:          createAccount.Gender,
		Email:                 createAccount.Email,
		Phone:                 createAccount.Phone,
		ToSAccepted:           createAccount.ToSAccepted,
		PrivacyPolicyAccepted: createAccount.PPAccepted,
		ByGuardian:            createAccount.ByGuardian,
		Signature:             createAccount.Signature,
	}

	var strategy businesslogic.ICreateAccountStrategy
	switch account.AccountTypeID {
	case businesslogic.ACCOUNT_TYPE_ORGANIZER:
		strategy = businesslogic.CreateOrganizerAccountStrategy{
			AccountRepo:   server.IAccountRepository,
			ProvisionRepo: server.IOrganizerProvisionRepository,
			HistoryRepo:   server.IOrganizerProvisionHistoryRepository,
		}
	case businesslogic.ACCOUNT_TYPE_ADMINISTRATOR:
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid account type", nil)
		return
	default:
		strategy = businesslogic.CreateAccountStrategy{
			AccountRepo: server.IAccountRepository,
		}
	}
	if err := strategy.CreateAccount(account, createAccount.Password); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// AccountAuthenticationHandler handles the request:
// 	POST /api/account/authenticate
// Accepted JSON parameters:
// 	{
//		"username": "user@email.com",
//		"password": "password"
//	}
// Sample returned response:
//	{
//		"status": 200,
//		"message": "authorized",
//		"data": {
//			"token": "some.jwt.token",
//			"type: 3
//		}
//	}
func (server Server) AccountAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	loginDTO := new(viewmodel.Login)
	err := util.ParseRequestBodyData(r, loginDTO)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid credential", nil)
		return
	} else if len(loginDTO.Email) < 4 || len(loginDTO.Password) < 8 {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid credential", nil)
		return
	}

	err = businesslogic.AuthenticateUser(loginDTO.Email, loginDTO.Password, server.IAccountRepository)
	// util.RespondJsonResult(w, http.StatusNotImplemented, "authentication is not implemented", nil)

	if err != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, err.Error(), nil)
		return
	} else {
		account := businesslogic.GetAccountByEmail(loginDTO.Email, server.IAccountRepository)

		// user jwt authentication
		authString := authentication.GenerateAuthenticationToken(account)
		if err != nil {
			log.Printf("[error] generating client credential: %s\n", err.Error())
			util.RespondJsonResult(w, http.StatusUnauthorized, "error in generating client credential", nil)
			return
		} else {
			response := struct {
				Token    string `json:"token"`
				UserType int    `json:"type"`
			}{Token: authString, UserType: account.AccountTypeID}
			util.RespondJsonResult(w, http.StatusOK, "authorized", response)
			return
		}
	}
}

/*
func authorizeSingleRole(h http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if role == businesslogic.ACCOUNT_TYPE_NOAUTH {
			h.ServeHTTP(w, r)
			return
		}

		account, accountErr := GetCurrentUser(r, accountRepository)
		if accountErr != nil {
			log.Println(accountErr)
			util.RespondJsonResult(w, http.StatusInternalServerError, "cannot find user account due to internal error", nil)
			return
		}

		if account.AccountTypeID != role {
			util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized to perform this action", nil)
			return
		}
		h.ServeHTTP(w, r)
	}
}

// caution: this method assumes that request r has already been authenticated and no security check is performed here.
func getAuthenticatedRequestToken(r *http.Request) (*jwt.Token, error) {
	authHeader := r.Header.Get("authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) < 2 {
		return nil, errors.New("no authorization information")
	}
	token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot authorize user")
		}
		return []byte(HMAC_SECRET), nil
	})
	return token, nil
}

func GetCurrentUser(r *http.Request, repo businesslogic.IAccountRepository) (*businesslogic.Account, error) {
	token, tokenErr := getAuthenticatedRequestToken(r)
	if tokenErr != nil {
		return nil, tokenErr
	}
	identity := getAuthenticatedRequestIdentity(token)
	account := businesslogic.GetAccountByUUID(identity.AccountID, repo)
	if account.ID == 0 {
		return nil, errors.New(fmt.Sprintf("account with identity %+v is not found", identity))
	}
	return &account, nil
}
*/
