package model

type OauthTokenRequest struct {
	GrantType    string
	Code         string
	RedirectURI  string
	ClientId     string
	ClientSecret string
	CodeVerifier string
	Username     string
	Password     string
	RefreshToken string
	Scope        []string
}
