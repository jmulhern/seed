package pkg

type Packet struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	CIDR        string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	Unique      string `json:"unique,omitempty" yaml:"unique,omitempty"`
	Account     string `json:"account,omitempty" yaml:"account,omitempty"`
	Region      string `json:"region,omitempty" yaml:"region,omitempty"`
	KeySecretID string `json:"key_secret_id,omitempty" yaml:"key-secret-id,omitempty"`
	Seeds       []Seed `json:"seeds,omitempty" yaml:"seeds,omitempty"`
}
