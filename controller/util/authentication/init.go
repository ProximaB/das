// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package authentication

import (
	"log"
	"os"
	"strconv"
)

var HMAC_SIGNING_KEY string
var HMAC_VALID_HOURS = 72

func init() {
	var durationErr error
	HMAC_SIGNING_KEY = os.Getenv("HMAC_SIGNING_KEY")
	HMAC_VALID_HOURS, durationErr = strconv.Atoi(os.Getenv("HMAC_VALID_HOURS"))

	if durationErr != nil {
		log.Println("[warning] HMAC_VALID_HOURS is not defined in this environment")
		log.Println("[info] Default duration will be used: 72 hours")
		HMAC_VALID_HOURS = 72
	}
	if len(HMAC_SIGNING_KEY) == 0 {
		// use a default key, not recommended!!!
		log.Println("[warning] HMAC_SIGNING_KEY is not defined in this environment")
		log.Println("[info] Default HMAC_SIGNING_KEY will be used")
		HMAC_SIGNING_KEY = "7ke7oi1+!q&11t!my0l)z$s-$p4j@fpt8+=ultj6=1zq8nsw$+"
	} else {
		log.Println("[info] HMAC_SIGNING_KEY is defined in this environment")
	}
}
