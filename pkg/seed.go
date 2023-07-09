package heritage

type Seed struct {
	Name                  string                `json:"name,omitempty" yaml:"name,omitempty"`
	HostedZone            string                `json:"hosted_zone,omitempty" yaml:"hosted-zone,omitempty"`
	FQDN                  string                `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	ContentSecurityPolicy ContentSecurityPolicy `json:"content_security_policy,omitempty" yaml:"content-security-policy,omitempty"`
	Bucket                Bucket                `json:"bucket,omitempty" yaml:"bucket,omitempty"`
	CDN                   CDN                   `json:"cdn,omitempty" yaml:"cdn,omitempty"`
	Cookie                Cookie                `json:"cookie,omitempty" yaml:"cookie,omitempty"`
	JWT                   JWT                   `json:"jwt,omitempty" yaml:"jwt,omitempty"`
	SMTP                  SMTP                  `json:"smtp,omitempty" yaml:"smtp,omitempty"`
	Email                 Email                 `json:"email,omitempty" yaml:"email,omitempty"`
	Login                 Login                 `json:"login,omitempty" yaml:"login,omitempty"`
	Site                  Site                  `json:"site,omitempty" yaml:"site,omitempty"`
}
