// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"golang.org/x/net/context"
	"net/http"

	"google.golang.org/api/option"
	"log"
)

type FirebaseAuthenticationStrategy struct {
	businesslogic.IAccountRepository
	app    *firebase.App
	client *auth.Client
}

func NewFirebaseAuthenticationStrategy(firebaseKey string, accountRepo businesslogic.IAccountRepository) FirebaseAuthenticationStrategy {
	opt := option.WithCredentialsFile(firebaseKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase authentication: %v", err)
	}
	client, err := app.Auth(context.Background())
	return FirebaseAuthenticationStrategy{
		accountRepo,
		app,
		client,
	}
}

func (strategy FirebaseAuthenticationStrategy) GetCurrentUser(r *http.Request) businesslogic.Account {
	ctx := r.Context()
	uid := "some_string_uid"
	client, err := strategy.app.Auth(ctx)
	if err != nil {
		log.Printf("[firebase-auth] error authenticating user: %v", err)
	}

	user, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Printf("[firebase-auth] error getting user: %v", err)
	}

	return businesslogic.Account{
		UUID:  user.UID,
		Email: user.Email,
		Phone: user.PhoneNumber,
	}
}

func (strategy FirebaseAuthenticationStrategy) CreateAccount(ctx context.Context, account businesslogic.Account, password string) *auth.UserRecord {
	params := (&auth.UserToCreate{}).Email(account.Email).EmailVerified(false).Password(password).DisplayName(account.FullName()).Disabled(false)
	u, err := strategy.client.CreateUser(ctx, params)
	if err != nil {
		log.Printf("[firebase-auth] error creating user: %v", err)
	}
	return u
}
