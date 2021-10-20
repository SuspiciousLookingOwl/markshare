package providers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/suspiciouslookingowl/markshare/server/config"
	"go.uber.org/fx"
)

// Constructor
type GitHubAuthProvider struct {
	ClientID     string
	ClientSecret string
}

type gitHubAuthProviderParam struct {
	fx.In

	*config.Env
}

func NewGitHubAuthProvider(p gitHubAuthProviderParam) *GitHubAuthProvider {
	return &GitHubAuthProvider{
		ClientID:     p.GitHubClientID,
		ClientSecret: p.GitHubClientSecret,
	}
}

// Get Access Token
type GitHubAccessToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (p *GitHubAuthProvider) GetAccessToken(credential string) (string, error) {
	requestUrl, _ := url.Parse("https://github.com/login/oauth/access_token")
	q := requestUrl.Query()
	q.Set("client_id", p.ClientID)
	q.Set("client_secret", p.ClientSecret)
	q.Set("code", credential)
	q.Set("redirect_url", "http://localhost:3000")
	requestUrl.RawQuery = q.Encode()

	req, _ := http.NewRequest("POST", requestUrl.String(), nil)
	req.Header.Add("Accept", "application/json")
	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		return "", errors.New("failed to get access token")
	}

	var accessToken GitHubAccessToken
	json.NewDecoder(res.Body).Decode(&accessToken)
	if accessToken.AccessToken == "" {
		return "", errors.New("failed to get access token")
	}

	return accessToken.AccessToken, nil
}

// Get User Emails
type GitHubUserEmail struct {
	Email      string  `json:"email"`
	Primary    bool    `json:"primary"`
	Verified   bool    `json:"verified"`
	Visibility *string `json:"visibility"`
}

func (p *GitHubAuthProvider) GetUserPrimaryEmail(accessToken string) (string, error) {
	req, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	req.Header.Add("Authorization", "token "+accessToken)
	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		return "", errors.New("failed to identify user")
	}

	var emails []GitHubUserEmail
	json.NewDecoder(res.Body).Decode(&emails)

	// get user id
	primaryEmail := emails[0]
	if !primaryEmail.Primary || !primaryEmail.Verified {
		return "", errors.New("email not found")
	}

	return primaryEmail.Email, nil
}

// Get User
type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func (p *GitHubAuthProvider) GetUser(accessToken string) (*GitHubUser, error) {
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Add("Authorization", "token "+accessToken)
	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		return nil, errors.New("failed to identify user")
	}

	var user GitHubUser
	json.NewDecoder(res.Body).Decode(&user)

	return &user, nil
}
