package client

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/oauth2"
	directory "google.golang.org/api/admin/directory/v1"
)

var OAuthScopes = []string{
	directory.AdminDirectoryCustomerReadonlyScope,
	directory.AdminDirectoryDomainReadonlyScope,
	directory.AdminDirectoryGroupMemberReadonlyScope,
	directory.AdminDirectoryGroupReadonlyScope,
	directory.AdminDirectoryOrgunitReadonlyScope,
	directory.AdminDirectoryUserAliasReadonlyScope,
	directory.AdminDirectoryUserReadonlyScope,
	directory.AdminDirectoryUserschemaReadonlyScope,
	directory.AdminDirectoryResourceCalendarReadonlyScope,
	directory.AdminDirectoryDeviceChromeosReadonlyScope,
	directory.AdminChromePrintersReadonlyScope,
}

type oauthSpec struct {
	TokenFile    string `json:"token_file,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func (o *oauthSpec) validate() error {
	switch {
	case len(o.TokenFile) > 0:
		return nil
	case len(o.ClientID) == 0:
		return fmt.Errorf("empty client_id in oauth spec")
	case len(o.ClientSecret) == 0:
		return fmt.Errorf("empty client_secret in oauth spec")
	default:
		return nil
	}
}

func (o *oauthSpec) getTokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(o.TokenFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func (o *oauthSpec) saveTokenToFile(token *oauth2.Token) error {
	f, err := os.OpenFile(o.TokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func (o *oauthSpec) getTokenSource(ctx context.Context, endpoint oauth2.Endpoint) (oauth2.TokenSource, error) {
	lst, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	config := &oauth2.Config{
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		Endpoint:     endpoint,
		RedirectURL:  "http://" + lst.Addr().String(),
		Scopes:       OAuthScopes,
	}

	if len(o.TokenFile) > 0 {
		savedToken, err := o.getTokenFromFile()
		if err == nil {
			return config.TokenSource(context.Background(), savedToken), nil
		} else if o.ClientID == "" || o.ClientSecret == "" {
			return nil, err
		}
	}

	b := make([]byte, 16)
	rand.Read(b)
	state := strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")

	handler := &oauthHandler{
		state: state,
		err:   make(chan error),
	}

	srv := http.Server{Handler: handler}

	go func() {
		defer srv.Close()
		_ = exec.CommandContext(ctx, "open", config.AuthCodeURL(state, oauth2.AccessTypeOffline)).Run()
		err = <-handler.err
	}()

	if serveErr := srv.Serve(lst); serveErr != http.ErrServerClosed {
		return nil, serveErr
	}

	if err != nil {
		return nil, err
	}

	// we have exchange token now
	token, err := config.Exchange(ctx, handler.code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, err
	}

	if o.TokenFile != "" {
		if err := o.saveTokenToFile(token); err != nil {
			return nil, err
		}
	}

	return config.TokenSource(context.Background(), token), nil
}

type oauthHandler struct {
	state string
	code  string
	err   chan error
}

func (o *oauthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer close(o.err)

	if state := r.FormValue("state"); state != o.state {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("incorrect \"state\" value: expected %q, got %q", o.state, state)
		fmt.Fprint(w, err.Error())
		o.err <- err
		return
	}

	o.code = r.FormValue("code")
	o.err <- nil

	fmt.Println("Code: " + o.code)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Authorization successful. You may close the window.")
}
