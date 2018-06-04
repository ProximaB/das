package authentication

import (
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/util/authentication"
)

var AuthenticationStrategy = authentication.JwtAuthenticationStrategy{
	database.AccountRepository,
}
