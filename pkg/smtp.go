package heritage

type SMTP struct {
	HostedZone string `json:"hosted_zone,omitempty" yaml:"hosted-zone,omitempty"`
	Host       string `json:"host,omitempty" yaml:"host,omitempty"`
	Port       int    `json:"port,omitempty" yaml:"port,omitempty"`
	Username   string `json:"username,omitempty" yaml:"username,omitempty"`
	Password   string `json:"password,omitempty" yaml:"password,omitempty"`
}
