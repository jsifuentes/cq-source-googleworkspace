package client

import (
	"testing"
)

func TestSpecValidate(t *testing.T) {
	tests := []struct {
		name    string
		spec    Spec
		wantErr bool
		errMsg  string
	}{
		{
			name:    "missing CustomerID",
			spec:    Spec{},
			wantErr: true,
			errMsg:  `required field "customer_id" is missing`,
		},
		{
			name: "both OAuth and ServiceAccount specified",
			spec: Spec{
				CustomerID:     "customer_id",
				OAuth:          &oauthSpec{},
				ServiceAccount: &serviceAccountSpec{},
			},
			wantErr: true,
			errMsg:  `only one of "oauth" or "service_account" can be specified`,
		},
		// OAuth
		{
			name: "valid OAuth specified",
			spec: Spec{
				CustomerID: "customer_id",
				OAuth: &oauthSpec{
					ClientID:     "client_id",
					ClientSecret: "client_secret",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid OAuth specified: missing ClientID",
			spec: Spec{
				CustomerID: "customer_id",
				OAuth: &oauthSpec{
					ClientID: "",
				},
			},
			wantErr: true,
			errMsg:  "invalid oauth config: empty client_id in oauth spec",
		},
		{
			name: "invalid OAuth specified: missing ClientSecret",
			spec: Spec{
				CustomerID: "customer_id",
				OAuth: &oauthSpec{
					ClientID:     "client_id",
					ClientSecret: "",
				},
			},
			wantErr: true,
			errMsg:  "invalid oauth config: empty client_secret in oauth spec",
		},
		// Service Account
		{
			name: "valid ServiceAccount specified",
			spec: Spec{
				CustomerID: "customer_id",
				ServiceAccount: &serviceAccountSpec{
					JSONString:       "json_string",
					ImpersonateEmail: "impersonate_email",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid ServiceAccount: missing JSONString",
			spec: Spec{
				CustomerID: "customer_id",
				ServiceAccount: &serviceAccountSpec{
					JSONString: "",
				},
			},
			wantErr: true,
			errMsg:  `invalid service_account config: required field "json_string" is missing. It should be equal to the service account JSON string`,
		},
		{
			name: "invalid ServiceAccount: missing ImpersonateEmail",
			spec: Spec{
				CustomerID: "customer_id",
				ServiceAccount: &serviceAccountSpec{
					JSONString:       "json_string",
					ImpersonateEmail: "",
				},
			},
			wantErr: true,
			errMsg:  `invalid service_account config: required field "impersonate_email" is missing. It should be equal to the email address of the user to impersonate`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.spec.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Spec.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("Spec.validate() error = %v, wantErrMsg %v", err, tt.errMsg)
			}
		})
	}
}
