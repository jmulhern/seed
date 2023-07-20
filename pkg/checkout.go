package heritage

type Checkout struct {
	SuccessURL string `json:"success_url,omitempty" yaml:"success-url,omitempty"`
	CancelURL  string `json:"cancel_url,omitempty" yaml:"cancel-url,omitempty"`
}
