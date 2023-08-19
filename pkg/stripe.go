package heritage

type Stripe struct {
	Enabled    bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	SecretKey  string `json:"secret_key,omitempty" yaml:"secret-key,omitempty"`
	SuccessURL string `json:"success_url,omitempty" yaml:"success-url,omitempty"`
	CancelURL  string `json:"cancel_url,omitempty" yaml:"cancel-url,omitempty"`
}
