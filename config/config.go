package config

type Config struct {
	SkKey string `env:"SK_KEY,required"`
	Port  int    `env:"PORT" envDefault:"8080"`

	IftttKey     string `env:"IFTTT_KEY"`
	IftttTurnOn  string `env:"IFTTT_TURN_ON"`
	IftttTurnOff string `env:"IFTTT_TURN_OFF"`
}
