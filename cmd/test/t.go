package te

type A struct {
	Abc *int   `env:"A"`
	B   int64  `env:"B"`
	C   uint   `env:"C"`
	D   *uint8 `env:"D"`
}
