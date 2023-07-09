package heritage

type Email struct {
	From string   `json:"from,omitempty" yaml:"from,omitempty"`
	To   []string `json:"to,omitempty" yaml:"to,omitempty"`
}
