package heritage

type Cookie struct {
	Secure bool   `json:"secure,omitempty" yaml:"secure,omitempty"`
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
