package routes

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/routes/internal/account"
	"github.com/yubing24/das/config/routes/internal/organizer"
	"github.com/yubing24/das/config/routes/internal/partnership"
	"github.com/yubing24/das/config/routes/internal/reference"
	"github.com/yubing24/das/controller/util"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/*
var restAPIRouter = []Route{
	// Accounts
	{"Authenticate", http.MethodPost, "/api/account/authenticate", authorizeSingleRole(AuthenticationHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Create DAS Account", http.MethodPost, "/api/account/register", authorizeSingleRole(registerAccountHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},

	// Competition
	{"Search competitions", http.MethodGet, "/api/public/competitions", authorizeSingleRole(setResponseHeader(publicSearchCompetitionHandler), businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Get competition status", http.MethodGet, "/api/competition/status", authorizeSingleRole(getCompetitionStatusHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Get unique federations of events at competition", http.MethodGet, "/api/competition/federation", authorizeSingleRole(getEventUniqueFederationsHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Get unique divisions of events at competition", http.MethodGet, "/api/competition/division", authorizeSingleRole(getEventUniqueDivisionsHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Get unique ages of events at competition", http.MethodGet, "/api/competition/age", authorizeSingleRole(getEventUniqueAgesHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Get unique proficiencies of events at competition", http.MethodGet, "/api/competition/proficiency", authorizeSingleRole(getEventUniqueProficienciesHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Get unique styles of events at competition", http.MethodGet, "/api/competition/style", authorizeSingleRole(getEventUniqueStylesHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},

	// Events
	{"Public view of events", http.MethodGet, "/api/event", authorizeSingleRole(getEventHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"Public view of competitive ballroom events", http.MethodGet, "/api/event/competitive/ballroom", authorizeSingleRole(getCompetitiveBallroomEventHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
	{"[Organizer] Create a competitive ballroom event", http.MethodPost, "/api/organizer/event", authorizeSingleRole(createEventHandler, businesslogic.ACCOUNT_TYPE_ORGANIZER)},

	// Entries
	{"add/drop competitive ballroom event entries", http.MethodPost, "/api/athlete/registration", authorizeSingleRole(createAthleteRegistrationHandler, businesslogic.ACCOUNT_TYPE_ATHLETE)},
	{"Get competitive ballroom entries for partnership", http.MethodGet, "/api/athlete/registration", authorizeSingleRole(getAthleteEventRegistrationHandler, businesslogic.ACCOUNT_TYPE_ATHLETE)},
	{"Get competitive ballroom entries for public view", http.MethodGet, "/api/public/entries", authorizeSingleRole(getCompetitiveBallroomEventEntryHandler, businesslogic.ACCOUNT_TYPE_NOAUTH)},
}*/

type Identity struct {
	Username    string
	Email       string
	AccountType int
	AccountID   string
}

type JwtToken struct {
	Token string `json:"token"`
}

const (
	JWT_AUTH_CLAIM_EMAIL    = "email"
	JWT_AUTH_CLAIM_USERNAME = "name"
	JWT_AUTH_CLAIM_TYPE     = "type"
	JWT_AUTH_CLAIM_UUID     = "uuid"
)

var HMAC_SECRET = ""

func setResponseHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cookie")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	}
}
func getAuthenticatedRequestIdentity(token *jwt.Token) Identity {
	var identity Identity
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identity.Username = claims[JWT_AUTH_CLAIM_USERNAME].(string)
		identity.Email = claims[JWT_AUTH_CLAIM_EMAIL].(string)
		identity.AccountID = claims[JWT_AUTH_CLAIM_UUID].(string)
		accountType, _ := claims[JWT_AUTH_CLAIM_TYPE].(string)
		identity.AccountType, _ = strconv.Atoi(accountType)
	}
	return identity
}
func getRequestUserRole(r *http.Request) (int, error) {
	authHeader := r.Header.Get("authorization")
	if authHeader != "" {
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return -1, errors.New("cannot authorize user")
				}
				return []byte(HMAC_SECRET), nil
			})
			if token.Valid && err == nil {
				context.Set(r, "decoded", token.Claims)
				identity := getAuthenticatedRequestIdentity(token)
				return identity.AccountType, nil
			}
		}
		return -1, errors.New("unauthorized")
	} else {
		return -1, errors.New("not authorized")
	}
}
func addDasController(router *mux.Router, handler util.DasController) {
	router.
		Methods(handler.Method, http.MethodOptions).
		Path(handler.Endpoint).
		Name(handler.Description).
		Handler(setResponseHeader(authorizeMultipleRoles(handler.Handler, handler.AllowedRoles)))
}
func authorizeMultipleRoles(h http.HandlerFunc, roles []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole, authErr := getRequestUserRole(r)
		allowNoAuth := false
		authorized := false
		for _, each := range roles {
			if each == businesslogic.ACCOUNT_TYPE_NOAUTH {
				allowNoAuth = true
			}
			if each == userRole {
				authorized = true
			}
		}

		if authErr != nil && !allowNoAuth {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		} else if allowNoAuth {
			h.ServeHTTP(w, r)
		} else if authorized {
			h.ServeHTTP(w, r)
		} else {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
	}
}

func addDasControllerGroup(router *mux.Router, group util.DasControllerGroup) {
	for _, each := range group.Controllers {
		addDasController(router, each)
	}
}

func DasRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Schemes("https")

	addDasControllerGroup(router, reference.CountryControllerGroup)
	addDasControllerGroup(router, reference.StateControllerGroup)
	addDasControllerGroup(router, reference.CityControllerGroup)
	addDasControllerGroup(router, reference.SchoolControllerGroup)
	addDasControllerGroup(router, reference.StudioControllerGroup)
	addDasControllerGroup(router, reference.FederationControllerGroup)
	addDasControllerGroup(router, reference.DivisionControllerGroup)
	addDasControllerGroup(router, reference.AgeControllerGroup)
	addDasControllerGroup(router, reference.ProficiencyControllerGroup)
	addDasControllerGroup(router, reference.StyleControllerGroup)
	addDasControllerGroup(router, reference.DanceControllerGroup)

	// account
	addDasControllerGroup(router, account.AccountControllerGroup)
	addDasController(router, account.AccountTypeController)
	addDasController(router, account.GenderController)

	// partnership request blacklist
	addDasController(router, partnership.GetPartnershipBlacklistReasonController)
	addDasControllerGroup(router, partnership.PartnershipRequestBlacklistControllerGroup)

	// partnership request
	addDasController(router, partnership.PartnershipRequestStatusController)
	addDasControllerGroup(router, partnership.PartnershipRequestControllerGroup)

	// partnership
	addDasControllerGroup(router, partnership.PartnershipControllerGroup)

	// organizer (multi-user group)
	addDasControllerGroup(router, organizer.ManageOrganizerProvisionControllerGroup)
	addDasControllerGroup(router, organizer.ProvisionControllerGroup)
	addDasControllerGroup(router, organizer.OrganizerProvisionControllerGroup)

	// organizer
	//addDasControllerGroup(router, organizer.OrganizerCompetitionManagementControllerGroup)

	// athlete

	// scrutineer

	// emcee

	// deck captain

	// adjudicator

	// administrator

	// public only

	log.Println("finishing controller initialization")
	return router
}
