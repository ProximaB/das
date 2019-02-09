package middleware

import (
	"github.com/DancesportSoftware/das/auth/firebase"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/env"
)

var AuthenticationStrategy = firebase.NewFirebaseAuthenticationStrategy(env.FirebaseAuthCredential, database.AccountRepository)
