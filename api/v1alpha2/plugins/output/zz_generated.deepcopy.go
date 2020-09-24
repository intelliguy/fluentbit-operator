// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package output

import (
	"kubesphere.io/fluentbit-operator/api/v1alpha2/plugins"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elasticsearch) DeepCopyInto(out *Elasticsearch) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.HTTPUser != nil {
		in, out := &in.HTTPUser, &out.HTTPUser
		*out = new(plugins.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPPasswd != nil {
		in, out := &in.HTTPPasswd, &out.HTTPPasswd
		*out = new(plugins.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.LogstashFormat != nil {
		in, out := &in.LogstashFormat, &out.LogstashFormat
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTagKey != nil {
		in, out := &in.IncludeTagKey, &out.IncludeTagKey
		*out = new(bool)
		**out = **in
	}
	if in.GenerateID != nil {
		in, out := &in.GenerateID, &out.GenerateID
		*out = new(bool)
		**out = **in
	}
	if in.ReplaceDots != nil {
		in, out := &in.ReplaceDots, &out.ReplaceDots
		*out = new(bool)
		**out = **in
	}
	if in.TraceOutput != nil {
		in, out := &in.TraceOutput, &out.TraceOutput
		*out = new(bool)
		**out = **in
	}
	if in.TraceError != nil {
		in, out := &in.TraceError, &out.TraceError
		*out = new(bool)
		**out = **in
	}
	if in.CurrentTimeIndex != nil {
		in, out := &in.CurrentTimeIndex, &out.CurrentTimeIndex
		*out = new(bool)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(plugins.TLS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elasticsearch.
func (in *Elasticsearch) DeepCopy() *Elasticsearch {
	if in == nil {
		return nil
	}
	out := new(Elasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Forward) DeepCopyInto(out *Forward) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.TimeAsInteger != nil {
		in, out := &in.TimeAsInteger, &out.TimeAsInteger
		*out = new(bool)
		**out = **in
	}
	if in.SendOptions != nil {
		in, out := &in.SendOptions, &out.SendOptions
		*out = new(bool)
		**out = **in
	}
	if in.RequireAckResponse != nil {
		in, out := &in.RequireAckResponse, &out.RequireAckResponse
		*out = new(bool)
		**out = **in
	}
	if in.EmptySharedKey != nil {
		in, out := &in.EmptySharedKey, &out.EmptySharedKey
		*out = new(bool)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(plugins.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(plugins.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(plugins.TLS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Forward.
func (in *Forward) DeepCopy() *Forward {
	if in == nil {
		return nil
	}
	out := new(Forward)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Http) DeepCopyInto(out *Http) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Http.
func (in *Http) DeepCopy() *Http {
	if in == nil {
		return nil
	}
	out := new(Http)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kafka) DeepCopyInto(out *Kafka) {
	*out = *in
	if in.Rdkafka != nil {
		in, out := &in.Rdkafka, &out.Rdkafka
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kafka.
func (in *Kafka) DeepCopy() *Kafka {
	if in == nil {
		return nil
	}
	out := new(Kafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Null) DeepCopyInto(out *Null) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Null.
func (in *Null) DeepCopy() *Null {
	if in == nil {
		return nil
	}
	out := new(Null)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stdout) DeepCopyInto(out *Stdout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stdout.
func (in *Stdout) DeepCopy() *Stdout {
	if in == nil {
		return nil
	}
	out := new(Stdout)
	in.DeepCopyInto(out)
	return out
}
