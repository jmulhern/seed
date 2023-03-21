package pkg

type Site struct {
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Favicon     string `json:"favicon,omitempty" yaml:"favicon,omitempty"`
	BodyClasses string `json:"body_classes,omitempty" yaml:"body-classes,omitempty"`
	Email       Email  `json:"email,omitempty" yaml:"email,omitempty"`
}
