package client

import "fmt"

type Spec struct {
	// plugin spec goes here
	CustomerID     string              `json:"customer_id,omitempty"`
	OAuth          *oauthSpec          `json:"oauth,omitempty"`
	ServiceAccount *serviceAccountSpec `json:"service_account,omitempty"`
}

func (s *Spec) validate() error {
	if len(s.CustomerID) == 0 {
		return fmt.Errorf(`required field "customer_id" is missing`)
	}

	if s.OAuth != nil && s.ServiceAccount != nil {
		return fmt.Errorf(`only one of "oauth" or "service_account" can be specified`)
	}

	if s.OAuth != nil {
		if err := s.OAuth.validate(); err != nil {
			return fmt.Errorf("invalid oauth config: %w", err)
		}
	}

	if s.ServiceAccount != nil {
		if err := s.ServiceAccount.validate(); err != nil {
			return fmt.Errorf("invalid service_account config: %w", err)
		}
	}

	return nil
}
