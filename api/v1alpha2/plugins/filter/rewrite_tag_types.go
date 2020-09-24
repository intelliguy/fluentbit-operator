package filter

import (
	"kubesphere.io/fluentbit-operator/api/v1alpha2/plugins"
)

// +kubebuilder:object:generate:=true

// The rewrite_tag filter, allows to re-emit a record under a new Tag. Once a record has been re-emitted, the original record can be preserved or discarded.
type RewriteTag struct {
	// Defines the matching criteria and the format of the Tag for the matching record.
	Rules []RewriteTagRule `json:"rules,omitempty"`
	// When the filter emits a record under the new Tag, there is an internal emitter plugin that takes care of the job.
	EmitterName string `json:"emitterName,omitempty"`
	// Define a buffering mechanism for the new records created. Note these records are part of the emitter plugin.
	// This option support the values memory (default) or filesystem.
	// +kubebuilder:validation:Enum:=memory;filesystem
	EmitterStorageType string `json:"emitterStorageType,omitempty"`
}

// +kubebuilder:object:generate:=true

// The plugin supports the following rules
type RewriteTagRule struct {
	// The key represents the name of the record key that holds the value that we want to use to match our regular expression.
	Key string `json:"key,omitempty"`
	// Using a simple regular expression we can specify a matching pattern to use against the value of the key specified above, also we can take advantage of group capturing to create custom placeholder values.
	Regex string `json:"regex,omitempty"`
	// If a regular expression has matched the value of the defined key in the rule, we are ready to compose a new Tag for that specific record.
	NewTag string `json:"newTag,omitempty"`
	// If a rule matches the criteria the filter will emit a copy of the record with the new defined Tag.
	Keep bool `json:"keep,omitempty"`
}

func (_ *RewriteTag) Name() string {
	return "rewrite_tag"
}

func (n *RewriteTag) Params(_ plugins.SecretLoader) (*plugins.KVs, error) {
	kvs := plugins.NewKVs()
	for _, rule := range n.Rules {
		kvs.Insert("Rule", rule.Key+" "+rule.Regex+" "+rule.NewTag+" "+FormatBool(rule.Keep))
	}
	if n.EmitterName != "" {
		kvs.Insert("Emitter_Name", n.EmitterName)
	}
	if n.EmitterStorageType != "" {
		kvs.Insert("Emitter_Storage.type", n.EmitterStorageType)
	}
	return kvs, nil
}

func FormatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
