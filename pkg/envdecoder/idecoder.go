package envdecoder

// Decoder is the interface implemented by types that can decode themselves from env values.
type Decoder interface {
	DecodeEnv() error
}
