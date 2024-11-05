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
	"runtime"
	"strings"

	"github.com/rs/zerolog"
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

func (o *oauthSpec) getTokenSource(ctx context.Context, logger zerolog.Logger, endpoint oauth2.Endpoint) (oauth2.TokenSource, error) {
	lst, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
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
			return nil, fmt.Errorf("failed to get token from file: %v", err)
		}
	}

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("failed to generate state: %v", err)
	}
	state := strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")

	handler := &oauthHandler{
		state: state,
		err:   make(chan error),
	}

	srv := http.Server{Handler: handler}

	go func() {
		defer srv.Close()

		url := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

		var openErr error
		switch runtime.GOOS {
		case "windows":
			openErr = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			openErr = exec.Command("open", url).Start()
		case "linux":
			openErr = exec.Command("xdg-open", url).Start()
		}

		if openErr != nil {
			logger.Err(openErr).Msg("unable to open browser automatically to the authorization URL")
		}

		err = <-handler.err
	}()

	if serveErr := srv.Serve(lst); serveErr != http.ErrServerClosed {
		return nil, serveErr
	}

	if err != nil {
		return nil, fmt.Errorf("failed to complete OAuth authorization: %v", err)
	}

	// we have exchange token now
	token, err := config.Exchange(ctx, handler.code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %v", err)
	}

	if o.TokenFile != "" {
		if err := o.saveTokenToFile(token); err != nil {
			return nil, fmt.Errorf("failed to save token to file: %v", err)
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
