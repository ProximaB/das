package firebase

import (
	"context"
	"errors"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
	"time"
)

// FirebaseAuthenticationStrategy implements IAuthenticationStrategy to use Firebase as the identity provider for DAS.
// Firebase's JWT uses RS256
type FirebaseAuthenticationStrategy struct {
	accountRepository businesslogic.IAccountRepository
	app               *firebase.App
	client            *auth.Client
	context           context.Context
}

func (strategy FirebaseAuthenticationStrategy) convertFirebaseUserToDasUser(user auth.UserRecord) businesslogic.Account {
	searchAccounts, searchErr := strategy.accountRepository.SearchAccount(businesslogic.SearchAccountCriteria{
		Email: user.Email,
	})
	if searchErr != nil || len(searchAccounts) != 1 || searchAccounts[0].Email != user.Email {
		var firstName, lastName string
		if len(strings.Split(user.DisplayName, " ")) > 1 {
			firstName = strings.Split(user.DisplayName, " ")[0]
			lastName = strings.Split(user.DisplayName, " ")[1]
		} else {
			firstName = "Unknown"
			lastName = "Unknown"
		}
		return businesslogic.Account{
			AccountStatusID: businesslogic.AccountStatusActivated,
			UserGenderID:    businesslogic.GENDER_UNKNOWN,
			FirstName:       firstName,
			LastName:        lastName,
			Email:           user.Email,
			Phone:           user.PhoneNumber,
			UID:             user.UID,
			DateTimeCreated: time.Unix(user.UserMetadata.CreationTimestamp, 0),
		}
	}
	return searchAccounts[0]
}

// NewFirebaseAuthenticationStrategy takes the credential (service account key) file and a handler to DAS account repository
// and instantiate an IAuthenticationStrategy that serves as the identity provider of DAS
func NewFirebaseAuthenticationStrategy(credential string, accountRepo businesslogic.IAccountRepository) FirebaseAuthenticationStrategy {
	opt := option.WithCredentialsJSON([]byte(credential))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase authentication: %v", err)
	}
	client, err := app.Auth(ctx)
	return FirebaseAuthenticationStrategy{
		accountRepository: accountRepo,
		app:               app,
		client:            client,
		context:           ctx,
	}
}

func (strategy FirebaseAuthenticationStrategy) GetUserByUID(uid string) (businesslogic.Account, error) {
	user, err := strategy.client.GetUser(strategy.context, uid)
	if err != nil {
		log.Printf("[error] cannot find user with UID %v in Firebase Auth", uid)
	}
	return strategy.convertFirebaseUserToDasUser(*user), nil
}

func (strategy FirebaseAuthenticationStrategy) GetUserByEmail(email string) (businesslogic.Account, error) {
	if user, err := strategy.client.GetUserByEmail(strategy.context, email); err != nil {
		return businesslogic.Account{}, err
	} else {
		return strategy.convertFirebaseUserToDasUser(*user), nil
	}
}

// GetCurrentUser attempts to get the user of from the HTTP request. The HTTP request should have an authorization token
// in its header. If the header is not found, then an empty user and an error will be returned to the caller function
func (strategy FirebaseAuthenticationStrategy) GetCurrentUser(r *http.Request) (businesslogic.Account, error) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < 1 {
		log.Printf("request %v misses authorization header", r.URL)
		return businesslogic.Account{}, errors.New("empty authentication token")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		log.Printf("request %v has invalid authorization header: %v", r.URL, authHeader)
		return businesslogic.Account{}, errors.New("invalid authentication token")
	}

	authToken := bearerToken[1]

	token, err := strategy.client.VerifyIDToken(strategy.context, authToken)
	if err != nil {
		log.Printf("[firebase-auth] error getting user: %v", err)
		return businesslogic.Account{}, err
	}
	if token.Expires < time.Now().Unix() {
		return businesslogic.Account{}, errors.New("token is expired")
	}
	user, findUserErr := strategy.GetUserByUID(token.UID)
	if findUserErr != nil {
		log.Printf("[error] cannot find user with UID %v: %v", token.UID, findUserErr)
		return businesslogic.Account{}, findUserErr
	}
	if user.Email == "" {
		return businesslogic.Account{}, errors.New("Invalid user")
	}
	return user, nil
}

// CreateUser assume that UID is provided, which will be used to retrieve account from Firebase. This will create a user
// in DAS instead of Firebase
func (strategy FirebaseAuthenticationStrategy) CreateUser(account *businesslogic.Account) error {
	var err error
	firstName := account.FirstName
	lastName := account.LastName
	phone := account.Phone
	*account, err = strategy.GetUserByUID(account.UID)
	if err != nil {
		log.Printf("[error] cannot get user with UID %v: %v", account.UID, err)
		return errors.New("Failed to create user in DAS")
	}
	account.FirstName = firstName
	account.LastName = lastName
	account.Phone = phone
	return strategy.accountRepository.CreateAccount(account)
}
