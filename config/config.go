package config

type Stripe struct {
	SkKey string `env:"SK_KEY,required"`
	Port  int    `env:"PORT" envDefault:"8080"`
}
