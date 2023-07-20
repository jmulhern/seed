package heritage

type Stripe struct {
	Enabled        bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	PublishableKey string `json:"publishable_key,omitempty" yaml:"publishable-key,omitempty"`
	SecretKey      string `json:"secret_key,omitempty" yaml:"secret-key,omitempty"`
	BasicPriceID   string `json:"basic_price_id,omitempty" yaml:"basic-price-id,omitempty"`
	ProPriceID     string `json:"pro_price_id:,omitempty" yaml:"pro-price-id:,omitempty"`
}
