package env

const (
	VarDatabaseConnectionString = "POSTGRES_CONNECTION"
	VarFirebaseAuthCredential   = "FIREBASE_AUTH_CREDENTIAL"
	VarFirebaseProjectId        = "FIREBASE_PROJECT_ID"
	VarHMACSigningKey           = "HMAC_SIGNING_KEY"
	VarHMACValidHours           = "HMAC_VALID_HOURS"
)

const (
	LogLevelInfo    = 1
	LogLevelWarning = 2
	LogLevelError   = 3
)

var (
	DatabaseConnectionString string
	FirebaseAuthCredential   string
	HmacSigningKey           string
	HmacValidHours           int
)
