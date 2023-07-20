package heritage

type Stripe struct {
	Enabled        bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	PublishableKey string `json:"publishable_key,omitempty" yaml:"publishable-key,omitempty"`
	SecretKey      string `json:"secret_key,omitempty" yaml:"secret-key,omitempty"`
}
