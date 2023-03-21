package pkg

type Seed struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	HostedZone string `json:"hosted_zone,omitempty" yaml:"hosted-zone,omitempty"`
	FQDN       string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	Bucket     Bucket `json:"bucket,omitempty" yaml:"bucket,omitempty"`
	SMTP       SMTP   `json:"smtp,omitempty" yaml:"smtp,omitempty"`
	Email      Email  `json:"email,omitempty" yaml:"email,omitempty"`
	Site       Site   `json:"site,omitempty" yaml:"site,omitempty"`
}
