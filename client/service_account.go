package client

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type serviceAccountSpec struct {
	JSONString       string `json:"json_string,omitempty"`
	ImpersonateEmail string `json:"impersonate_email,omitempty"`
}

func (s *serviceAccountSpec) validate() error {
	if len(s.JSONString) == 0 {
		return fmt.Errorf(`required field "json_string" is missing. It should be equal to the service account JSON string`)
	}
	if len(s.ImpersonateEmail) == 0 {
		return fmt.Errorf(`required field "impersonate_email" is missing. It should be equal to the email address of the user to impersonate`)
	}
	return nil
}

func (s *serviceAccountSpec) getTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	conf, err := google.JWTConfigFromJSON([]byte(s.JSONString))
	if err != nil {
		return nil, err
	}
	conf.Subject = s.ImpersonateEmail
	conf.Scopes = OAuthScopes
	return conf.TokenSource(ctx), nil
}
