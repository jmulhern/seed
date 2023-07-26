package heritage

type Site struct {
	ID          string `json:"id" yaml:"id"`
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Favicon     string `json:"favicon,omitempty" yaml:"favicon,omitempty"`
	BodyClasses string `json:"body_classes,omitempty" yaml:"body-classes,omitempty"`
	HTMLClasses string `json:"html_classes,omitempty" yaml:"html-classes,omitempty"`
	Email       Email  `json:"email,omitempty" yaml:"email,omitempty"`
}
