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
	JWT_AUTH_CLAIM_TYPE       = "roles"
	JWT_AUTH_CLAIM_UUID       = "uuid"
	JWT_AUTH_CLAIM_ISSUEDON   = "issued" // issue time stamp (unix time)
	JWT_AUTH_CLAIM_EXPIRATION = "exp"    // expiration time stamp (unix time)
)

type JWTAuthenticationStrategy struct {
	businesslogic.IAccountRepository
}

type AuthorizedIdentity struct {
	Username    string
	Email       string
	AccountType int
	AccountID   string
}

func (strategy JWTAuthenticationStrategy) GetCurrentUser(r *http.Request) (businesslogic.Account, error) {
	token, tokenErr := getAuthenticatedRequestToken(r)
	if tokenErr != nil {
		log.Printf(" %v %v: authorization failed for request: %v\n", r.Method, r.RequestURI, tokenErr)
		return businesslogic.Account{}, tokenErr
	}
	identity := getAuthenticatedRequestIdentity(token)
	searchResults, searchErr := strategy.SearchAccount(businesslogic.SearchAccountCriteria{UUID: identity.AccountID})
	if searchErr != nil || len(searchResults) != 1 {
		log.Println(searchErr)
		return businesslogic.Account{}, errors.New("cannot be authorized")
	}
	account := searchResults[0]
	if account.ID == 0 {
		err := errors.New(fmt.Sprintf("account with identity %+v is not found", identity))
		log.Println(err)
		return businesslogic.Account{}, err
	}
	return account, nil
}

func (strategy JWTAuthenticationStrategy) SetAuthorizationResponse(w http.ResponseWriter) {
}

func GenerateAuthenticationToken(account businesslogic.Account) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		JWT_AUTH_CLAIM_EMAIL:      account.Email,
		JWT_AUTH_CLAIM_USERNAME:   account.FullName(),
		JWT_AUTH_CLAIM_UUID:       account.UUID,
		JWT_AUTH_CLAIM_ISSUEDON:   time.Now().Unix(),
		JWT_AUTH_CLAIM_EXPIRATION: time.Now().Add(time.Hour * time.Duration(HMAC_VALID_HOURS)).Unix(),
	})
	authString, err := token.SignedString([]byte(HMAC_SIGNING_KEY))
	if err != nil {
		log.Printf("failed to generate authentication token for legit user: %s\n", err)
	}
	return authString
}

func ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid authentication token")
		}
		return []byte(HMAC_SIGNING_KEY), nil
	})
	return err
}

func getAuthenticatedRequestIdentity(token *jwt.Token) AuthorizedIdentity {
	var identity AuthorizedIdentity
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
	authHeader := r.Header.Get("Authorization")
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

func (strategy JWTAuthenticationStrategy) Authenticate(r *http.Request) (*businesslogic.Account, error) {
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
