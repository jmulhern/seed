package pkg

type CDN struct {
	FQDN   string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
