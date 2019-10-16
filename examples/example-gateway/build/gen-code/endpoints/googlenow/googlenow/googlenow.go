// Code generated by thriftrw v1.20.1. DO NOT EDIT.
// @generated

package googlenow

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

// GoogleNow_AddCredentials_Args represents the arguments for the GoogleNow.addCredentials function.
//
// The arguments for addCredentials are sent and received over the wire as this struct.
type GoogleNow_AddCredentials_Args struct {
	AuthCode string `json:"authCode,required"`
}

// ToWire translates a GoogleNow_AddCredentials_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *GoogleNow_AddCredentials_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.AuthCode), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GoogleNow_AddCredentials_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GoogleNow_AddCredentials_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GoogleNow_AddCredentials_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GoogleNow_AddCredentials_Args) FromWire(w wire.Value) error {
	var err error

	authCodeIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.AuthCode, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				authCodeIsSet = true
			}
		}
	}

	if !authCodeIsSet {
		return errors.New("field AuthCode of GoogleNow_AddCredentials_Args is required")
	}

	return nil
}

// String returns a readable string representation of a GoogleNow_AddCredentials_Args
// struct.
func (v *GoogleNow_AddCredentials_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("AuthCode: %v", v.AuthCode)
	i++

	return fmt.Sprintf("GoogleNow_AddCredentials_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this GoogleNow_AddCredentials_Args match the
// provided GoogleNow_AddCredentials_Args.
//
// This function performs a deep comparison.
func (v *GoogleNow_AddCredentials_Args) Equals(rhs *GoogleNow_AddCredentials_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.AuthCode == rhs.AuthCode) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GoogleNow_AddCredentials_Args.
func (v *GoogleNow_AddCredentials_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("authCode", v.AuthCode)
	return err
}

// GetAuthCode returns the value of AuthCode if it is set or its
// zero value if it is unset.
func (v *GoogleNow_AddCredentials_Args) GetAuthCode() (o string) {
	if v != nil {
		o = v.AuthCode
	}
	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "addCredentials" for this struct.
func (v *GoogleNow_AddCredentials_Args) MethodName() string {
	return "addCredentials"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *GoogleNow_AddCredentials_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// GoogleNow_AddCredentials_Helper provides functions that aid in handling the
// parameters and return values of the GoogleNow.addCredentials
// function.
var GoogleNow_AddCredentials_Helper = struct {
	// Args accepts the parameters of addCredentials in-order and returns
	// the arguments struct for the function.
	Args func(
		authCode string,
	) *GoogleNow_AddCredentials_Args

	// IsException returns true if the given error can be thrown
	// by addCredentials.
	//
	// An error can be thrown by addCredentials only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for addCredentials
	// given the error returned by it. The provided error may
	// be nil if addCredentials did not fail.
	//
	// This allows mapping errors returned by addCredentials into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// addCredentials
	//
	//   err := addCredentials(args)
	//   result, err := GoogleNow_AddCredentials_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from addCredentials: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*GoogleNow_AddCredentials_Result, error)

	// UnwrapResponse takes the result struct for addCredentials
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if addCredentials threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := GoogleNow_AddCredentials_Helper.UnwrapResponse(result)
	UnwrapResponse func(*GoogleNow_AddCredentials_Result) error
}{}

func init() {
	GoogleNow_AddCredentials_Helper.Args = func(
		authCode string,
	) *GoogleNow_AddCredentials_Args {
		return &GoogleNow_AddCredentials_Args{
			AuthCode: authCode,
		}
	}

	GoogleNow_AddCredentials_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	GoogleNow_AddCredentials_Helper.WrapResponse = func(err error) (*GoogleNow_AddCredentials_Result, error) {
		if err == nil {
			return &GoogleNow_AddCredentials_Result{}, nil
		}

		return nil, err
	}
	GoogleNow_AddCredentials_Helper.UnwrapResponse = func(result *GoogleNow_AddCredentials_Result) (err error) {
		return
	}

}

// GoogleNow_AddCredentials_Result represents the result of a GoogleNow.addCredentials function call.
//
// The result of a addCredentials execution is sent and received over the wire as this struct.
type GoogleNow_AddCredentials_Result struct {
}

// ToWire translates a GoogleNow_AddCredentials_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *GoogleNow_AddCredentials_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GoogleNow_AddCredentials_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GoogleNow_AddCredentials_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GoogleNow_AddCredentials_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GoogleNow_AddCredentials_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a GoogleNow_AddCredentials_Result
// struct.
func (v *GoogleNow_AddCredentials_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("GoogleNow_AddCredentials_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this GoogleNow_AddCredentials_Result match the
// provided GoogleNow_AddCredentials_Result.
//
// This function performs a deep comparison.
func (v *GoogleNow_AddCredentials_Result) Equals(rhs *GoogleNow_AddCredentials_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GoogleNow_AddCredentials_Result.
func (v *GoogleNow_AddCredentials_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "addCredentials" for this struct.
func (v *GoogleNow_AddCredentials_Result) MethodName() string {
	return "addCredentials"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *GoogleNow_AddCredentials_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

// GoogleNow_CheckCredentials_Args represents the arguments for the GoogleNow.checkCredentials function.
//
// The arguments for checkCredentials are sent and received over the wire as this struct.
type GoogleNow_CheckCredentials_Args struct {
}

// ToWire translates a GoogleNow_CheckCredentials_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *GoogleNow_CheckCredentials_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GoogleNow_CheckCredentials_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GoogleNow_CheckCredentials_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GoogleNow_CheckCredentials_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GoogleNow_CheckCredentials_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a GoogleNow_CheckCredentials_Args
// struct.
func (v *GoogleNow_CheckCredentials_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("GoogleNow_CheckCredentials_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this GoogleNow_CheckCredentials_Args match the
// provided GoogleNow_CheckCredentials_Args.
//
// This function performs a deep comparison.
func (v *GoogleNow_CheckCredentials_Args) Equals(rhs *GoogleNow_CheckCredentials_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GoogleNow_CheckCredentials_Args.
func (v *GoogleNow_CheckCredentials_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "checkCredentials" for this struct.
func (v *GoogleNow_CheckCredentials_Args) MethodName() string {
	return "checkCredentials"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *GoogleNow_CheckCredentials_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// GoogleNow_CheckCredentials_Helper provides functions that aid in handling the
// parameters and return values of the GoogleNow.checkCredentials
// function.
var GoogleNow_CheckCredentials_Helper = struct {
	// Args accepts the parameters of checkCredentials in-order and returns
	// the arguments struct for the function.
	Args func() *GoogleNow_CheckCredentials_Args

	// IsException returns true if the given error can be thrown
	// by checkCredentials.
	//
	// An error can be thrown by checkCredentials only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for checkCredentials
	// given the error returned by it. The provided error may
	// be nil if checkCredentials did not fail.
	//
	// This allows mapping errors returned by checkCredentials into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// checkCredentials
	//
	//   err := checkCredentials(args)
	//   result, err := GoogleNow_CheckCredentials_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from checkCredentials: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*GoogleNow_CheckCredentials_Result, error)

	// UnwrapResponse takes the result struct for checkCredentials
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if checkCredentials threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := GoogleNow_CheckCredentials_Helper.UnwrapResponse(result)
	UnwrapResponse func(*GoogleNow_CheckCredentials_Result) error
}{}

func init() {
	GoogleNow_CheckCredentials_Helper.Args = func() *GoogleNow_CheckCredentials_Args {
		return &GoogleNow_CheckCredentials_Args{}
	}

	GoogleNow_CheckCredentials_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	GoogleNow_CheckCredentials_Helper.WrapResponse = func(err error) (*GoogleNow_CheckCredentials_Result, error) {
		if err == nil {
			return &GoogleNow_CheckCredentials_Result{}, nil
		}

		return nil, err
	}
	GoogleNow_CheckCredentials_Helper.UnwrapResponse = func(result *GoogleNow_CheckCredentials_Result) (err error) {
		return
	}

}

// GoogleNow_CheckCredentials_Result represents the result of a GoogleNow.checkCredentials function call.
//
// The result of a checkCredentials execution is sent and received over the wire as this struct.
type GoogleNow_CheckCredentials_Result struct {
}

// ToWire translates a GoogleNow_CheckCredentials_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *GoogleNow_CheckCredentials_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a GoogleNow_CheckCredentials_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a GoogleNow_CheckCredentials_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v GoogleNow_CheckCredentials_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *GoogleNow_CheckCredentials_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a GoogleNow_CheckCredentials_Result
// struct.
func (v *GoogleNow_CheckCredentials_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("GoogleNow_CheckCredentials_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this GoogleNow_CheckCredentials_Result match the
// provided GoogleNow_CheckCredentials_Result.
//
// This function performs a deep comparison.
func (v *GoogleNow_CheckCredentials_Result) Equals(rhs *GoogleNow_CheckCredentials_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of GoogleNow_CheckCredentials_Result.
func (v *GoogleNow_CheckCredentials_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "checkCredentials" for this struct.
func (v *GoogleNow_CheckCredentials_Result) MethodName() string {
	return "checkCredentials"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *GoogleNow_CheckCredentials_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
