package filter

import "kubesphere.io/fluentbit-operator/api/v1alpha2/plugins"

// +kubebuilder:object:generate:=true

// The Grep Filter plugin allows to match or exclude specific records based in regular expression patterns.
type Grep struct {
	// Keep records which field matches the regular expression.
	// Value Format: FIELD REGEX
	Regex string `json:"regex,omitempty"`
	// Exclude records which field matches the regular expression.
	// Value Format: FIELD REGEX
	Exclude string `json:"exclude,omitempty"`
}

func (_ *Grep) Name() string {
	return "grep"
}

func (g *Grep) Params(_ plugins.SecretLoader) (*plugins.KVs, error) {
	kvs := plugins.NewKVs()
	if g.Regex != "" {
		kvs.Insert("Regex", g.Regex)
	}
	if g.Exclude != "" {
		kvs.Insert("Exclude", g.Exclude)
	}
	return kvs, nil
}
