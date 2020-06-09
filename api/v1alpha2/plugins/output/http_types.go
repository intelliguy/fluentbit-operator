package output

import (
	"fmt"
	"kubesphere.io/fluentbit-operator/api/v1alpha2/plugins"
)

// +kubebuilder:object:generate:=true

// The http output plugin allows to flush your records into a HTTP endpoint.
type Http struct {
	// IP address or hostname of the target HTTP Server
	Host string `json:"host,omitempty"`
	// Basic Auth Username
	HttpUser string `json:"hostUser,omitempty"`
	// Basic Auth Password. Requires HTTP_User to be set
	HttpPasswd string `json:"hostPasswd,omitempty"`
	// TCP Port of the target service.
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:validation:Maximum:=65535
	Port *int32 `json:"port,omitempty"`
	// Specify an optional HTTP URI for the target web server, e.g: /something
	URI string `json:"uri,omitempty"`
	// Specify the data format to be used in the HTTP request body, by default it uses msgpack. Other supported formats are json, json_stream and json_lines and gelf.
	// +kubebuilder:validation:Enum:=msgpack;json;json_stream;json_lines;gelf
	Format string `json:"format,omitempty"`
}

func (_ *Http) Name() string {
	return "http"
}

// implement Section() method
func (f *Http) Params(sl plugins.SecretLoader) (*plugins.KVs, error) {
	kvs := plugins.NewKVs()
	if f.Host != "" {
		kvs.Insert("Host", f.Host)
	}
	if f.HttpUser != "" {
		kvs.Insert("HTTP_User", f.HttpUser)
	}
	if f.HttpPasswd != "" {
		kvs.Insert("HTTP_Passwd", f.HttpPasswd)
	}
	if f.Port != nil {
		kvs.Insert("Port", fmt.Sprint(f.Port))
	}
	if f.URI != "" {
		kvs.Insert("URI", f.URI)
	}
	if f.Format != "" {
		kvs.Insert("Format", f.Format)
	}
	return kvs, nil
}
