package heritage

type Login struct {
	Enabled      bool   `json:"enabled,omitempty" yaml:"enabled,omitempty,omitempty"`
	ClientID     string `json:"oidc_client_id,omitempty" yaml:"oidc-client-id,omitempty"`
	ClientSecret string `json:"oidc_client_secret,omitempty" yaml:"oidc-client-secret,omitempty"`
	RedirectURL  string `json:"oidc_redirect_url,omitempty" yaml:"oidc-redirect-url,omitempty"`
	ProviderURL  string `json:"provider_url,omitempty" yaml:"provider-url,omitempty"`
}
