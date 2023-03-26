package pkg

type CDN struct {
	FQDM   string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
