package te

type B struct {
	A *int `env:"a" json:"a,omitempty"`
}

type A struct {
	*B
	C []string
}
