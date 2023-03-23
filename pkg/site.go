package pkg

type Site struct {
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Favicon     string `json:"favicon,omitempty" yaml:"favicon,omitempty"`
	BodyClasses string `json:"body_classes,omitempty" yaml:"body-classes,omitempty"`
	Email       Email  `json:"email,omitempty" yaml:"email,omitempty"`
}
