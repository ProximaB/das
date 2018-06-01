package util

import (
	"github.com/yubing24/das/businesslogic"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Identity struct {
	Username    string
	Email       string
	AccountType int
	AccountID   string
}
type AuthenticationStrategy interface {
	Authenticate(r *http.Request) (*businesslogic.Account, error)
}

type DatabaseAuthenticationStrategy struct {
	businesslogic.IAccountRepository
}

func (strategy DatabaseAuthenticationStrategy) Authenticate(r *http.Request) (*businesslogic.Account, error) {
	token, tokenErr := getAuthenticatedRequestToken(r)
	if tokenErr != nil {
		return nil, tokenErr
	}
	identity := getAuthenticatedRequestIdentity(token)
	account := businesslogic.GetAccountByUUID(identity.AccountID, strategy.IAccountRepository)
	if account.ID == 0 {
		return nil, errors.New(fmt.Sprintf("account with identity %+v is not found", identity))
	}
	return &account, nil
}

var HMAC_SECRET = ""

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

const (
	JWT_AUTH_CLAIM_EMAIL    = "email"
	JWT_AUTH_CLAIM_USERNAME = "name"
	JWT_AUTH_CLAIM_TYPE     = "type"
	JWT_AUTH_CLAIM_UUID     = "uuid"
)

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

func GenerateAuthenticationToken(account businesslogic.Account) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		JWT_AUTH_CLAIM_EMAIL:    account.Email,
		JWT_AUTH_CLAIM_TYPE:     strconv.Itoa(account.AccountTypeID),
		JWT_AUTH_CLAIM_USERNAME: account.FirstName + " " + account.LastName,
		JWT_AUTH_CLAIM_UUID:     account.UUID,
	})
	authString, err := token.SignedString([]byte(HMAC_SECRET))
	if err != nil {
		log.Panicf("failed to generate authentication token for legit user: %s\n", err)
	}
	return authString
}
