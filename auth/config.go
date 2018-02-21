package auth

// What type of authentication and authorization will be used
type Method int

const (
	// No authentication or authorization
	AuthTypeNone Method = iota
	// "Plain" OpenID Connect and OAuth 2.0
	AuthTypeOIDC
	// HEART profiled OpenID Connect and OAuth 2.0
	AuthTypeHEART
)

// Config represents configuration information necessary to set up authentication
// and authorization for the FHIR server
type Config struct {
	Method           Method
	ClientID         string
	ClientSecret     string
	AuthorizationURL string
	TokenURL         string
	IntrospectionURL string
	UserInfoURL      string
	JWKPath          string
	OPURL            string
	SessionSecret    string
}

// None provides a server config where no authorization or authentication will
// be provided
func None() Config {
	return Config{Method: AuthTypeNone}
}
