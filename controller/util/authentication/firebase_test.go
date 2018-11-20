package authentication_test

import (
	"context"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestNewFirebaseAuthenticationStrategy(t *testing.T) {
	log.Println(os.Getwd())
	strategy := authentication.NewFirebaseAuthenticationStrategy("../../../secret/firebasePrivateKey.json", nil)
	user := strategy.CreateAccount(context.Background(), businesslogic.Account{
		Email: "test@test.com",
	}, "abcdefgH3df")
	assert.Equal(t, "test@test.com", user.Email)
}
