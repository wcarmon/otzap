package otzap

// -- Zap & OpenTelemetry Span Attribute Keys
const (
	defaultContextKey        = "ctx"
	defaultLevelKey          = "level"
	defaultLogEventSourceKey = "logEventSource"
	defaultSpanIdKey         = "spanId"
	defaultSpanKey           = "span"
	defaultTimestampKey      = "timestamp"
)

// -- Attribute Values
const (
	defaultOtelSourceValue = "otelApi"
	defaultZapSourceValue  = "zapApi"
)

// BlockedEnvVars lists keys which must NOT be logged
// not case sensitive
var BlockedEnvVars = []string{
	"AUTH_PASS",
	"AWS_ACCESS_KEY",
	"AWS_SECRET_ACCESS_KEY",
	"DB_PASS",
	"DB_PASSWORD",
	"GOOGLE_APPLICATION_CREDENTIALS",
	"JDBC_DATABASE_PASSWORD",
	"KAFKA_PASSWORD",
	"MARIADB_ROOT_PASSWORD",
	"MYSQL_PWD",
	"MYSQL_ROOT_PASSWORD",
	"PGPASSFILE",
	"PGPASSWORD",
	"PGSSLKEY",
	"PIN",
	"POSTGRES_PASSWORD",
	"PS1",
	"PS2",
	"SYSTEM_ACCESSTOKEN",
	"VISUAL",
}

// BlockedEnvVarSubstrings lists Key substrings which must NOT be logged
// not case sensitive
var BlockedEnvVarSubstrings = []string{
	"access_key",
	"accesstoken",
	"admin_token",
	"auth_pass",
	"authentication",
	"credentials",
	"database_password",
	"jwt",
	"kerberos",
	"keystore",
	"ldap",
	"openssl",
	"passwd",
	"password",
	"password_file",
	"root_password",
	"rsa",
	"secret",
	"sslkey",
	"token",
	"truststore",
}
