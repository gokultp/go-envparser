package envdecoder

import "errors"

var (
	// ErrDecoderNotImplemented is thrown if the type has not implemeted the Decoder interface
	ErrDecoderNotImplemented = errors.New("env Decoder interface is not implemented")
)
