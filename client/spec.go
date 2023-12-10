package client

import "fmt"

type Spec struct {
	// plugin spec goes here
	CustomerID string     `json:"customer_id,omitempty"`
	OAuth      *oauthSpec `json:"oauth,omitempty"`
}

func (s *Spec) validate() error {
	if len(s.CustomerID) == 0 {
		return fmt.Errorf(`required field "customer_id" is missing`)
	}
	return s.OAuth.validate()
}
