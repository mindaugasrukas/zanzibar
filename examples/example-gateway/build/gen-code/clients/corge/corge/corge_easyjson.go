// Code generated by zanzibar
// @generated
// Checksum : b04N5OMJcHtH4hxo6sbAbA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package corge

import (
	json "encoding/json"
	fmt "fmt"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(in *jlexer.Lexer, out *Corge_EchoString_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(out *jwriter.Writer, in Corge_EchoString_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_EchoString_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_EchoString_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_EchoString_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_EchoString_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(in *jlexer.Lexer, out *Corge_EchoString_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "arg":
			out.Arg = string(in.String())
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(out *jwriter.Writer, in Corge_EchoString_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"arg\":"
		out.RawString(prefix[1:])
		out.String(string(in.Arg))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_EchoString_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_EchoString_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_EchoString_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_EchoString_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool(in *jlexer.Lexer, out *Corge_EchoBool_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(bool)
				}
				*out.Success = bool(in.Bool())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool(out *jwriter.Writer, in Corge_EchoBool_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_EchoBool_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_EchoBool_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_EchoBool_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_EchoBool_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool1(in *jlexer.Lexer, out *Corge_EchoBool_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "arg":
			out.Arg = bool(in.Bool())
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool1(out *jwriter.Writer, in Corge_EchoBool_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"arg\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Arg))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_EchoBool_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_EchoBool_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_EchoBool_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_EchoBool_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoBool1(l, v)
}
