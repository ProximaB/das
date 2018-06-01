package businesslogic

import (
	"errors"
	"log"
	"reflect"
)

func AuthenticateUser(email, password string, repo IAccountRepository) error {
	if len(email) < 1 || len(password) < 1 {
		return errors.New("username or password is missing")
	}

	accounts, err := repo.SearchAccount(&SearchAccountCriteria{Email: email})
	if err != nil {
		return err
	}

	if len(accounts) != 1 {
		return errors.New("invalid email address")
	}

	if accounts[0].AccountStatusID == ACCOUNT_STATUS_SUSPENDED {
		return errors.New("account is suspended")
	}

	if accounts[0].AccountStatusID == ACCOUNT_STATUS_LOCKED {
		return errors.New("account is locked")
	}

	expectedHash := GenerateHash(accounts[0].PasswordSalt, []byte(password))
	if reflect.DeepEqual(expectedHash, accounts[0].PasswordHash) {
		// TODO: UpdateAccountSecurity (email, "login", true)
		log.Printf("%s was authenticated", email)
		return nil // user is authenticated
	} else {
		log.Printf("%s failed being authenticated", email)
		// TODO: UpdateAccountSecurity (email, "login", false)
		return errors.New("incorrect username or password")
	}
}
