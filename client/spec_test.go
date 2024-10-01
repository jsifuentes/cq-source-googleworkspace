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
		{
			name: "valid OAuth",
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
			name: "invalid OAuth - missing ClientID",
			spec: Spec{
				CustomerID: "customer_id",
				OAuth: &oauthSpec{
					ClientSecret: "client_secret",
				},
			},
			wantErr: true,
			errMsg:  "empty client_id in oauth spec",
		},
		{
			name: "invalid OAuth - missing ClientSecret",
			spec: Spec{
				CustomerID: "customer_id",
				OAuth: &oauthSpec{
					ClientID: "client_id",
				},
			},
			wantErr: true,
			errMsg:  "empty client_secret in oauth spec",
		},
		{
			name: "valid ServiceAccount",
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
			name: "invalid ServiceAccount - missing JSONString",
			spec: Spec{
				CustomerID: "customer_id",
				ServiceAccount: &serviceAccountSpec{
					ImpersonateEmail: "impersonate_email",
				},
			},
			wantErr: true,
			errMsg:  `required field "json_string" is missing. It should be equal to the service account JSON string`,
		},
		{
			name: "invalid ServiceAccount - missing ImpersonateEmail",
			spec: Spec{
				CustomerID: "customer_id",
				ServiceAccount: &serviceAccountSpec{
					JSONString: "json_string",
				},
			},
			wantErr: true,
			errMsg:  `required field "impersonate_email" is missing. It should be equal to the email address of the user to impersonate`,
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
