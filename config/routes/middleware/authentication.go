package middleware

import (
	"github.com/ProximaB/das/auth/firebase"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/env"
)

var AuthenticationStrategy = firebase.NewFirebaseAuthenticationStrategy(env.FirebaseAuthCredential, database.AccountRepository)
