// Code generated by zanzibar
// @generated
// Checksum : Qb0icL6muCAtJBX/hFlReQ==
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

func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge(in *jlexer.Lexer, out *NotModified) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge(out *jwriter.Writer, in NotModified) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NotModified) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NotModified) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NotModified) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NotModified) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge1(in *jlexer.Lexer, out *Foo) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge1(out *jwriter.Writer, in Foo) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Foo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Foo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Foo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Foo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorge1(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent(in *jlexer.Lexer, out *Corge_NoContent_Result) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "notModified":
			if in.IsNull() {
				in.Skip()
				out.NotModified = nil
			} else {
				if out.NotModified == nil {
					out.NotModified = new(NotModified)
				}
				(*out.NotModified).UnmarshalEasyJSON(in)
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent(out *jwriter.Writer, in Corge_NoContent_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.NotModified != nil {
		const prefix string = ",\"notModified\":"
		first = false
		out.RawString(prefix[1:])
		(*in.NotModified).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_NoContent_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_NoContent_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_NoContent_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_NoContent_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent1(in *jlexer.Lexer, out *Corge_NoContent_Args) {
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
		key := in.UnsafeFieldName(false)
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent1(out *jwriter.Writer, in Corge_NoContent_Args) {
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
func (v Corge_NoContent_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_NoContent_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_NoContent_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_NoContent_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContent1(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException(in *jlexer.Lexer, out *Corge_NoContentOnException_Result) {
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
		key := in.UnsafeFieldName(false)
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
					out.Success = new(Foo)
				}
				(*out.Success).UnmarshalEasyJSON(in)
			}
		case "notModified":
			if in.IsNull() {
				in.Skip()
				out.NotModified = nil
			} else {
				if out.NotModified == nil {
					out.NotModified = new(NotModified)
				}
				(*out.NotModified).UnmarshalEasyJSON(in)
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException(out *jwriter.Writer, in Corge_NoContentOnException_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Success).MarshalEasyJSON(out)
	}
	if in.NotModified != nil {
		const prefix string = ",\"notModified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.NotModified).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_NoContentOnException_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_NoContentOnException_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_NoContentOnException_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_NoContentOnException_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException1(in *jlexer.Lexer, out *Corge_NoContentOnException_Args) {
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
		key := in.UnsafeFieldName(false)
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException1(out *jwriter.Writer, in Corge_NoContentOnException_Args) {
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
func (v Corge_NoContentOnException_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_NoContentOnException_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_NoContentOnException_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_NoContentOnException_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentOnException1(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException(in *jlexer.Lexer, out *Corge_NoContentNoException_Result) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException(out *jwriter.Writer, in Corge_NoContentNoException_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Corge_NoContentNoException_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_NoContentNoException_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_NoContentNoException_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_NoContentNoException_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException(l, v)
}
func easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException1(in *jlexer.Lexer, out *Corge_NoContentNoException_Args) {
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
		key := in.UnsafeFieldName(false)
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
func easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException1(out *jwriter.Writer, in Corge_NoContentNoException_Args) {
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
func (v Corge_NoContentNoException_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_NoContentNoException_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAbe23ddeEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_NoContentNoException_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_NoContentNoException_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAbe23ddeDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeNoContentNoException1(l, v)
}
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
		key := in.UnsafeFieldName(false)
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
		key := in.UnsafeFieldName(false)
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
		key := in.UnsafeFieldName(false)
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
		key := in.UnsafeFieldName(false)
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
