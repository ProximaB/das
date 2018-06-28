// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package authentication

import (
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/util/authentication"
)

var AuthenticationStrategy = authentication.JwtAuthenticationStrategy{
	database.AccountRepository,
}
