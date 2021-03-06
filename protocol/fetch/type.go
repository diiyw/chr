package fetch

import (
	"github.com/diiyw/cuto/protocol/network"
)

// Unique request identifier.
type RequestId string

// Stages of the request to handle. Request will intercept before the request is
	// sent. Response will intercept after the response is received (but before response
	// body is received.
type RequestStage string

// 
type RequestPattern  struct {

	// Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to "*".
	UrlPattern	string	`json:"urlPattern,omitempty"`

	// If set, only requests for matching resource types will be intercepted.
	ResourceType	network.ResourceType	`json:"resourceType,omitempty"`

	// Stage at wich to begin intercepting requests. Default is Request.
	RequestStage	RequestStage	`json:"requestStage,omitempty"`
}

// Response HTTP header entry
type HeaderEntry  struct {

	// 
	Name	string	`json:"name"`

	// 
	Value	string	`json:"value"`
}

// Authorization challenge for HTTP status code 401 or 407.
type AuthChallenge  struct {

	// Source of the authentication challenge.
	Source	string	`json:"source,omitempty"`

	// Origin of the challenger.
	Origin	string	`json:"origin"`

	// The authentication scheme used, such as basic or digest
	Scheme	string	`json:"scheme"`

	// The realm of the challenge. May be empty.
	Realm	string	`json:"realm"`
}

// Response to an AuthChallenge.
type AuthChallengeResponse  struct {

	// The decision on what to do in response to the authorization challenge.  Default means
	// deferring to the default behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	Response	string	`json:"response"`

	// The username to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Username	string	`json:"username,omitempty"`

	// The password to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	Password	string	`json:"password,omitempty"`
}
