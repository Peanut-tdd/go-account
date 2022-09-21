package vo

type GetAccessToken struct {
	Result      int    `json:"result"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
