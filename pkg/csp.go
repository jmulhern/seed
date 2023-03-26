package pkg

import (
	"fmt"
	"strings"
)

type ContentSecurityPolicy struct {
	Enabled       bool     `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ReportOnly    bool     `json:"report_only,omitempty" yaml:"report-only,omitempty"`
	Default       []string `json:"default_src,omitempty" yaml:"default-src,omitempty"`
	Script        []string `json:"script_src,omitempty" yaml:"script-src,omitempty"`
	ScriptElement []string `json:"script_src_elem,omitempty" yaml:"script-src-elem,omitempty"`
	Style         []string `json:"style_src,omitempty" yaml:"style-src,omitempty"`
	Image         []string `json:"image_src,omitempty" yaml:"image-src,omitempty"`
	Object        []string `json:"object_src,omitempty" yaml:"object-src,omitempty"`
	BaseURI       []string `json:"base_uri,omitempty" yaml:"base-uri,omitempty"`
	ReportURI     string   `json:"report_uri,omitempty" yaml:"report-uri,omitempty"`
}

func (csp ContentSecurityPolicy) Make(nonce string) string {
	if !csp.Enabled {
		return ""
	}
	var sb strings.Builder
	if len(csp.Default) > 0 {
		sb.WriteString("default-src ")
		policyParams := params(csp.Default)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.Script) > 0 {
		sb.WriteString("script-src ")
		policyParams := strings.ReplaceAll(params(csp.Script), "<nonce>", nonce)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.ScriptElement) > 0 {
		sb.WriteString("script-src-elem ")
		policyParams := strings.ReplaceAll(params(csp.ScriptElement), "<nonce>", nonce)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.Style) > 0 {
		sb.WriteString("style-src ")
		policyParams := params(csp.Style)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.Image) > 0 {
		sb.WriteString("img-src ")
		policyParams := params(csp.Image)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.Object) > 0 {
		sb.WriteString("object-src ")
		policyParams := params(csp.Object)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.BaseURI) > 0 {
		sb.WriteString("base-uri ")
		policyParams := params(csp.BaseURI)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if len(csp.BaseURI) > 0 {
		sb.WriteString("base-uri ")
		policyParams := params(csp.BaseURI)
		sb.WriteString(policyParams)
		sb.WriteString(" ; ")
	}
	if csp.ReportURI != "" {
		sb.WriteString("report-uri ")
		sb.WriteString(csp.ReportURI)
		sb.WriteString(" ; ")
	}
	return sb.String()
}

func params(values []string) string {
	for i, v := range values {
		if !strings.HasPrefix(v, "'") {
			values[i] = fmt.Sprintf("'%s'", v)
		}
	}
	return strings.Join(values, " ")
}
