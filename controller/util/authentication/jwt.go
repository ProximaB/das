// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package authentication

import (
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	JWT_AUTH_CLAIM_EMAIL      = "email"
	JWT_AUTH_CLAIM_USERNAME   = "name"
	JWT_AUTH_CLAIM_TYPE       = "type"
	JWT_AUTH_CLAIM_UUID       = "uuid"
	JWT_AUTH_CLAIM_ISSUEDON   = "issued"
	JWT_AUTH_CLAIM_EXPIRATION = "exp" // default expiration is 48 hours
)

type JwtAuthenticationStrategy struct {
	businesslogic.IAccountRepository
}
type Identity struct {
	Username    string
	Email       string
	AccountType int
	AccountID   string
}

func (strategy JwtAuthenticationStrategy) GetCurrentUser(r *http.Request, repo businesslogic.IAccountRepository) (businesslogic.Account, error) {
	token, tokenErr := getAuthenticatedRequestToken(r)
	if tokenErr != nil {
		return businesslogic.Account{}, tokenErr
	}
	identity := getAuthenticatedRequestIdentity(token)
	searchResults, searchErr := repo.SearchAccount(businesslogic.SearchAccountCriteria{UUID: identity.AccountID})
	if searchErr != nil || len(searchResults) != 1 {
		log.Println(searchErr)
		return businesslogic.Account{}, errors.New("cannot be authorized")
	}
	account := searchResults[0]
	if account.ID == 0 {
		return businesslogic.Account{}, errors.New(fmt.Sprintf("account with identity %+v is not found", identity))
	}
	return businesslogic.Account{}, nil
}

func (strategy JwtAuthenticationStrategy) SetAuthorizationResponse(w http.ResponseWriter) {
}

func GenerateAuthenticationToken(account businesslogic.Account) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		JWT_AUTH_CLAIM_EMAIL:      account.Email,
		JWT_AUTH_CLAIM_USERNAME:   account.FirstName + " " + account.LastName,
		JWT_AUTH_CLAIM_UUID:       account.UUID,
		JWT_AUTH_CLAIM_EXPIRATION: time.Now().Add(time.Hour * time.Duration(HMAC_VALID_HOURS)).Unix(),
	})
	authString, err := token.SignedString([]byte(HMAC_SIGNING_KEY))
	if err != nil {
		log.Panicf("failed to generate authentication token for legit user: %s\n", err)
	}
	return authString
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

// caution: this method assumes that request r has already been authenticated and no security check is performed here.
func getAuthenticatedRequestToken(r *http.Request) (*jwt.Token, error) {
	authHeader := r.Header.Get("authorization")
	if len(authHeader) < 1 {
		return nil, errors.New("empty authentication token")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) < 2 {
		return nil, errors.New("invalid authentication token")
	}

	token, tokenParseErr := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid authentication token")
		}
		return []byte(HMAC_SIGNING_KEY), nil
	})
	return token, tokenParseErr
}

func (strategy JwtAuthenticationStrategy) Authenticate(r *http.Request) (*businesslogic.Account, error) {
	// check if authentication token is valid
	token, tokenErr := getAuthenticatedRequestToken(r)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// token is good, check if account is valid
	identity := getAuthenticatedRequestIdentity(token)
	account := businesslogic.GetAccountByUUID(identity.AccountID, strategy.IAccountRepository)
	if account.ID == 0 {
		return nil, errors.New(fmt.Sprintf("account with identity %+v is not found", identity))
	}
	return &account, nil
}
