package envdecoder

import "errors"

// Decode will decode environment variables to given value
func Decode(val interface{}) error {
	if t, ok := val.(Decoder); ok {
		return t.DecodeEnv()
	}
	return errors.New("decoder interface is not implemented for  interface")
}
